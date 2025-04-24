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
	Bulletin       *string  `json:"bulletin,omitempty" parquet:"bulletin,optional" ch:"bulletin" ch:"bulletin"`
	Classification *string  `json:"classification,omitempty" parquet:"classification,optional" ch:"classification"`
	CreatedTime    *int     `json:"created_time,omitempty" parquet:"created_time,optional" ch:"created_time"`
	CreatedTimeDt  *string  `json:"created_time_dt,omitempty" parquet:"created_time_dt,optional" ch:"created_time_dt"`
	IsSuperseded   *bool    `json:"is_superseded,omitempty" parquet:"is_superseded,optional" ch:"is_superseded"`
	OS             *OS      `json:"os,omitempty" parquet:"os,optional" ch:"os"`
	Product        *Product `json:"product,omitempty" parquet:"product,optional" ch:"product"`
	Severity       *string  `json:"severity,omitempty" parquet:"severity,optional" ch:"severity"`
	Size           *int     `json:"size,omitempty" parquet:"size,optional" ch:"size"`
	SrcURL         *string  `json:"src_url,omitempty" parquet:"src_url,optional" ch:"src_url"`
	Title          *string  `json:"title,omitempty" parquet:"title,optional" ch:"title"`
	UID            string   `json:"uid" parquet:"uid" ch:"uid"` // required field
}
