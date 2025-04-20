package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var EnrichmentFields = []arrow.Field{
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var EnrichmentStruct = arrow.StructOf(EnrichmentFields...)
var EnrichmentClassname = "enrichment"

// Enrichment represents an enrichment element.
type Enrichment struct {
	Data     string  `json:"data" parquet:"data"` // JSON string
	Name     string  `json:"name" parquet:"name"` // JSON string
	Provider *string `json:"provider,omitempty" parquet:"provider,optional"`
	Type     *string `json:"type,omitempty" parquet:"type,optional"`
	Value    string  `json:"value" parquet:"value"`
}
