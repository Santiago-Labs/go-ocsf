package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
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
	Build         *string `json:"build,omitempty" parquet:"build"`
	Country       *string `json:"country,omitempty" parquet:"country"`
	CPEName       *string `json:"cpe_name,omitempty" parquet:"cpe_name"`
	CPUBits       *int    `json:"cpu_bits,omitempty" parquet:"cpu_bits"`
	Edition       *string `json:"edition,omitempty" parquet:"edition"`
	KernelRelease *string `json:"kernel_release,omitempty" parquet:"kernel_release"`
	Lang          *string `json:"lang,omitempty" parquet:"lang"`
	Name          string  `json:"name" parquet:"name"`
	SPName        *string `json:"sp_name,omitempty" parquet:"sp_name"`
	SPVer         *int    `json:"sp_ver,omitempty" parquet:"sp_ver"`
	Type          *string `json:"type,omitempty" parquet:"type"`
	TypeID        int     `json:"type_id" parquet:"type_id"`
	Version       *string `json:"version,omitempty" parquet:"version"`
}
