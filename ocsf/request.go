package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// RequestFields defines the Arrow fields for Request.
var RequestFields = []arrow.Field{
	{Name: "containers", Type: arrow.ListOf(ContainerStruct), Nullable: true},
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "flags", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var RequestStruct = arrow.StructOf(RequestFields...)
var RequestClassname = "request"

type Request struct {
	Containers []*Container `json:"containers,omitempty" parquet:"containers,list,optional"`
	Data       *string      `json:"data,omitempty" parquet:"data,optional"`
	Flags      []string     `json:"flags,omitempty" parquet:"flags,list,optional"`
	UID        string       `json:"uid" parquet:"uid"`
}
