package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AffectedCodeFields defines the Arrow fields for AffectedCode.
var AffectedCodeFields = []arrow.Field{
	{Name: "end_line", Type: arrow.PrimitiveTypes.Int32},
	{Name: "start_line", Type: arrow.PrimitiveTypes.Int32},
	{Name: "file", Type: arrow.StructOf(FileFields...)},
	{Name: "owner", Type: arrow.StructOf(UserFields...)},
	{Name: "remediation", Type: arrow.StructOf(RemediationFields...)},
}

// AffectedCodeSchema is the Arrow schema for AffectedCode.
var AffectedCodeSchema = arrow.NewSchema(AffectedCodeFields, nil)

type AffectedCode struct {
	EndLine     int32        `json:"end_line" parquet:"end_line"`
	StartLine   int32        `json:"start_line" parquet:"start_line"`
	File        File         `json:"file" parquet:"file"`
	Owner       *User        `json:"owner" parquet:"owner"`
	Remediation *Remediation `json:"remediation" parquet:"remediation"`
}
