package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var TechniqueFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var TechniqueStruct = arrow.StructOf(TechniqueFields...)
var TechniqueClassname = "technique"

type Technique struct {
	Name   string `json:"name" parquet:"name"`
	SrcURL string `json:"src_url" parquet:"src_url"`
	UID    string `json:"uid" parquet:"uid"`
}
