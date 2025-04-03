package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// RequestFields defines the Arrow fields for Request.
var RequestFields = []arrow.Field{
	{Name: "containers", Type: arrow.ListOf(ContainerStruct)},
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "flags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var RequestStruct = arrow.StructOf(RequestFields...)
var RequestClassname = "request"

type Request struct {
	Containers []*Container `json:"containers,omitempty" parquet:"containers"`
	Data       *string      `json:"data,omitempty" parquet:"data"`
	Flags      []string     `json:"flags,omitempty" parquet:"flags"`
	UID        string       `json:"uid" parquet:"uid"`
}
