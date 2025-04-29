package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AutonomousSystemFields defines the Arrow fields for AutonomousSystem.
var AutonomousSystemFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "number", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
}

var AutonomousSystemStruct = arrow.StructOf(AutonomousSystemFields...)
var AutonomousSystemClassname = "autonomous_system"

type AutonomousSystem struct {
	Name   *string `json:"name,omitempty" parquet:"name,optional"`
	Number *int32  `json:"number,omitempty" parquet:"number,optional"`
}
