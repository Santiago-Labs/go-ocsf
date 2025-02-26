package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// RemediationFields defines the Arrow fields for Remediation.
var RemediationFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "kb_article_list", Type: arrow.ListOf(arrow.StructOf(KBArticleFields...))},
	{Name: "references", Type: arrow.ListOf(arrow.BinaryTypes.String)},
}

// RemediationSchema is the Arrow schema for Remediation.
var RemediationSchema = arrow.NewSchema(RemediationFields, nil)

// Remediation represents remediation guidance.
type Remediation struct {
	Description   string      `json:"desc"`                 // required field
	KbArticleList []KBArticle `json:"kb_article_list"`      // list of KBArticle
	References    []string    `json:"references,omitempty"` // optional list of strings
}

// WriteToParquet writes the Remediation fields to the provided Arrow StructBuilder.
func (r *Remediation) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Description.
	descB := sb.FieldBuilder(0).(*array.StringBuilder)
	descB.Append(r.Description)

	// Field 1: KbArticleList (list of KBArticle).
	kbArticlesB := sb.FieldBuilder(1).(*array.ListBuilder)
	if len(r.KbArticleList) > 0 {
		kbArticlesB.Append(true)
		kbArticlesValB := kbArticlesB.ValueBuilder().(*array.StructBuilder)
		for _, kb := range r.KbArticleList {
			kbArticlesValB.Append(true)
			kb.WriteToParquet(kbArticlesValB)
		}
	} else {
		kbArticlesB.AppendNull()
	}

	// Field 2: References (list of strings).
	referencesB := sb.FieldBuilder(2).(*array.ListBuilder)
	if len(r.References) > 0 {
		referencesB.Append(true)
		referencesValB := referencesB.ValueBuilder().(*array.StringBuilder)
		for _, ref := range r.References {
			referencesValB.Append(ref)
		}
	} else {
		referencesB.AppendNull()
	}
}
