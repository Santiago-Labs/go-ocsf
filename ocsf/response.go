package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ResponseFields defines the Arrow fields for Response.
var ResponseFields = []arrow.Field{
	{Name: "code", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "containers", Type: arrow.ListOf(ContainerStruct), Nullable: true},
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "error", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "error_message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "flags", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "message", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ResponseStruct = arrow.StructOf(ResponseFields...)
var ResponseClassname = "response"

type Response struct {
	Code         *int32       `json:"code,omitempty" parquet:"code,optional"`
	Containers   []*Container `json:"containers,omitempty" parquet:"containers,list,optional"`
	Data         *string      `json:"data,omitempty" parquet:"data,optional"`
	Error        *string      `json:"error,omitempty" parquet:"error,optional"`
	ErrorMessage *string      `json:"error_message,omitempty" parquet:"error_message,optional"`
	Flags        []*string    `json:"flags,omitempty" parquet:"flags,list,optional"`
	Message      *string      `json:"message,omitempty" parquet:"message,optional"`
}
