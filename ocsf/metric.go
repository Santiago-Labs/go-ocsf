package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// MetricFields defines the Arrow fields for Metric.
var MetricFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: true},
}

var MetricStruct = arrow.StructOf(MetricFields...)
var MetricClassname = "metric"

type Metric struct {
	Name  string  `json:"name" parquet:"name"`
	Value *string `json:"value" parquet:"value"`
}
