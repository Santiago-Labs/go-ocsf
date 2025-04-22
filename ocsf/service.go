package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ServiceFields defines the Arrow fields for Service.
var ServiceFields = []arrow.Field{
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "tags", Type: arrow.ListOf(KeyValueObjectStruct), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ServiceStruct = arrow.StructOf(ServiceFields...)
var ServiceClassname = "service"

type Service struct {
	Labels  []string          `json:"labels,omitempty" parquet:"labels,list,optional"`
	Name    *string           `json:"name,omitempty" parquet:"name,optional"`
	Tags    []*KeyValueObject `json:"tags,omitempty" parquet:"tags,list,optional"`
	UID     *string           `json:"uid,omitempty" parquet:"uid,optional"`
	Version *string           `json:"version,omitempty" parquet:"version,optional"`
}
