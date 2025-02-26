package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// KBArticleFields defines the Arrow fields for KBArticle.
var KBArticleFields = []arrow.Field{
	{Name: "bulletin", Type: arrow.BinaryTypes.String},
	{Name: "classification", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "is_superseded", Type: arrow.FixedWidthTypes.Boolean},
	// Field for OS; assume OSSchema is defined in the OS implementation.
	{Name: "os", Type: arrow.StructOf(OSFields...)},
	// Field for Product; assume ProductSchema is defined in the Product implementation.
	{Name: "product", Type: arrow.StructOf(ProductFields...)},
	{Name: "severity", Type: arrow.BinaryTypes.String},
	{Name: "size", Type: arrow.PrimitiveTypes.Int32},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "title", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// KBArticleSchema is the Arrow schema for KBArticle.
var KBArticleSchema = arrow.NewSchema(KBArticleFields, nil)

// KBArticle represents a knowledgebase article.
type KBArticle struct {
	Bulletin       *string  `json:"bulletin,omitempty"`
	Classification *string  `json:"classification,omitempty"`
	CreatedTime    *int     `json:"created_time,omitempty"`
	CreatedTimeDt  *string  `json:"created_time_dt,omitempty"`
	IsSuperseded   *bool    `json:"is_superseded,omitempty"`
	OS             *OS      `json:"os,omitempty"`
	Product        *Product `json:"product,omitempty"`
	Severity       *string  `json:"severity,omitempty"`
	Size           *int     `json:"size,omitempty"`
	SrcURL         *string  `json:"src_url,omitempty"`
	Title          *string  `json:"title,omitempty"`
	UID            string   `json:"uid"` // required field
}

// WriteToParquet writes the KBArticle fields to the provided Arrow StructBuilder.
func (k *KBArticle) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: Bulletin.
	bulletinB := sb.FieldBuilder(0).(*array.StringBuilder)
	if k.Bulletin != nil {
		bulletinB.Append(*k.Bulletin)
	} else {
		bulletinB.AppendNull()
	}

	// Field 1: Classification.
	classificationB := sb.FieldBuilder(1).(*array.StringBuilder)
	if k.Classification != nil {
		classificationB.Append(*k.Classification)
	} else {
		classificationB.AppendNull()
	}

	// Field 2: CreatedTime.
	createdTimeB := sb.FieldBuilder(2).(*array.Int32Builder)
	if k.CreatedTime != nil {
		createdTimeB.Append(int32(*k.CreatedTime))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 3: CreatedTimeDt.
	createdTimeDtB := sb.FieldBuilder(3).(*array.StringBuilder)
	if k.CreatedTimeDt != nil {
		createdTimeDtB.Append(*k.CreatedTimeDt)
	} else {
		createdTimeDtB.AppendNull()
	}

	// Field 4: IsSuperseded.
	isSupersededB := sb.FieldBuilder(4).(*array.BooleanBuilder)
	if k.IsSuperseded != nil {
		isSupersededB.Append(*k.IsSuperseded)
	} else {
		isSupersededB.AppendNull()
	}

	// Field 5: OS (nested struct).
	osB := sb.FieldBuilder(5).(*array.StructBuilder)
	if k.OS != nil {
		k.OS.WriteToParquet(osB)
	} else {
		osB.AppendNull()
	}

	// Field 6: Product (nested struct).
	productB := sb.FieldBuilder(6).(*array.StructBuilder)
	if k.Product != nil {
		productB.Append(true)
		k.Product.WriteToParquet(productB)
	} else {
		productB.AppendNull()
	}

	// Field 7: Severity.
	severityB := sb.FieldBuilder(7).(*array.StringBuilder)
	if k.Severity != nil {
		severityB.Append(*k.Severity)
	} else {
		severityB.AppendNull()
	}

	// Field 8: Size.
	sizeB := sb.FieldBuilder(8).(*array.Int32Builder)
	if k.Size != nil {
		sizeB.Append(int32(*k.Size))
	} else {
		sizeB.AppendNull()
	}

	// Field 9: SrcURL.
	srcURLB := sb.FieldBuilder(9).(*array.StringBuilder)
	if k.SrcURL != nil {
		srcURLB.Append(*k.SrcURL)
	} else {
		srcURLB.AppendNull()
	}

	// Field 10: Title.
	titleB := sb.FieldBuilder(10).(*array.StringBuilder)
	if k.Title != nil {
		titleB.Append(*k.Title)
	} else {
		titleB.AppendNull()
	}

	// Field 11: UID (required).
	uidB := sb.FieldBuilder(11).(*array.StringBuilder)
	uidB.Append(k.UID)
}
