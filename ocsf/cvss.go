package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// CVSSFields defines the Arrow fields for the CVSS type.
var CVSSFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "depth", Type: arrow.BinaryTypes.String},
	{Name: "metrics", Type: arrow.ListOf(arrow.StructOf(MetricFields...))},
	{Name: "overall_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "severity", Type: arrow.BinaryTypes.String},
	{Name: "vector_string", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// CVSSSchema is the Arrow schema for CVSS.
var CVSSSchema = arrow.NewSchema(CVSSFields, nil)

type CVSS struct {
	BaseScore    float64  `json:"base_score"`
	Depth        *string  `json:"depth,omitempty"`
	Metrics      []Metric `json:"metrics,omitempty"`
	OverallScore *float64 `json:"overall_score,omitempty"`
	Severity     *string  `json:"severity,omitempty"`
	VectorString *string  `json:"vector_string,omitempty"`
	Version      string   `json:"version"`
}

func (c *CVSS) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: BaseScore.
	baseScoreB := sb.FieldBuilder(0).(*array.Float64Builder)
	baseScoreB.Append(c.BaseScore)

	// Field 1: Depth.
	depthB := sb.FieldBuilder(1).(*array.StringBuilder)
	if c.Depth != nil {
		depthB.Append(*c.Depth)
	} else {
		depthB.AppendNull()
	}

	// Field 2: Metrics - for simplicity, we marshal Metrics slice to JSON.
	metricsB := sb.FieldBuilder(2).(*array.ListBuilder)
	metricsBIfValB := metricsB.ValueBuilder().(*array.StructBuilder)
	if len(c.Metrics) > 0 {
		metricsB.Append(true)
		for _, m := range c.Metrics {
			metricsBIfValB.Append(true)
			m.WriteToParquet(metricsBIfValB)
		}
	} else {
		metricsB.AppendNull()
	}

	// Field 3: OverallScore.
	overallB := sb.FieldBuilder(3).(*array.Float64Builder)
	if c.OverallScore != nil {
		overallB.Append(*c.OverallScore)
	} else {
		overallB.AppendNull()
	}

	// Field 4: Severity.
	severityB := sb.FieldBuilder(4).(*array.StringBuilder)
	if c.Severity != nil {
		severityB.Append(*c.Severity)
	} else {
		severityB.AppendNull()
	}

	// Field 5: VectorString.
	vectorB := sb.FieldBuilder(5).(*array.StringBuilder)
	if c.VectorString != nil {
		vectorB.Append(*c.VectorString)
	} else {
		vectorB.AppendNull()
	}

	// Field 6: Version.
	versionB := sb.FieldBuilder(6).(*array.StringBuilder)
	versionB.Append(c.Version)
}
