package ocsf

import "github.com/apache/arrow-go/v18/arrow"

// LongString represents a large string value in the OCSF schema
type LongString struct {
	IsTruncated     *bool   `json:"is_truncated,omitempty" parquet:"is_truncated,optional"`
	UntruncatedSize *int64  `json:"untruncated_size,omitempty" parquet:"untruncated_size,optional"`
	Value           *string `json:"value,omitempty" parquet:"value,optional"`
}

var LongStringFields = []arrow.Field{
	{Name: "is_truncated", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "untruncated_size", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var LongStringStruct = arrow.StructOf(LongStringFields...)
var LongStringClassname = "long_string"
