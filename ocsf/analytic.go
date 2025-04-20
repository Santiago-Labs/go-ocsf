package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// baseAnalyticFields defines the non‐recursive fields for Analytic.
var relatedAnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var RelatedAnalyticStruct = arrow.StructOf(relatedAnalyticFields...)
var RelatedAnalyticClassname = "related_analytics"

// AnalyticFields defines the Arrow fields for Analytic.
// To avoid infinite recursion in the "related_analytics" field,
// we include only the base (non‐recursive) fields.
var AnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "related_analytics", Type: arrow.ListOf(RelatedAnalyticStruct), Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var AnalyticStruct = arrow.StructOf(AnalyticFields...)
var AnalyticClassname = "analytic"

type Analytic struct {
	Category         *string            `json:"category,omitempty" parquet:"category,optional"`
	Desc             *string            `json:"desc,omitempty" parquet:"desc,optional"`
	Name             *string            `json:"name,omitempty" parquet:"name,optional"`
	RelatedAnalytics []*RelatedAnalytic `json:"related_analytics,omitempty" parquet:"related_analytics,list,optional"`
	Type             *string            `json:"type,omitempty" parquet:"type,optional"`
	TypeID           string             `json:"type_id" parquet:"type_id"`
	UID              *string            `json:"uid,omitempty" parquet:"uid,optional"`
	Version          *string            `json:"version,omitempty" parquet:"version,optional"`
}

type RelatedAnalytic struct {
	Category *string `json:"category,omitempty" parquet:"category,optional"`
	Desc     *string `json:"desc,omitempty" parquet:"desc,optional"`
	Name     *string `json:"name,omitempty" parquet:"name,optional"`
	Type     *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID   string  `json:"type_id" parquet:"type_id"`
	UID      *string `json:"uid,omitempty" parquet:"uid,optional"`
	Version  *string `json:"version,omitempty" parquet:"version,optional"`
}
