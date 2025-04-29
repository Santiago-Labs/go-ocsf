package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// KeyboardInfoFields defines the Arrow fields for KeyboardInfo.
var KeyboardInfoFields = []arrow.Field{
	{Name: "function_keys", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "ime", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "keyboard_layout", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "keyboard_subtype", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "keyboard_type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var KeyboardInfoStruct = arrow.StructOf(KeyboardInfoFields...)
var KeyboardInfoClassname = "keyboard_info"

type KeyboardInfo struct {
	FunctionKeys    *int32  `json:"function_keys,omitempty" parquet:"function_keys,optional"`
	IME             *string `json:"ime,omitempty" parquet:"ime,optional"`
	KeyboardLayout  *string `json:"keyboard_layout,omitempty" parquet:"keyboard_layout,optional"`
	KeyboardSubtype *int32  `json:"keyboard_subtype,omitempty" parquet:"keyboard_subtype,optional"`
	KeyboardType    *string `json:"keyboard_type,omitempty" parquet:"keyboard_type,optional"`
}
