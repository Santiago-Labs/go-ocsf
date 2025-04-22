package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// KeyValueObjectFields defines the Arrow fields for KeyValueObject.
var KeyValueObjectFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "values", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
}

var KeyValueObjectStruct = arrow.StructOf(KeyValueObjectFields...)
var KeyValueObjectClassname = "key_value_object"

type KeyValueObject struct {
	Name   string   `json:"name" parquet:"name"`
	Value  *string  `json:"value,omitempty" parquet:"value,optional"`
	Values []string `json:"values,omitempty" parquet:"values,list,optional"`
}
