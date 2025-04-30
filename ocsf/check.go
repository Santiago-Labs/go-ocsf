package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// CheckFields defines the Arrow fields for Check.
var CheckFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "standards", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var CheckStruct = arrow.StructOf(CheckFields...)
var CheckClassname = "check"

type Check struct {
	Desc       *string  `json:"desc,omitempty" parquet:"desc,optional"`
	Name       *string  `json:"name,omitempty" parquet:"name,optional"`
	Severity   *string  `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID *int32   `json:"severity_id,omitempty" parquet:"severity_id,optional"`
	Standards  []string `json:"standards,omitempty" parquet:"standards,optional"`
	Status     *string  `json:"status,omitempty" parquet:"status,optional"`
	StatusID   *int32   `json:"status_id" parquet:"status_id"`
	UID        *string  `json:"uid,omitempty" parquet:"uid,optional"`
	Version    *string  `json:"version,omitempty" parquet:"version,optional"`
}
