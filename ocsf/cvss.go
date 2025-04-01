package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// CVSSFields defines the Arrow fields for the CVSS type.
var CVSSFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "depth", Type: arrow.BinaryTypes.String},
	{Name: "metrics", Type: arrow.ListOf(MetricStruct)},
	{Name: "overall_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "severity", Type: arrow.BinaryTypes.String},
	{Name: "vector_string", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var CVSSStruct = arrow.StructOf(CVSSFields...)

// CVSSSchema is the Arrow schema for CVSS.
var CVSSSchema = arrow.NewSchema(CVSSFields, nil)

type CVSS struct {
	BaseScore    float64  `json:"base_score" parquet:"base_score"`
	Depth        *string  `json:"depth,omitempty" parquet:"depth"`
	Metrics      []Metric `json:"metrics,omitempty" parquet:"metrics"`
	OverallScore *float64 `json:"overall_score,omitempty" parquet:"overall_score"`
	Severity     *string  `json:"severity,omitempty" parquet:"severity"`
	VectorString *string  `json:"vector_string,omitempty" parquet:"vector_string"`
	Version      string   `json:"version" parquet:"version"`
}
