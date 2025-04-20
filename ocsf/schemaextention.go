package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var SchemaExtensionFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var SchemaExtensionStruct = arrow.StructOf(SchemaExtensionFields...)
var SchemaExtensionClassname = "extension"

type SchemaExtension struct {
	Name    string `json:"name" parquet:"name"`
	UID     string `json:"uid" parquet:"uid"`
	Version string `json:"version" parquet:"version"`
}
