package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var EnrichmentFields = []arrow.Field{
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

var EnrichmentSchema = arrow.NewSchema(EnrichmentFields, nil)

// Enrichment represents an enrichment element.
type Enrichment struct {
	Data     string  `json:"data"` // JSON string
	Name     string  `json:"name"` // JSON string
	Provider *string `json:"provider,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    string  `json:"value"`
}

// WriteToParquet writes the Enrichment fields to the provided Arrow StructBuilder.
func (e *Enrichment) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Data.
	dataB := sb.FieldBuilder(0).(*array.StringBuilder)
	dataB.Append(e.Data)

	// Field 1: Name.
	nameB := sb.FieldBuilder(1).(*array.StringBuilder)
	nameB.Append(e.Name)

	// Field 2: Provider.
	providerB := sb.FieldBuilder(2).(*array.StringBuilder)
	if e.Provider != nil {
		providerB.Append(*e.Provider)
	} else {
		providerB.AppendNull()
	}

	// Field 3: Type.
	typeB := sb.FieldBuilder(3).(*array.StringBuilder)
	if e.Type != nil {
		typeB.Append(*e.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 4: Value.
	valueB := sb.FieldBuilder(4).(*array.StringBuilder)
	valueB.Append(e.Value)
}
