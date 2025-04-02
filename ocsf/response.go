package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// ResponseFields defines the Arrow fields for Response.
var ResponseFields = []arrow.Field{
	{Name: "code", Type: arrow.PrimitiveTypes.Int32},
	{Name: "containers", Type: arrow.ListOf(ContainerStruct)},
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "error", Type: arrow.BinaryTypes.String},
	{Name: "error_message", Type: arrow.BinaryTypes.String},
	{Name: "flags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "message", Type: arrow.BinaryTypes.String},
}

var ResponseStruct = arrow.StructOf(ResponseFields...)

type Response struct {
	Code         *int32       `json:"code,omitempty" parquet:"code"`
	Containers   []*Container `json:"containers,omitempty" parquet:"containers"`
	Data         *string      `json:"data,omitempty" parquet:"data"`
	Error        *string      `json:"error,omitempty" parquet:"error"`
	ErrorMessage *string      `json:"error_message,omitempty" parquet:"error_message"`
	Flags        []string     `json:"flags,omitempty" parquet:"flags"`
	Message      *string      `json:"message,omitempty" parquet:"message"`
}
