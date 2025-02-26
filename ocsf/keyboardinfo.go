package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// KeyboardInfoFields defines the Arrow fields for KeyboardInfo.
var KeyboardInfoFields = []arrow.Field{
	{Name: "function_keys", Type: arrow.PrimitiveTypes.Int32},
	{Name: "ime", Type: arrow.BinaryTypes.String},
	{Name: "keyboard_layout", Type: arrow.BinaryTypes.String},
	{Name: "keyboard_subtype", Type: arrow.PrimitiveTypes.Int32},
	{Name: "keyboard_type", Type: arrow.BinaryTypes.String},
}

// KeyboardInfoSchema is the Arrow schema for KeyboardInfo.
var KeyboardInfoSchema = arrow.NewSchema(KeyboardInfoFields, nil)

// KeyboardInfo represents keyboard details.
type KeyboardInfo struct {
	FunctionKeys    *int    `json:"function_keys,omitempty"`
	IME             *string `json:"ime,omitempty"`
	KeyboardLayout  *string `json:"keyboard_layout,omitempty"`
	KeyboardSubtype *int    `json:"keyboard_subtype,omitempty"`
	KeyboardType    *string `json:"keyboard_type,omitempty"`
}

// WriteToParquet writes the KeyboardInfo fields to the provided Arrow StructBuilder.
func (ki *KeyboardInfo) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: FunctionKeys.
	fkB := sb.FieldBuilder(0).(*array.Int32Builder)
	if ki.FunctionKeys != nil {
		fkB.Append(int32(*ki.FunctionKeys))
	} else {
		fkB.AppendNull()
	}

	// Field 1: IME.
	imeB := sb.FieldBuilder(1).(*array.StringBuilder)
	if ki.IME != nil {
		imeB.Append(*ki.IME)
	} else {
		imeB.AppendNull()
	}

	// Field 2: KeyboardLayout.
	layoutB := sb.FieldBuilder(2).(*array.StringBuilder)
	if ki.KeyboardLayout != nil {
		layoutB.Append(*ki.KeyboardLayout)
	} else {
		layoutB.AppendNull()
	}

	// Field 3: KeyboardSubtype.
	subtypeB := sb.FieldBuilder(3).(*array.Int32Builder)
	if ki.KeyboardSubtype != nil {
		subtypeB.Append(int32(*ki.KeyboardSubtype))
	} else {
		subtypeB.AppendNull()
	}

	// Field 4: KeyboardType.
	typeB := sb.FieldBuilder(4).(*array.StringBuilder)
	if ki.KeyboardType != nil {
		typeB.Append(*ki.KeyboardType)
	} else {
		typeB.AppendNull()
	}
}
