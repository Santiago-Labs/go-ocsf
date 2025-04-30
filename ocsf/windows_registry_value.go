package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var WindowsRegistryValueFields = []arrow.Field{
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "is_default", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_system", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
}

var WindowsRegistryValueStruct = arrow.StructOf(WindowsRegistryValueFields...)
var WindowsRegistryValueClassname = "windows_registry_value"

type WindowsRegistryValue struct {
	Data         *string `json:"data,omitempty" parquet:"data,optional"`
	IsDefault    *bool   `json:"is_default,omitempty" parquet:"is_default,optional"`
	IsSystem     *bool   `json:"is_system,omitempty" parquet:"is_system,optional"`
	ModifiedTime *int64  `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Name         *string `json:"name,omitempty" parquet:"name,optional"`
	Path         *string `json:"path,omitempty" parquet:"path,optional"`
	Type         *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID       *int32  `json:"type_id,omitempty" parquet:"type_id,optional"`
}
