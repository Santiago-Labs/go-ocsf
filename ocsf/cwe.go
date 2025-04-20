package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var CWEFields = []arrow.Field{
	{Name: "caption", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var CWEStruct = arrow.StructOf(CWEFields...)
var CWEClassname = "cwe"

type CWE struct {
	Caption   *string `json:"caption" parquet:"caption,optional"`
	SourceURL *string `json:"src_url" parquet:"src_url,optional"`
	UID       string  `json:"uid" parquet:"uid"`
}
