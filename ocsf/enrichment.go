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
	CreatedTime *int64      `json:"created_time,omitempty" parquet:"created_time,optional"`
	Data        string      `json:"data" parquet:"data"` // JSON string
	Desc        *string     `json:"desc,omitempty" parquet:"desc,optional"`
	Name        string      `json:"name" parquet:"name"` // JSON string
	Reputation  *Reputation `json:"reputation,omitempty" parquet:"reputation,optional"`
	Provider    *string     `json:"provider,omitempty" parquet:"provider,optional"`
	ShortDesc   *string     `json:"short_desc,omitempty" parquet:"short_desc,optional"`
	SrcURL      *string     `json:"src_url,omitempty" parquet:"src_url,optional"`
	Type        *string     `json:"type,omitempty" parquet:"type,optional"`
	Value       string      `json:"value" parquet:"value"`
}
