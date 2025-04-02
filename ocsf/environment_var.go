package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// EnvironmentVarFields defines the Arrow fields for EnvironmentVar.
var EnvironmentVarFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

var EnvironmentVarStruct = arrow.StructOf(EnvironmentVarFields...)

type EnvironmentVar struct {
	Name  string `json:"name" parquet:"name"`
	Value string `json:"value" parquet:"value"`
}
