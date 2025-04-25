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
	Code         *int32       `json:"code" parquet:"code,optional" ch:"code"`
	Containers   []*Container `json:"containers" parquet:"containers,list,optional" ch:"containers"`
	Data         *string      `json:"data" parquet:"data,optional" ch:"data"`
	Error        *string      `json:"error" parquet:"error,optional" ch:"error"`
	ErrorMessage *string      `json:"error_message" parquet:"error_message,optional" ch:"error_message"`
	Flags        []string     `json:"flags" parquet:"flags,list,optional" ch:"flags"`
	Message      *string      `json:"message" parquet:"message,optional" ch:"message"`
}
