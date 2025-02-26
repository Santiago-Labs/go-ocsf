package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// MetricFields defines the Arrow fields for Metric.
var MetricFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

// MetricSchema is the Arrow schema for Metric.
var MetricSchema = arrow.NewSchema(MetricFields, nil)

// Metric represents a metric with a name and a value.
type Metric struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// WriteToParquet writes the Metric fields to the provided Arrow StructBuilder.
func (m *Metric) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(m.Name)

	// Field 1: Value.
	valueB := sb.FieldBuilder(1).(*array.StringBuilder)
	valueB.Append(m.Value)
}
