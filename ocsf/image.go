package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
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
	Labels []string `json:"labels,omitempty" parquet:"labels"`
	Name   *string  `json:"name,omitempty" parquet:"name"`
	Path   *string  `json:"path,omitempty" parquet:"path"`
	Tag    *string  `json:"tag,omitempty" parquet:"tag"`
	UID    string   `json:"uid" parquet:"uid"`
}
