package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AffectedCodeFields defines the Arrow fields for AffectedCode.
var AffectedCodeFields = []arrow.Field{
	{Name: "end_line", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "start_line", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "file", Type: FileStruct, Nullable: false},
	{Name: "owner", Type: UserStruct, Nullable: true},
	{Name: "remediation", Type: RemediationStruct, Nullable: true},
}

var AffectedCodeStruct = arrow.StructOf(AffectedCodeFields...)
var AffectedCodeClassname = "affected_code"

type AffectedCode struct {
	EndLine     int32        `json:"end_line" parquet:"end_line" ch:"end_line"`
	StartLine   int32        `json:"start_line" parquet:"start_line" ch:"start_line"`
	File        File         `json:"file" parquet:"file" ch:"file"`
	Owner       *User        `json:"owner" parquet:"owner,optional" ch:"owner,omitempty"`
	Remediation *Remediation `json:"remediation" parquet:"remediation,optional" ch:"remediation,omitempty"`
}
