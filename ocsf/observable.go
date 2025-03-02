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
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "reputation", Type: arrow.StructOf(ReputationFields...)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "value", Type: arrow.BinaryTypes.String},
}
