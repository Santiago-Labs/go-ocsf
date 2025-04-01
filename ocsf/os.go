package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var OSFields = []arrow.Field{
	{Name: "build", Type: arrow.BinaryTypes.String},
	{Name: "country", Type: arrow.BinaryTypes.String},
	{Name: "cpe_name", Type: arrow.BinaryTypes.String},
	{Name: "cpu_bits", Type: arrow.PrimitiveTypes.Int32},
	{Name: "edition", Type: arrow.BinaryTypes.String},
	{Name: "lang", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "sp_name", Type: arrow.BinaryTypes.String},
	{Name: "sp_ver", Type: arrow.PrimitiveTypes.Int32},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var OSStruct = arrow.StructOf(OSFields...)

type OS struct {
	Build   *string `json:"build,omitempty" parquet:"build"`
	Country *string `json:"country,omitempty" parquet:"country"`
	CpeName *string `json:"cpe_name,omitempty" parquet:"cpe_name"`
	CPUBits *int    `json:"cpu_bits,omitempty" parquet:"cpu_bits"`
	Edition *string `json:"edition,omitempty" parquet:"edition"`
	Lang    *string `json:"lang,omitempty" parquet:"lang"`
	Name    string  `json:"name" parquet:"name"`
	SPName  *string `json:"sp_name,omitempty" parquet:"sp_name"`
	SPVer   *int    `json:"sp_ver,omitempty" parquet:"sp_ver"`
	Type    *string `json:"type,omitempty" parquet:"type"`
	TypeID  int     `json:"type_id" parquet:"type_id"`
	Version *string `json:"version,omitempty" parquet:"version"`
}
