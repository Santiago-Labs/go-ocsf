package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var FileFields = []arrow.Field{
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "owner", Type: arrow.BinaryTypes.String},
}

var FileSchema = arrow.NewSchema(FileFields, nil)

type File struct {
	Path  string `json:"path" parquet:"path"`
	Owner string `json:"owner" parquet:"owner"`
}
