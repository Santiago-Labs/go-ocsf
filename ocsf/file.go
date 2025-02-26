package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var FileFields = []arrow.Field{
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "owner", Type: arrow.BinaryTypes.String},
}

var FileSchema = arrow.NewSchema(FileFields, nil)

// File represents file details.
type File struct {
	Path  string `json:"path"`
	Owner string `json:"owner"`
}

// WriteToParquet writes the File fields to the provided Arrow StructBuilder.
func (f *File) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Path.
	pathB := sb.FieldBuilder(0).(*array.StringBuilder)
	pathB.Append(f.Path)

	// Field 1: Owner.
	ownerB := sb.FieldBuilder(1).(*array.StringBuilder)
	ownerB.Append(f.Owner)
}
