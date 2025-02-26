package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// SubTechniqueFields defines the Arrow fields for SubTechnique.
var SubTechniqueFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// SubTechniqueSchema is the Arrow schema for SubTechnique.
var SubTechniqueSchema = arrow.NewSchema(SubTechniqueFields, nil)

// SubTechnique represents a sub-technique.
type SubTechnique struct {
	Name   string `json:"name"`
	SrcURL string `json:"src_url"`
	UID    string `json:"uid"`
}

// WriteToParquet writes the SubTechnique fields to the provided Arrow StructBuilder.
func (st *SubTechnique) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(st.Name)

	// Field 1: src_url.
	srcURLB := sb.FieldBuilder(1).(*array.StringBuilder)
	srcURLB.Append(st.SrcURL)

	// Field 2: uid.
	uidB := sb.FieldBuilder(2).(*array.StringBuilder)
	uidB.Append(st.UID)
}
