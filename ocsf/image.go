package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// ImageFields defines the fields for the Image Arrow schema.
var ImageFields = []arrow.Field{
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "tag", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// ImageSchema is the Arrow schema for Image.
var ImageSchema = arrow.NewSchema(ImageFields, nil)

// Image represents image details.
type Image struct {
	Labels []string `json:"labels,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Path   *string  `json:"path,omitempty"`
	Tag    *string  `json:"tag,omitempty"`
	UID    string   `json:"uid"`
}

// WriteToParquet writes the Image fields to the provided Arrow StructBuilder.
func (img *Image) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Labels (list of strings).
	labelsB := sb.FieldBuilder(0).(*array.ListBuilder)
	if len(img.Labels) > 0 {
		labelsB.Append(true) // Start a new list element.
		labelsValB := labelsB.ValueBuilder().(*array.StringBuilder)
		for _, label := range img.Labels {
			labelsValB.Append(label)
		}
	} else {
		labelsB.AppendNull()
	}

	// Field 1: Name.
	nameB := sb.FieldBuilder(1).(*array.StringBuilder)
	if img.Name != nil {
		nameB.Append(*img.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 2: Path.
	pathB := sb.FieldBuilder(2).(*array.StringBuilder)
	if img.Path != nil {
		pathB.Append(*img.Path)
	} else {
		pathB.AppendNull()
	}

	// Field 3: Tag.
	tagB := sb.FieldBuilder(3).(*array.StringBuilder)
	if img.Tag != nil {
		tagB.Append(*img.Tag)
	} else {
		tagB.AppendNull()
	}

	// Field 4: UID (required).
	uidB := sb.FieldBuilder(4).(*array.StringBuilder)
	uidB.Append(img.UID)
}
