package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// TacticFields defines the Arrow fields for Tactic.
var TacticFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var TacticStruct = arrow.StructOf(TacticFields...)
var TacticClassname = "tactic"

type Tactic struct {
	Name   string `json:"name" parquet:"name"`
	SrcURL string `json:"src_url" parquet:"src_url"`
	UID    string `json:"uid" parquet:"uid"`
}
