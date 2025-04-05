package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var SchemaExtensionFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var SchemaExtensionStruct = arrow.StructOf(SchemaExtensionFields...)
var SchemaExtensionClassname = "extension"

type SchemaExtension struct {
	Name    string `json:"name" parquet:"name"`
	UID     string `json:"uid" parquet:"uid"`
	Version string `json:"version" parquet:"version"`
}
