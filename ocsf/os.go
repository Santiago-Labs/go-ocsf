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
	Build         *string `json:"build" parquet:"build,optional" ch:"build" ch:"build"`
	Country       *string `json:"country" parquet:"country,optional" ch:"country"`
	CPEName       *string `json:"cpe_name" parquet:"cpe_name,optional" ch:"cpe_name"`
	CPUBits       *int64  `json:"cpu_bits" parquet:"cpu_bits,optional" ch:"cpu_bits"`
	Edition       *string `json:"edition" parquet:"edition,optional" ch:"edition"`
	KernelRelease *string `json:"kernel_release" parquet:"kernel_release,optional" ch:"kernel_release"`
	Lang          *string `json:"lang" parquet:"lang,optional" ch:"lang"`
	Name          string  `json:"name" parquet:"name" ch:"name"`
	SPName        *string `json:"sp_name" parquet:"sp_name,optional" ch:"sp_name"`
	SPVer         *int64  `json:"sp_ver" parquet:"sp_ver,optional" ch:"sp_ver"`
	Type          *string `json:"type" parquet:"type,optional" ch:"type"`
	TypeID        int     `json:"type_id" parquet:"type_id" ch:"type_id"`
	Version       *string `json:"version" parquet:"version,optional" ch:"version"`
}
