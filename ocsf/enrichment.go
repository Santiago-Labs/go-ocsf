package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var EnrichmentFields = []arrow.Field{
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

var EnrichmentStruct = arrow.StructOf(EnrichmentFields...)
var EnrichmentClassname = "enrichment"

// Enrichment represents an enrichment element.
type Enrichment struct {
	Data     string  `json:"data" parquet:"data"` // JSON string
	Name     string  `json:"name" parquet:"name"` // JSON string
	Provider *string `json:"provider,omitempty" parquet:"provider"`
	Type     *string `json:"type,omitempty" parquet:"type"`
	Value    string  `json:"value" parquet:"value"`
}
