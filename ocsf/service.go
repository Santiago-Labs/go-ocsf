package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// ServiceFields defines the Arrow fields for Service.
var ServiceFields = []arrow.Field{
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "tags", Type: arrow.ListOf(KeyValueObjectStruct)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var ServiceStruct = arrow.StructOf(ServiceFields...)
var ServiceClassname = "service"

type Service struct {
	Labels  []string          `json:"labels,omitempty" parquet:"labels"`
	Name    *string           `json:"name,omitempty" parquet:"name"`
	Tags    []*KeyValueObject `json:"tags,omitempty" parquet:"tags"`
	UID     *string           `json:"uid,omitempty" parquet:"uid"`
	Version *string           `json:"version,omitempty" parquet:"version"`
}
