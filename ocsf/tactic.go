package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// TacticFields defines the Arrow fields for Tactic.
var TacticFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// TacticSchema is the Arrow schema for Tactic.
var TacticSchema = arrow.NewSchema(TacticFields, nil)

// Tactic represents a tactic.
type Tactic struct {
	Name   string `json:"name"`
	SrcURL string `json:"src_url"`
	UID    string `json:"uid"`
}

// WriteToParquet writes the Tactic fields to the provided Arrow StructBuilder.
func (t *Tactic) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(t.Name)

	// Field 1: src_url.
	srcURLB := sb.FieldBuilder(1).(*array.StringBuilder)
	srcURLB.Append(t.SrcURL)

	// Field 2: uid.
	uidB := sb.FieldBuilder(2).(*array.StringBuilder)
	uidB.Append(t.UID)
}
