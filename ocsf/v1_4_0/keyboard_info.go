// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_4_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type KeyboardInformation struct {

	// Function Keys: The number of function keys on client keyboard.
	FunctionKeys *int32 `json:"function_keys,omitempty" parquet:"function_keys,optional"`

	// IME: The Input Method Editor (IME) file name.
	Ime *string `json:"ime,omitempty" parquet:"ime,optional"`

	// Keyboard Layout: The keyboard locale identifier name (e.g., en-US).
	KeyboardLayout *string `json:"keyboard_layout,omitempty" parquet:"keyboard_layout,optional"`

	// Keyboard Subtype: The keyboard numeric code.
	KeyboardSubtype *int32 `json:"keyboard_subtype,omitempty" parquet:"keyboard_subtype,optional"`

	// Keyboard Type: The keyboard type (e.g., xt, ico).
	KeyboardType *string `json:"keyboard_type,omitempty" parquet:"keyboard_type,optional"`
}

var KeyboardInformationFields = []arrow.Field{
	{Name: "function_keys", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "ime", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "keyboard_layout", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "keyboard_subtype", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "keyboard_type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var KeyboardInformationStruct = arrow.StructOf(KeyboardInformationFields...)

var KeyboardInformationSchema = arrow.NewSchema(KeyboardInformationFields, nil)
var KeyboardInformationClassname = "keyboard_info"
