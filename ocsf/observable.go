package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type Observable struct {
	Name       string      `json:"name" parquet:"name" ch:"name"`
	Reputation *Reputation `json:"reputation" parquet:"reputation,optional" ch:"reputation"`
	Type       *string     `json:"type" parquet:"type,optional" ch:"type"`
	TypeID     int         `json:"type_id" parquet:"type_id" ch:"type_id"`
	Value      *string     `json:"value" parquet:"value,optional" ch:"value"`
}

var ObservableFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "reputation", Type: ReputationStruct, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ObservableStruct = arrow.StructOf(ObservableFields...)
var ObservableClassname = "observable"
