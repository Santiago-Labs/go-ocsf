package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var SubTechniqueFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var SubTechniqueStruct = arrow.StructOf(SubTechniqueFields...)
var SubTechniqueClassname = "sub_technique"

type SubTechnique struct {
	Name   string `json:"name" parquet:"name"`
	SrcURL string `json:"src_url" parquet:"src_url"`
	UID    string `json:"uid" parquet:"uid"`
}
