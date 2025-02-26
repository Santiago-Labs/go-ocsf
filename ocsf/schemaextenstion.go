package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// SchemaExtensionFields defines the Arrow fields for SchemaExtension.
var SchemaExtensionFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// SchemaExtensionSchema is the Arrow schema for SchemaExtension.
var SchemaExtensionSchema = arrow.NewSchema(SchemaExtensionFields, nil)

// SchemaExtension represents a schema extension.
type SchemaExtension struct {
	Name    string `json:"name"`
	UID     string `json:"uid"`
	Version string `json:"version"`
}

// WriteToParquet writes the SchemaExtension fields to the provided Arrow StructBuilder.
func (se *SchemaExtension) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(se.Name)

	// Field 1: uid.
	uidB := sb.FieldBuilder(1).(*array.StringBuilder)
	uidB.Append(se.UID)

	// Field 2: version.
	versionB := sb.FieldBuilder(2).(*array.StringBuilder)
	versionB.Append(se.Version)
}
