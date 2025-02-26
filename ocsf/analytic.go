package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// baseAnalyticFields defines the non‐recursive fields for Analytic.
var baseAnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// AnalyticFields defines the Arrow fields for Analytic.
// To avoid infinite recursion in the "related_analytics" field,
// we include only the base (non‐recursive) fields.
var AnalyticFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "related_analytics", Type: arrow.ListOf(arrow.StructOf(baseAnalyticFields...))},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// AnalyticSchema is the Arrow schema for Analytic.
var AnalyticSchema = arrow.NewSchema(AnalyticFields, nil)

type Analytic struct {
	Category         *string    `json:"category,omitempty"`
	Desc             *string    `json:"desc,omitempty"`
	Name             *string    `json:"name,omitempty"`
	RelatedAnalytics []Analytic `json:"related_analytics,omitempty"`
	Type             *string    `json:"type,omitempty"`
	TypeID           string     `json:"type_id"`
	UID              *string    `json:"uid,omitempty"`
	Version          *string    `json:"version,omitempty"`
}

func (a *Analytic) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Category.
	catB := sb.FieldBuilder(0).(*array.StringBuilder)
	if a.Category != nil {
		catB.Append(*a.Category)
	} else {
		catB.AppendNull()
	}

	// Field 1: Desc.
	descB := sb.FieldBuilder(1).(*array.StringBuilder)
	if a.Desc != nil {
		descB.Append(*a.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 2: Name.
	nameB := sb.FieldBuilder(2).(*array.StringBuilder)
	if a.Name != nil {
		nameB.Append(*a.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 3: RelatedAnalytics (list of Analytic).
	relatedB := sb.FieldBuilder(3).(*array.ListBuilder)
	relatedValB := relatedB.ValueBuilder().(*array.StructBuilder)
	if len(a.RelatedAnalytics) > 0 {
		relatedB.Append(true)
		for _, r := range a.RelatedAnalytics {
			relatedValB.Append(true)
			r.WriteToParquet(relatedValB)
		}
	} else {
		relatedB.AppendNull()
	}

	// Field 4: Type.
	typeB := sb.FieldBuilder(4).(*array.StringBuilder)
	if a.Type != nil {
		typeB.Append(*a.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 5: TypeID.
	typeIDB := sb.FieldBuilder(5).(*array.StringBuilder)
	typeIDB.Append(a.TypeID)

	// Field 6: UID.
	uidB := sb.FieldBuilder(6).(*array.StringBuilder)
	if a.UID != nil {
		uidB.Append(*a.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 7: Version.
	versionB := sb.FieldBuilder(7).(*array.StringBuilder)
	if a.Version != nil {
		versionB.Append(*a.Version)
	} else {
		versionB.AppendNull()
	}
}
