package datastore

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/apache/iceberg-go"
)

func SliceToRecordBatch(slice interface{}, schema *arrow.Schema) (arrow.Record, error) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("SliceToRecordBatch: input is not a slice")
	}
	nrows := val.Len()
	elemType := val.Type().Elem()
	var ptrToStruct bool
	if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
		ptrToStruct = true
		elemType = elemType.Elem()
	}
	if elemType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("SliceToRecordBatch: slice elements are not structs")
	}

	fieldIndexByTag := make(map[string]int)
	for i := 0; i < elemType.NumField(); i++ {
		sf := elemType.Field(i)
		tag := sf.Tag.Get("parquet")
		if tag != "" {
			tag = strings.Split(tag, ",")[0]
			fieldIndexByTag[tag] = i
		}
	}

	pool := memory.DefaultAllocator
	builders := make([]array.Builder, len(schema.Fields()))
	fieldIndices := make([]int, len(schema.Fields()))
	for j, field := range schema.Fields() {
		idx, ok := fieldIndexByTag[field.Name]
		if !ok {
			return nil, fmt.Errorf("no struct field with parquet tag %q", field.Name)
		}
		var b array.Builder
		dt := field.Type
		switch dt := dt.(type) {
		case *arrow.Int8Type:
			b = array.NewInt8Builder(pool)
		case *arrow.Int16Type:
			b = array.NewInt16Builder(pool)
		case *arrow.Int32Type:
			b = array.NewInt32Builder(pool)
		case *arrow.Int64Type:
			b = array.NewInt64Builder(pool)
		case *arrow.Uint8Type:
			b = array.NewUint8Builder(pool)
		case *arrow.Uint16Type:
			b = array.NewUint16Builder(pool)
		case *arrow.Uint32Type:
			b = array.NewUint32Builder(pool)
		case *arrow.Uint64Type:
			b = array.NewUint64Builder(pool)
		case *arrow.Float32Type:
			b = array.NewFloat32Builder(pool)
		case *arrow.Float64Type:
			b = array.NewFloat64Builder(pool)
		case *arrow.BooleanType:
			b = array.NewBooleanBuilder(pool)
		case *arrow.StringType:
			b = array.NewStringBuilder(pool)
		case *arrow.BinaryType:
			b = array.NewBinaryBuilder(pool, arrow.BinaryTypes.Binary)
		case *arrow.FixedSizeBinaryType:
			b = array.NewFixedSizeBinaryBuilder(pool, dt)
		case *arrow.Decimal128Type:
			b = array.NewDecimal128Builder(pool, dt)
		case *arrow.Decimal256Type:
			b = array.NewDecimal256Builder(pool, dt)
		case *arrow.Date32Type:
			b = array.NewDate32Builder(pool)
		case *arrow.Date64Type:
			b = array.NewDate64Builder(pool)
		case *arrow.TimestampType:
			b = array.NewTimestampBuilder(pool, dt)
		case *arrow.ListType:
			elemType := dt.Elem()
			b = array.NewListBuilder(pool, elemType)
		case *arrow.FixedSizeListType:
			elemType := dt.Elem()
			listSize := dt.Len()
			b = array.NewFixedSizeListBuilder(pool, listSize, elemType)
		case *arrow.StructType:
			b = array.NewStructBuilder(pool, dt)
		case *arrow.MapType:
			keyType := dt.KeyType()
			itemType := dt.ItemType()
			b = array.NewMapBuilder(pool, keyType, itemType, dt.KeysSorted)
		default:

			return nil, fmt.Errorf("unsupported arrow type: %s", dt)
		}
		builders[j] = b
		fieldIndices[j] = idx
	}

	var appendValue func(b array.Builder, arrowType arrow.DataType, val reflect.Value) error
	appendValue = func(b array.Builder, arrowType arrow.DataType, val reflect.Value) error {
		if !val.IsValid() {

			switch builder := b.(type) {
			case *array.StructBuilder:
				builder.AppendNull()
			case *array.ListBuilder:
				builder.AppendNull()
			case *array.FixedSizeListBuilder:
				builder.AppendNull()
			case *array.MapBuilder:
				builder.AppendNull()
			default:
				builder.(array.Builder).AppendNull()
			}
			return nil
		}

		if val.Kind() == reflect.Ptr {
			if val.IsNil() {

				switch builder := b.(type) {
				case *array.StructBuilder:
					builder.AppendNull()
				case *array.ListBuilder:
					builder.AppendNull()
				case *array.FixedSizeListBuilder:
					builder.AppendNull()
				case *array.MapBuilder:
					builder.AppendNull()
				default:
					builder.(array.Builder).AppendNull()
				}
				return nil
			}
			val = val.Elem()
		}

		switch builder := b.(type) {

		case *array.Int8Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int8 {
				return fmt.Errorf("type mismatch: expected int8 for arrow INT8. got %s", val.Kind())
			}
			builder.Append(int8(val.Int()))
		case *array.Int16Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int16 {
				return fmt.Errorf("type mismatch: expected int16 for arrow INT16. got %s", val.Kind())
			}
			builder.Append(int16(val.Int()))
		case *array.Int32Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int32 {
				return fmt.Errorf("type mismatch: expected int32 for arrow INT32. got %s", val.Kind())
			}
			builder.Append(int32(val.Int()))
		case *array.Int64Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int64 {
				return fmt.Errorf("type mismatch: expected int64 for arrow INT64. got %s", val.Kind())
			}
			builder.Append(val.Int())
		case *array.Uint8Builder:
			if val.Kind() != reflect.Uint && val.Kind() != reflect.Uint8 {
				return fmt.Errorf("type mismatch: expected uint8 for arrow UINT8. got %s", val.Kind())
			}
			builder.Append(uint8(val.Uint()))
		case *array.Uint16Builder:
			if val.Kind() != reflect.Uint && val.Kind() != reflect.Uint16 {
				return fmt.Errorf("type mismatch: expected uint16 for arrow UINT16. got %s", val.Kind())
			}
			builder.Append(uint16(val.Uint()))
		case *array.Uint32Builder:
			if val.Kind() != reflect.Uint && val.Kind() != reflect.Uint32 {
				return fmt.Errorf("type mismatch: expected uint32 for arrow UINT32. got %s", val.Kind())
			}
			builder.Append(uint32(val.Uint()))
		case *array.Uint64Builder:
			if val.Kind() != reflect.Uint && val.Kind() != reflect.Uint64 {
				return fmt.Errorf("type mismatch: expected uint64 for arrow UINT64. got %s", val.Kind())
			}
			builder.Append(val.Uint())
		case *array.Float32Builder:
			if val.Kind() != reflect.Float32 && val.Kind() != reflect.Float64 {
				return fmt.Errorf("type mismatch: expected float32 for arrow FLOAT32. got %s", val.Kind())
			}
			builder.Append(float32(val.Float()))
		case *array.Float64Builder:
			if val.Kind() != reflect.Float32 && val.Kind() != reflect.Float64 {
				return fmt.Errorf("type mismatch: expected float64 for arrow FLOAT64. got %s", val.Kind())
			}
			builder.Append(val.Float())
		case *array.BooleanBuilder:
			if val.Kind() != reflect.Bool {
				return fmt.Errorf("type mismatch: expected bool for arrow BOOL. got %s", val.Kind())
			}
			builder.Append(val.Bool())
		case *array.StringBuilder:
			if val.Kind() != reflect.String {
				return fmt.Errorf("type mismatch: expected string for arrow STRING. got %s", val.Kind())
			}
			builder.AppendString(val.String())
		case *array.BinaryBuilder:

			if val.Kind() == reflect.Slice && val.Type().Elem().Kind() == reflect.Uint8 {
				builder.Append(val.Bytes())
			} else if val.Kind() == reflect.String {
				builder.Append([]byte(val.String()))
			} else {
				return fmt.Errorf("type mismatch: expected []byte for arrow BINARY. got %s", val.Kind())
			}
		case *array.Date32Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int32 {
				return fmt.Errorf("type mismatch: expected int32 (days) for arrow DATE32. got %s", val.Kind())
			}
			builder.Append(arrow.Date32(val.Int()))
		case *array.Date64Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int64 {
				return fmt.Errorf("type mismatch: expected int64 (ms) for arrow DATE64. got %s", val.Kind())
			}
			builder.Append(arrow.Date64(val.Int()))
		case *array.Time32Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int32 {
				return fmt.Errorf("type mismatch: expected int32 for arrow TIME32. got %s", val.Kind())
			}
			builder.Append(arrow.Time32(val.Int()))
		case *array.Time64Builder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int64 {
				return fmt.Errorf("type mismatch: expected int64 for arrow TIME64. got %s", val.Kind())
			}
			builder.Append(arrow.Time64(val.Int()))
		case *array.TimestampBuilder:
			if val.Kind() != reflect.Int && val.Kind() != reflect.Int64 {
				return fmt.Errorf("type mismatch: expected int64 for arrow TIMESTAMP. got %s", val.Kind())
			}
			builder.Append(arrow.Timestamp(val.Int()))

		case *array.StructBuilder:

			if val.Kind() != reflect.Struct {
				return fmt.Errorf("type mismatch: expected struct for arrow STRUCT. got %s", val.Kind())
			}

			builder.Append(true)
			structType := val.Type()
			structFieldsByTag := make(map[string]int)
			for i := 0; i < structType.NumField(); i++ {
				tag := structType.Field(i).Tag.Get("parquet")
				if tag != "" {
					tag = strings.Split(tag, ",")[0]
					structFieldsByTag[tag] = i
				}
			}
			structDT := arrowType.(*arrow.StructType)
			for childIdx, childField := range structDT.Fields() {
				childBuilder := builder.FieldBuilder(childIdx)
				if childBuilder == nil {
					return fmt.Errorf("failed to get child builder for struct field %q", childField.Name)
				}
				sfIdx, ok := structFieldsByTag[childField.Name]
				if !ok {

					return fmt.Errorf("no struct field with tag %q in nested struct", childField.Name)
				}
				childVal := val.Field(sfIdx)
				if err := appendValue(childBuilder, childField.Type, childVal); err != nil {
					return err
				}
			}
		case *array.ListBuilder:
			if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
				return fmt.Errorf("type mismatch: expected slice/array for arrow LIST. got %s", val.Kind())
			}
			if val.Kind() == reflect.Slice && val.IsNil() {
				builder.AppendNull()
				break
			}

			builder.Append(true)
			elemType := arrowType.(*arrow.ListType).Elem()
			valueBuilder := builder.ValueBuilder()
			length := val.Len()
			for i := 0; i < length; i++ {
				itemVal := val.Index(i)
				if err := appendValue(valueBuilder, elemType, itemVal); err != nil {
					return err
				}
			}
		case *array.FixedSizeListBuilder:

			if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
				return fmt.Errorf("type mismatch: expected array for arrow FixedSizeList. got %s", val.Kind())
			}
			if val.Kind() == reflect.Slice && val.IsNil() {
				builder.AppendNull()
				break
			}

			listLen := arrowType.(*arrow.FixedSizeListType).Len()
			if val.Len() != int(listLen) {
				return fmt.Errorf("length mismatch: expected %d elements for fixed-size list, got %d", listLen, val.Len())
			}
			builder.Append(true)
			elemType := arrowType.(*arrow.FixedSizeListType).Elem()
			valueBuilder := builder.ValueBuilder()

			for i := 0; i < val.Len(); i++ {
				itemVal := val.Index(i)
				if err := appendValue(valueBuilder, elemType, itemVal); err != nil {
					return err
				}
			}
		case *array.MapBuilder:

			if val.Kind() != reflect.Map {
				return fmt.Errorf("type mismatch: expected map for arrow MAP. got %s", val.Kind())
			}
			if val.IsNil() {
				builder.AppendNull()
				break
			}
			builder.Append(true)
			keyBuilder := builder.KeyBuilder()
			itemBuilder := builder.ItemBuilder()
			iter := val.MapRange()

			for iter.Next() {
				keyVal := iter.Key()
				valVal := iter.Value()

				keyType := arrowType.(*arrow.MapType).KeyType()
				if err := appendValue(keyBuilder, keyType, keyVal); err != nil {
					return err
				}

				itemType := arrowType.(*arrow.MapType).ItemType()
				if err := appendValue(itemBuilder, itemType, valVal); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unsupported builder type: %T", b)
		}
		return nil
	}

	for i := 0; i < nrows; i++ {
		elemVal := val.Index(i)
		if ptrToStruct {
			if elemVal.IsNil() {
				for j, b := range builders {
					switch b.(type) {
					case *array.StructBuilder, *array.ListBuilder, *array.FixedSizeListBuilder, *array.MapBuilder:

						appendValue(b, schema.Field(j).Type, reflect.Value{})
					default:
						b.AppendNull()
					}
				}
				continue
			}
			elemVal = elemVal.Elem()
		}
		for j, field := range schema.Fields() {
			b := builders[j]
			structFieldIdx := fieldIndices[j]
			fieldVal := elemVal.Field(structFieldIdx)
			err := appendValue(b, field.Type, fieldVal)
			if err != nil {
				for _, rb := range builders {
					if rb != nil {
						rb.Release()
					}
				}
				return nil, fmt.Errorf("error appending field %q (row %d): %w", field.Name, i, err)
			}
		}
	}

	arrays := make([]arrow.Array, len(builders))
	for j, b := range builders {
		arr := b.NewArray()
		arrays[j] = arr
		b.Release()
	}
	rec := array.NewRecord(schema, arrays, int64(nrows))
	return rec, nil
}

