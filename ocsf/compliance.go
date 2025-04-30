package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ComplianceFields defines the Arrow fields for Compliance.
var ComplianceFields = []arrow.Field{
	{Name: "assessments", Type: arrow.ListOf(AssessmentStruct), Nullable: true},
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "checks", Type: arrow.ListOf(CheckStruct), Nullable: true},
	{Name: "control", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "control_parameters", Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String), Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "requirements", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "standards", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_code", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_details", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
}

var ComplianceStruct = arrow.StructOf(ComplianceFields...)
var ComplianceClassname = "compliance"

type Compliance struct {
	Assessments       []*Assessment     `json:"assessments,omitempty" parquet:"assessments,optional"`
	Category          *string           `json:"category,omitempty" parquet:"category,optional"`
	Checks            []*Check          `json:"checks,omitempty" parquet:"checks,optional"`
	Control           *string           `json:"control,omitempty" parquet:"control,optional"`
	ControlParameters map[string]string `json:"control_parameters,omitempty" parquet:"control_parameters,optional"`
	Desc              *string           `json:"desc,omitempty" parquet:"desc,optional"`
	Requirements      []string          `json:"requirements,omitempty" parquet:"requirements,optional"`
	Standards         []string          `json:"standards,omitempty" parquet:"standards,optional"`
	Status            *string           `json:"status,omitempty" parquet:"status,optional"`
	StatusCode        *string           `json:"status_code,omitempty" parquet:"status_code,optional"`
	StatusDetails     []string          `json:"status_details,omitempty" parquet:"status_details,optional"`
	StatusID          *int32            `json:"status_id,omitempty" parquet:"status_id,optional"`
}
