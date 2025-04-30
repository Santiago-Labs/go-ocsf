package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var WindowsRegistryKeyFields = []arrow.Field{
	{Name: "is_system", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "security_descriptor", Type: arrow.BinaryTypes.String, Nullable: true},
}

var WindowsRegistryKeyStruct = arrow.StructOf(WindowsRegistryKeyFields...)
var WindowsRegistryKeyClassname = "windows_registry_key"

type WindowsRegistryKey struct {
	IsSystem           *bool   `json:"is_system,omitempty" parquet:"is_system,optional"`
	ModifiedTime       *int64  `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Path               *string `json:"path,omitempty" parquet:"path,optional"`
	SecurityDescriptor *string `json:"security_descriptor,omitempty" parquet:"security_descriptor,optional"`
}
