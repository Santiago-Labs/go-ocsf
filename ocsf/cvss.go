package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// CVSSFields defines the Arrow fields for the CVSS type.
var CVSSFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
	{Name: "depth", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "metrics", Type: arrow.ListOf(MetricStruct), Nullable: true},
	{Name: "overall_score", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vector_string", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var CVSSStruct = arrow.StructOf(CVSSFields...)
var CVSSClassname = "cvss"

type CVSS struct {
	BaseScore    float64   `json:"base_score" parquet:"base_score" ch:"base_score"`
	Depth        *string   `json:"depth,omitempty" parquet:"depth,optional" ch:"depth"`
	Metrics      []*Metric `json:"metrics,omitempty" parquet:"metrics,list,optional" ch:"metrics"`
	OverallScore *float64  `json:"overall_score,omitempty" parquet:"overall_score,optional" ch:"overall_score"`
	Severity     *string   `json:"severity,omitempty" parquet:"severity,optional" ch:"severity"`
	VectorString *string   `json:"vector_string,omitempty" parquet:"vector_string,optional" ch:"vector_string"`
	Version      string    `json:"version" parquet:"version" ch:"version"`
}
