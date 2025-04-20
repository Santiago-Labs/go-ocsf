package datastore

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/Santiago-Labs/go-ocsf/ocsf"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/arrow-go/v18/parquet/file"    // For reading parquet back to arrow
	"github.com/apache/arrow-go/v18/parquet/pqarrow" // For reading parquet back to arrow
	"github.com/apache/iceberg-go"
	"github.com/apache/iceberg-go/table"
	"github.com/parquet-go/parquet-go"
)

func serializeFindingsToParquet(findings []ocsf.VulnerabilityFinding, schema *parquet.Schema) ([]byte, error) {
	buf := new(bytes.Buffer)

	pw := parquet.NewGenericWriter[ocsf.VulnerabilityFinding](buf, schema, parquet.Compression(&parquet.Gzip))
	defer pw.Close()

	if _, err := pw.Write(findings); err != nil {
		return nil, err
	}
	if err := pw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func serializeActivitiesToParquet(activities []ocsf.APIActivity, schema *parquet.Schema) ([]byte, error) {
	buf := new(bytes.Buffer)

	pw := parquet.NewGenericWriter[ocsf.APIActivity](buf, schema, parquet.Compression(&parquet.Gzip))
	defer pw.Close()

	if _, err := pw.Write(activities); err != nil {
		return nil, err
	}

	if err := pw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func parquetBytesToArrowTable(data []byte, mem memory.Allocator) (arrow.Table, error) {
	reader, err := file.NewParquetReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	arrowRdr, err := pqarrow.NewFileReader(reader, pqarrow.ArrowReadProperties{BatchSize: 1024}, mem)
	if err != nil {
		return nil, err
	}

	table, err := arrowRdr.ReadTable(context.Background())
	if err != nil {
		return nil, err
	}

	return table, nil
}

func icebergTypeToParquetNode(f iceberg.NestedField) parquet.Node {
	var node parquet.Node
	id := f.ID
	optional := !f.Required

	switch t := f.Type.(type) {
	case iceberg.BooleanType:
		node = parquet.Leaf(parquet.BooleanType)
	case iceberg.Int32Type:
		node = parquet.Leaf(parquet.Int32Type)
	case iceberg.Int64Type:
		node = parquet.Leaf(parquet.Int64Type)
	case iceberg.Float32Type:
		node = parquet.Leaf(parquet.FloatType)
	case iceberg.Float64Type:
		node = parquet.Leaf(parquet.DoubleType)
	case iceberg.DateType:
		node = parquet.Date()
	case iceberg.TimeType:
		node = parquet.Time(parquet.Millisecond)
	case iceberg.TimestampType:
		node = parquet.Timestamp(parquet.Millisecond)
	case iceberg.TimestampTzType:
		node = parquet.Timestamp(parquet.Microsecond)
	case iceberg.StringType:
		node = parquet.String()
	case *iceberg.StructType:
		group := parquet.Group{}
		for _, subField := range t.Fields() {
			group[subField.Name] = icebergTypeToParquetNode(subField)
		}
		node = group
	case *iceberg.ListType:
		elemF := t.ElementField()
		elemPayload := icebergTypeToParquetNode(elemF)

		node = parquet.List(elemPayload)
	case *iceberg.MapType:
		keyNode := icebergTypeToParquetNode(t.KeyField())
		keyNode = parquet.Required(keyNode) // keys are required

		valNode := icebergTypeToParquetNode(t.ValueField())

		node = parquet.Map(keyNode, valNode)
	default:
		log.Fatalf("Unsupported Iceberg type: %T (%v)", f.Type, f.Type)
	}

	if optional {
		node = parquet.Optional(node)
	} else {
		node = parquet.Required(node)
	}

	node = parquet.FieldID(node, id)

	return node
}

func buildParquetSchemaFromIceberg(tbl table.Table) *parquet.Schema {
	rootFields := tbl.Schema().Fields()
	rootGroup := parquet.Group{}

	for _, f := range rootFields {
		rootGroup[f.Name] = icebergTypeToParquetNode(f)
	}

	return parquet.NewSchema("root", rootGroup)
}

func applyFieldIDs(iceberg parquet.Node, parquetSchema parquet.Node) (parquet.Node, error) {
	if iceberg.Leaf() != parquetSchema.Leaf() {
		name := "<root>"
		if f, ok := iceberg.(parquet.Field); ok {
			name = f.Name()
		}
		return nil, fmt.Errorf("schema mismatch at field %q: one node is leaf, the other is group", name)
	}
	if iceberg.Optional() != parquetSchema.Optional() || iceberg.Repeated() != parquetSchema.Repeated() {
		name := "<root>"
		if f, ok := iceberg.(parquet.Field); ok {
			name = f.Name()
		}
		return nil, fmt.Errorf("schema nullability mismatch at field %q (Iceberg optional=%v, Parquet optional=%v)",
			name, iceberg.Optional(), parquetSchema.Optional())
	}

	if iceberg.Leaf() && parquetSchema.Leaf() {
		if iceberg.ID() != 0 {
			parquetSchema = parquet.FieldID(parquetSchema, iceberg.ID())
		}
		return parquetSchema, nil
	}

	childrenIce := iceberg.Fields()
	childrenPar := parquetSchema.Fields()

	parFieldByName := make(map[string]parquet.Node, len(childrenPar))
	for _, pf := range childrenPar {
		parFieldByName[pf.Name()] = pf
	}

	resultGroup := parquet.Group{}
	for _, iceField := range childrenIce {
		name := iceField.Name()
		parField, ok := parFieldByName[name]
		if !ok {
			return nil, fmt.Errorf("field %q present in Iceberg schema not found in Parquet schema", name)
		}
		updatedChild, err := applyFieldIDs(iceField, parField)
		if err != nil {
			return nil, err
		}
		resultGroup[name] = updatedChild
	}
	if len(childrenPar) != len(childrenIce) {
		return nil, fmt.Errorf("Parquet schema has fields not present in Iceberg schema")
	}

	var updatedNode parquet.Node = resultGroup
	if parquetSchema.Optional() {
		updatedNode = parquet.Optional(updatedNode)
	}
	if parquetSchema.Repeated() {
		updatedNode = parquet.Repeated(updatedNode)
	}
	if iceberg.ID() != 0 {
		updatedNode = parquet.FieldID(updatedNode, iceberg.ID())
	}
	return updatedNode, nil
}

func attachFieldIDs(arSchema *arrow.Schema, iceSchema *iceberg.Schema) (*arrow.Schema, error) {
	nextMeta := func(md arrow.Metadata, k, v string) arrow.Metadata {
		keys, vals := md.Keys(), md.Values()
		keys = append(keys, k)
		vals = append(vals, v)
		return arrow.NewMetadata(keys, vals)
	}

	var convert func(arrow.Field, iceberg.NestedField) (arrow.Field, error)

	convert = func(arField arrow.Field, iceField iceberg.NestedField) (arrow.Field, error) {
		md := arField.Metadata
		md = nextMeta(md, "iceberg.field_id", fmt.Sprint(iceField.ID))
		arField.Metadata = md

		switch it := iceField.Type.(type) {

		case *iceberg.StructType:
			st, ok := arField.Type.(*arrow.StructType)
			if !ok {
				return arrow.Field{}, fmt.Errorf(
					"field %q: Arrow ≠ Struct while Iceberg is Struct", arField.Name)
			}

			iceChildren := make(map[string]iceberg.NestedField, len(it.Fields()))
			for _, f := range it.Fields() {
				iceChildren[f.Name] = f
			}

			newChildren := make([]arrow.Field, st.NumFields())
			for i, childA := range st.Fields() {
				iceChild, ok := iceChildren[childA.Name]
				if !ok {
					return arrow.Field{}, fmt.Errorf(
						"struct field %q: child %q not present in Iceberg schema",
						arField.Name, childA.Name)
				}
				conv, err := convert(childA, iceChild)
				if err != nil {
					return arrow.Field{}, err
				}
				newChildren[i] = conv
			}
			arField.Type = arrow.StructOf(newChildren...)

		case *iceberg.ListType:
			al, ok := arField.Type.(*arrow.ListType)
			if !ok {
				return arrow.Field{}, fmt.Errorf(
					"field %q: Arrow ≠ List while Iceberg is List", arField.Name)
			}

			elemA := al.ElemField()
			elemIce := it.ElementField()

			convElem, err := convert(elemA, elemIce)
			if err != nil {
				return arrow.Field{}, err
			}

			if elemIce.Required {
				arField.Type = arrow.ListOfField(arrow.Field{
					Name:     convElem.Name,
					Type:     convElem.Type,
					Nullable: false,
					Metadata: convElem.Metadata,
				})
			} else {
				arField.Type = arrow.ListOfField(convElem)
			}

		case *iceberg.MapType:
			am, ok := arField.Type.(*arrow.MapType)
			if !ok {
				return arrow.Field{}, fmt.Errorf(
					"field %q: Arrow ≠ Map while Iceberg is Map", arField.Name)
			}

			keyA, valA := am.KeyField(), am.ItemField()
			keyIce, valIce := it.KeyField(), it.ValueField()

			newKey, err := convert(keyA, keyIce)
			if err != nil {
				return arrow.Field{}, err
			}
			newVal, err := convert(valA, valIce)
			if err != nil {
				return arrow.Field{}, err
			}

			arField.Type = arrow.MapOf(newKey.Type, newVal.Type)
		}
		return arField, nil
	}

	newFields := make([]arrow.Field, len(arSchema.Fields()))
	for i, af := range arSchema.Fields() {
		iceF, ok := iceSchema.FindFieldByName(af.Name)
		if !ok {
			return nil, fmt.Errorf("column %q not found in Iceberg schema", af.Name)
		}
		nf, err := convert(af, iceF)
		if err != nil {
			return nil, err
		}
		newFields[i] = nf
	}

	return arrow.NewSchema(newFields, nil), nil
}

func fieldByName(s *iceberg.StructType, name string) (iceberg.NestedField, bool) {
	for _, f := range s.Fields() {
		if f.Name == name {
			return f, true
		}
	}
	return iceberg.NestedField{}, false
}

func dumpArray(f arrow.Field, arr arrow.Array, indent string) {
	fmt.Printf("%s%s (nullable=%v)  len=%d  nulls=%d  type=%s\n",
		indent, f.Name, f.Nullable, arr.Len(), arr.NullN(), arr.DataType())

	switch a := arr.(type) {
	case *array.List:
		values := a.ListValues()
		dumpArray(
			arrow.Field{Name: f.Name + ".element", Nullable: true, Type: values.DataType()},
			values, indent+"  ",
		)
	case *array.Struct:
		sa := a
		for j := 0; j < sa.NumField(); j++ {
			childField := sa.Field(j)
			childArrowField := a.DataType().(*arrow.StructType).Field(j)
			dumpArray(
				arrow.Field{
					Name:     childArrowField.Name,
					Nullable: childArrowField.Nullable,
					Type:     childArrowField.Type,
				},
				childField, indent+"  ",
			)
		}
	}
}
