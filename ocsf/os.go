package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// OSFields defines the Arrow fields for OS.
var OSFields = []arrow.Field{
	{Name: "build", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "country", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cpe_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cpu_bits", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "edition", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "kernel_release", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "lang", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "sp_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "sp_ver", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var OSStruct = arrow.StructOf(OSFields...)
var OSClassname = "os"

type OS struct {
	Build         *string `json:"build,omitempty" parquet:"build,optional"`
	Country       *string `json:"country,omitempty" parquet:"country,optional"`
	CPEName       *string `json:"cpe_name,omitempty" parquet:"cpe_name,optional"`
	CPUBits       *int    `json:"cpu_bits,omitempty" parquet:"cpu_bits,optional"`
	Edition       *string `json:"edition,omitempty" parquet:"edition,optional"`
	KernelRelease *string `json:"kernel_release,omitempty" parquet:"kernel_release,optional"`
	Lang          *string `json:"lang,omitempty" parquet:"lang,optional"`
	Name          string  `json:"name" parquet:"name"`
	SPName        *string `json:"sp_name,omitempty" parquet:"sp_name,optional"`
	SPVer         *int    `json:"sp_ver,omitempty" parquet:"sp_ver,optional"`
	Type          *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID        int     `json:"type_id" parquet:"type_id"`
	Version       *string `json:"version,omitempty" parquet:"version,optional"`
}
