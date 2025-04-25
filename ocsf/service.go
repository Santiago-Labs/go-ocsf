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
	Labels  []string          `json:"labels" parquet:"labels,list,optional" ch:"labels"`
	Name    *string           `json:"name" parquet:"name,optional" ch:"name"`
	Tags    []*KeyValueObject `json:"tags" parquet:"tags,list,optional" ch:"tags"`
	UID     *string           `json:"uid" parquet:"uid,optional" ch:"uid"`
	Version *string           `json:"version" parquet:"version,optional" ch:"version"`
}
