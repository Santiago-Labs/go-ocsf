package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// KeyboardInfoFields defines the Arrow fields for KeyboardInfo.
var KeyboardInfoFields = []arrow.Field{
	{Name: "function_keys", Type: arrow.PrimitiveTypes.Int32},
	{Name: "ime", Type: arrow.BinaryTypes.String},
	{Name: "keyboard_layout", Type: arrow.BinaryTypes.String},
	{Name: "keyboard_subtype", Type: arrow.PrimitiveTypes.Int32},
	{Name: "keyboard_type", Type: arrow.BinaryTypes.String},
}

var KeyboardInfoStruct = arrow.StructOf(KeyboardInfoFields...)
var KeyboardInfoClassname = "keyboard_info"

type KeyboardInfo struct {
	FunctionKeys    *int    `json:"function_keys,omitempty" parquet:"function_keys"`
	IME             *string `json:"ime,omitempty" parquet:"ime"`
	KeyboardLayout  *string `json:"keyboard_layout,omitempty" parquet:"keyboard_layout"`
	KeyboardSubtype *int    `json:"keyboard_subtype,omitempty" parquet:"keyboard_subtype"`
	KeyboardType    *string `json:"keyboard_type,omitempty" parquet:"keyboard_type"`
}
