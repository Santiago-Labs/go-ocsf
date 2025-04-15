package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

type Observable struct {
	Name       string      `json:"name" parquet:"name"`
	Reputation *Reputation `json:"reputation,omitempty" parquet:"reputation"`
	Type       *string     `json:"type,omitempty" parquet:"type"`
	TypeID     int         `json:"type_id" parquet:"type_id"`
	Value      *string     `json:"value,omitempty" parquet:"value"`
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
