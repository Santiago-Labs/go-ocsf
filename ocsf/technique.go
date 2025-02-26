package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// TechniqueFields defines the Arrow fields for Technique.
var TechniqueFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// TechniqueSchema is the Arrow schema for Technique.
var TechniqueSchema = arrow.NewSchema(TechniqueFields, nil)

// Technique represents a technique.
type Technique struct {
	Name   string `json:"name"`
	SrcURL string `json:"src_url"`
	UID    string `json:"uid"`
}

// WriteToParquet writes the Technique fields to the provided Arrow StructBuilder.
func (t *Technique) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(t.Name)

	// Field 1: SrcURL.
	srcURLB := sb.FieldBuilder(1).(*array.StringBuilder)
	srcURLB.Append(t.SrcURL)

	// Field 2: UID.
	uidB := sb.FieldBuilder(2).(*array.StringBuilder)
	uidB.Append(t.UID)
}
