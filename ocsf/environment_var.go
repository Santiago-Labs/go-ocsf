package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// EnvironmentVarFields defines the Arrow fields for EnvironmentVar.
var EnvironmentVarFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var EnvironmentVarStruct = arrow.StructOf(EnvironmentVarFields...)
var EnvironmentVarClassname = "environment_variable"

type EnvironmentVar struct {
	Name  string `json:"name" parquet:"name" ch:"name"`
	Value string `json:"value" parquet:"value" ch:"value"`
}
