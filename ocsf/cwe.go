package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var CWEFields = []arrow.Field{
	{Name: "caption", Type: arrow.BinaryTypes.String},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var CWEStruct = arrow.StructOf(CWEFields...)
var CWESchema = arrow.NewSchema(CWEFields, nil)

type CWE struct {
	Caption   *string `json:"caption" parquet:"caption"`
	SourceURL *string `json:"src_url" parquet:"src_url"`
	UID       string  `json:"uid" parquet:"uid"`
}
