package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// CWEFields defines the Arrow fields for CWE.
var CWEFields = []arrow.Field{
	{Name: "caption", Type: arrow.BinaryTypes.String},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// CWESchema is the Arrow schema for CWE.
var CWESchema = arrow.NewSchema(CWEFields, nil)

type CWE struct {
	Caption   *string `json:"caption"`
	SourceURL *string `json:"src_url"`
	UID       string  `json:"uid"`
}

func (c *CWE) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: Caption.
	captionB := sb.FieldBuilder(0).(*array.StringBuilder)
	if c.Caption != nil {
		captionB.Append(*c.Caption)
	} else {
		captionB.AppendNull()
	}

	// Field 1: SourceURL.
	srcB := sb.FieldBuilder(1).(*array.StringBuilder)
	if c.SourceURL != nil {
		srcB.Append(*c.SourceURL)
	} else {
		srcB.AppendNull()
	}

	// Field 2: UID.
	uidB := sb.FieldBuilder(2).(*array.StringBuilder)
	uidB.Append(c.UID)
}
