package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// KBArticleFields defines the Arrow fields for KBArticle.
var KBArticleFields = []arrow.Field{
	{Name: "bulletin", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "classification", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "is_superseded", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "os", Type: OSStruct, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "size", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var KBArticleStruct = arrow.StructOf(KBArticleFields...)
var KBArticleClassname = "kb_article"

type KBArticle struct {
	Bulletin       *string  `json:"bulletin,omitempty" parquet:"bulletin,optional"`
	Classification *string  `json:"classification,omitempty" parquet:"classification,optional"`
	CreatedTime    *int     `json:"created_time,omitempty" parquet:"created_time,optional"`
	CreatedTimeDt  *string  `json:"created_time_dt,omitempty" parquet:"created_time_dt,optional"`
	IsSuperseded   *bool    `json:"is_superseded,omitempty" parquet:"is_superseded,optional"`
	OS             *OS      `json:"os,omitempty" parquet:"os,optional"`
	Product        *Product `json:"product,omitempty" parquet:"product,optional"`
	Severity       *string  `json:"severity,omitempty" parquet:"severity,optional"`
	Size           *int     `json:"size,omitempty" parquet:"size,optional"`
	SrcURL         *string  `json:"src_url,omitempty" parquet:"src_url,optional"`
	Title          *string  `json:"title,omitempty" parquet:"title,optional"`
	UID            string   `json:"uid" parquet:"uid"` // required field
}
