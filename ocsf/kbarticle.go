package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// KBArticleFields defines the Arrow fields for KBArticle.
var KBArticleFields = []arrow.Field{
	{Name: "bulletin", Type: arrow.BinaryTypes.String},
	{Name: "classification", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "is_superseded", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "os", Type: OSStruct},
	{Name: "product", Type: ProductStruct},
	{Name: "severity", Type: arrow.BinaryTypes.String},
	{Name: "size", Type: arrow.PrimitiveTypes.Int32},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "title", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var KBArticleStruct = arrow.StructOf(KBArticleFields...)
var KBArticleClassname = "kb_article"

type KBArticle struct {
	Bulletin       *string  `json:"bulletin,omitempty" parquet:"bulletin"`
	Classification *string  `json:"classification,omitempty" parquet:"classification"`
	CreatedTime    *int     `json:"created_time,omitempty" parquet:"created_time"`
	CreatedTimeDt  *string  `json:"created_time_dt,omitempty" parquet:"created_time_dt"`
	IsSuperseded   *bool    `json:"is_superseded,omitempty" parquet:"is_superseded"`
	OS             *OS      `json:"os,omitempty" parquet:"os"`
	Product        *Product `json:"product,omitempty" parquet:"product"`
	Severity       *string  `json:"severity,omitempty" parquet:"severity"`
	Size           *int     `json:"size,omitempty" parquet:"size"`
	SrcURL         *string  `json:"src_url,omitempty" parquet:"src_url"`
	Title          *string  `json:"title,omitempty" parquet:"title"`
	UID            string   `json:"uid" parquet:"uid"` // required field
}
