package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ImageFields defines the fields for the Image Arrow schema.
var ImageFields = []arrow.Field{
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "tag", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var ImageStruct = arrow.StructOf(ImageFields...)
var ImageClassname = "image"

// Image represents image details.
type Image struct {
	Labels []*string `json:"labels,omitempty" parquet:"labels,list,optional"`
	Name   *string   `json:"name,omitempty" parquet:"name,optional"`
	Path   *string   `json:"path,omitempty" parquet:"path,optional"`
	Tag    *string   `json:"tag,omitempty" parquet:"tag,optional"`
	UID    string    `json:"uid" parquet:"uid"`
}
