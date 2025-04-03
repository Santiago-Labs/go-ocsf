package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// baseAnalyticFields defines the non‐recursive fields for Analytic.
var relatedAnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var RelatedAnalyticStruct = arrow.StructOf(relatedAnalyticFields...)
var RelatedAnalyticClassname = "related_analytics"

// AnalyticFields defines the Arrow fields for Analytic.
// To avoid infinite recursion in the "related_analytics" field,
// we include only the base (non‐recursive) fields.
var AnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "related_analytics", Type: arrow.ListOf(RelatedAnalyticStruct)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var AnalyticStruct = arrow.StructOf(AnalyticFields...)
var AnalyticClassname = "analytic"

type Analytic struct {
	Category         *string           `json:"category,omitempty" parquet:"category"`
	Desc             *string           `json:"desc,omitempty" parquet:"desc"`
	Name             *string           `json:"name,omitempty" parquet:"name"`
	RelatedAnalytics []RelatedAnalytic `json:"related_analytics,omitempty" parquet:"related_analytics"`
	Type             *string           `json:"type,omitempty" parquet:"type"`
	TypeID           string            `json:"type_id" parquet:"type_id"`
	UID              *string           `json:"uid,omitempty" parquet:"uid"`
	Version          *string           `json:"version,omitempty" parquet:"version"`
}

type RelatedAnalytic struct {
	Category *string `json:"category,omitempty" parquet:"category"`
	Desc     *string `json:"desc,omitempty" parquet:"desc"`
	Name     *string `json:"name,omitempty" parquet:"name"`
	Type     *string `json:"type,omitempty" parquet:"type"`
	TypeID   string  `json:"type_id" parquet:"type_id"`
	UID      *string `json:"uid,omitempty" parquet:"uid"`
	Version  *string `json:"version,omitempty" parquet:"version"`
}
