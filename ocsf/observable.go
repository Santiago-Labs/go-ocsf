package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// Observable represents an observable entity.
type Observable struct {
	Name       string      `json:"name"`
	Reputation *Reputation `json:"reputation,omitempty"`
	Type       *string     `json:"type,omitempty"`
	TypeID     int         `json:"type_id"`
	Value      *string     `json:"value,omitempty"`
}

// ObservableFields defines the Arrow fields for Observable in the desired order.
var ObservableFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "reputation", Type: arrow.StructOf(ReputationFields...)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

// ObservableSchema is the Arrow schema for Observable.
var ObservableSchema = arrow.NewSchema(ObservableFields, nil)

// WriteToParquet writes the Observable fields to the provided Arrow StructBuilder.
func (o *Observable) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	nameB.Append(o.Name)

	// Field 1: Reputation (nested struct).
	repB := sb.FieldBuilder(1).(*array.StructBuilder)
	if o.Reputation != nil {
		repB.Append(true)
		o.Reputation.WriteToParquet(repB)
	} else {
		repB.AppendNull()
	}

	// Field 2: Type.
	typeB := sb.FieldBuilder(2).(*array.StringBuilder)
	if o.Type != nil {
		typeB.Append(*o.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 3: TypeID.
	typeIDB := sb.FieldBuilder(3).(*array.Int32Builder)
	typeIDB.Append(int32(o.TypeID))

	// Field 4: Value.
	valueB := sb.FieldBuilder(4).(*array.StringBuilder)
	if o.Value != nil {
		valueB.Append(*o.Value)
	} else {
		valueB.AppendNull()
	}
}
