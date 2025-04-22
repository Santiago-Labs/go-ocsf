package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var FileFields = []arrow.Field{
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "owner", Type: arrow.BinaryTypes.String, Nullable: false},
}

var FileStruct = arrow.StructOf(FileFields...)
var FileClassname = "file"

type File struct {
	Path  string `json:"path" parquet:"path"`
	Owner string `json:"owner" parquet:"owner"`
}
