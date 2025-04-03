package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// KeyValueObjectFields defines the Arrow fields for KeyValueObject.
var KeyValueObjectFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
	{Name: "values", Type: arrow.ListOf(arrow.BinaryTypes.String)},
}

var KeyValueObjectStruct = arrow.StructOf(KeyValueObjectFields...)
var KeyValueObjectClassname = "key_value_object"

type KeyValueObject struct {
	Name   string    `json:"name" parquet:"name"`
	Value  *string   `json:"value,omitempty" parquet:"value"`
	Values []*string `json:"values,omitempty" parquet:"values"`
}
