package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AssessmentFields defines the Arrow fields for Assessment.
var AssessmentFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "meets_criteria", Type: arrow.FixedWidthTypes.Boolean, Nullable: false},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "policy", Type: PolicyStruct, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var AssessmentStruct = arrow.StructOf(AssessmentFields...)
var AssessmentClassname = "assessment"

type Assessment struct {
	Category      *string `json:"category,omitempty" parquet:"category,optional"`
	Desc          *string `json:"desc,omitempty" parquet:"desc,optional"`
	MeetsCriteria bool    `json:"meets_criteria" parquet:"meets_criteria"`
	Name          *string `json:"name,omitempty" parquet:"name,optional"`
	Policy        *Policy `json:"policy,omitempty" parquet:"policy,optional"`
	UID           *string `json:"uid,omitempty" parquet:"uid,optional"`
}
