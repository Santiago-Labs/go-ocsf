package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AutonomousSystemFields defines the Arrow fields for AutonomousSystem.
var AutonomousSystemFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "number", Type: arrow.PrimitiveTypes.Int32},
}

var AutonomousSystemStruct = arrow.StructOf(AutonomousSystemFields...)
var AutonomousSystemClassname = "autonomous_system"

type AutonomousSystem struct {
	Name   *string `json:"name,omitempty" parquet:"name"`
	Number *int    `json:"number,omitempty" parquet:"number"`
}