func ArrowSchemaToIceberg(schema *arrow.Schema) (*iceberg.Schema, error) {
	var nextID int = 1

	var convertField func(field arrow.Field) (iceberg.NestedField, error)
	convertField = func(field arrow.Field) (iceberg.NestedField, error) {
		var icebergType iceberg.Type
		switch t := field.Type.(type) {

		case *arrow.BooleanType:
			icebergType = iceberg.BooleanType{}
		case *arrow.Int8Type, *arrow.Int16Type, *arrow.Int32Type:
			icebergType = iceberg.Int32Type{}
		case *arrow.Int64Type:
			icebergType = iceberg.Int64Type{}
		case *arrow.Uint8Type, *arrow.Uint16Type:
			icebergType = iceberg.Int32Type{}
		case *arrow.Uint32Type:
			icebergType = iceberg.Int64Type{}
		case *arrow.Uint64Type:

			icebergType = iceberg.Int64Type{}
		case *arrow.Float32Type:
			icebergType = iceberg.Float32Type{}
		case *arrow.Float64Type:
			icebergType = iceberg.Float64Type{}
		case *arrow.StringType:
			icebergType = iceberg.StringType{}
		case *arrow.BinaryType:
			icebergType = iceberg.BinaryType{}
		case *arrow.FixedSizeBinaryType:
			icebergType = iceberg.FixedTypeOf(t.ByteWidth)
		case *arrow.Decimal128Type:
			icebergType = iceberg.DecimalTypeOf(int(t.Precision), int(t.Scale))
		case *arrow.Decimal256Type:
			icebergType = iceberg.DecimalTypeOf(int(t.Precision), int(t.Scale))
		case *arrow.Date32Type, *arrow.Date64Type:
			icebergType = iceberg.DateType{}
		case *arrow.Time32Type, *arrow.Time64Type:
			icebergType = iceberg.TimeType{}
		case *arrow.TimestampType:
			if t.TimeZone != "" {
				icebergType = iceberg.TimestampTzType{}
			} else {
				icebergType = iceberg.TimestampType{}
			}

		case *arrow.StructType:
			fields := make([]iceberg.NestedField, 0, len(t.Fields()))
			for _, child := range t.Fields() {
				nf, err := convertField(child)
				if err != nil {
					return iceberg.NestedField{}, err
				}
				fields = append(fields, nf)
			}
			structType := &iceberg.StructType{FieldList: fields}
			icebergType = structType
		case *arrow.ListType:
			elemField := t.ElemField()
			nf, err := convertField(elemField)
			if err != nil {
				return iceberg.NestedField{}, err
			}
			listType := &iceberg.ListType{
				ElementID:       nf.ID,
				Element:         nf.Type,
				ElementRequired: nf.Required,
			}
			icebergType = listType
		case *arrow.MapType:
			keyField := t.KeyField()
			valField := t.ItemField()
			keyNF, err := convertField(keyField)
			if err != nil {
				return iceberg.NestedField{}, err
			}
			valNF, err := convertField(valField)
			if err != nil {
				return iceberg.NestedField{}, err
			}
			keyNF.Required = true
			mapType := &iceberg.MapType{
				KeyID:         keyNF.ID,
				KeyType:       keyNF.Type,
				ValueID:       valNF.ID,
				ValueType:     valNF.Type,
				ValueRequired: valNF.Required,
			}
			icebergType = mapType

		default:
			return iceberg.NestedField{}, fmt.Errorf("unsupported Arrow type: %s", t)
		}

		fieldID := nextID
		nextID++

		isRequired := !field.Nullable
		fieldName := field.Name

		nf := iceberg.NestedField{
			ID:       fieldID,
			Name:     fieldName,
			Type:     icebergType,
			Required: isRequired,
		}
		return nf, nil
	}

	arrowFields := schema.Fields()
	icebergFields := make([]iceberg.NestedField, 0, len(arrowFields))
	for _, f := range arrowFields {
		nf, err := convertField(f)
		if err != nil {
			return nil, err
		}
		icebergFields = append(icebergFields, nf)
	}

	icebergSchema := iceberg.NewSchema(0, icebergFields...)
	return icebergSchema, nil
}
