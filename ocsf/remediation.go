package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// RemediationFields defines the Arrow fields for Remediation.
var RemediationFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "kb_article_list", Type: arrow.ListOf(KBArticleStruct)},
	{Name: "references", Type: arrow.ListOf(arrow.BinaryTypes.String)},
}

var RemediationStruct = arrow.StructOf(RemediationFields...)
var RemediationClassname = "remediation"

type Remediation struct {
	Description   string      `json:"desc" parquet:"desc"`
	KbArticleList []KBArticle `json:"kb_article_list" parquet:"kb_article_list"`
	References    []string    `json:"references,omitempty" parquet:"references"`
}
