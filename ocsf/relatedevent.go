package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// RelatedEventFields defines the Arrow fields for RelatedEvent.
var RelatedEventFields = []arrow.Field{
	{Name: "attacks", Type: arrow.ListOf(arrow.StructOf(MITREATTCKFields...))},
	{Name: "kill_chain", Type: arrow.ListOf(arrow.StructOf(KillChainPhaseFields...))},
	{Name: "observables", Type: arrow.ListOf(arrow.StructOf(ObservableFields...))},
	{Name: "product_uid", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int64},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// RelatedEventSchema is the Arrow schema for RelatedEvent.
var RelatedEventSchema = arrow.NewSchema(RelatedEventFields, nil)

// RelatedEvent represents a related event.
type RelatedEvent struct {
	Attacks     []MITREATTCK     `json:"attacks,omitempty"`
	KillChain   []KillChainPhase `json:"kill_chain,omitempty"`
	Observables []Observable     `json:"observables,omitempty"`
	ProductUID  *string          `json:"product_uid,omitempty"`
	Type        *string          `json:"type,omitempty"`
	TypeUID     *int64           `json:"type_uid,omitempty"`
	UID         string           `json:"uid"` // required field
}

// WriteToParquet writes the RelatedEvent fields to the provided Arrow StructBuilder.
func (r *RelatedEvent) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: attacks (list of MITREATTCK).
	attacksB := sb.FieldBuilder(0).(*array.ListBuilder)
	if len(r.Attacks) > 0 {
		attacksB.Append(true)
		attacksValB := attacksB.ValueBuilder().(*array.StructBuilder)
		for _, a := range r.Attacks {
			attacksValB.Append(true)
			a.WriteToParquet(attacksValB)
		}
	} else {
		attacksB.AppendNull()
	}

	// Field 1: kill_chain (list of KillChainPhase).
	killChainB := sb.FieldBuilder(1).(*array.ListBuilder)
	if len(r.KillChain) > 0 {
		killChainB.Append(true)
		killChainValB := killChainB.ValueBuilder().(*array.StructBuilder)
		for _, kc := range r.KillChain {
			killChainValB.Append(true)
			kc.WriteToParquet(killChainValB)
		}
	} else {
		killChainB.AppendNull()
	}

	// Field 2: observables (list of Observable).
	observablesB := sb.FieldBuilder(2).(*array.ListBuilder)
	if len(r.Observables) > 0 {
		observablesB.Append(true)
		observablesValB := observablesB.ValueBuilder().(*array.StructBuilder)
		for _, o := range r.Observables {
			observablesValB.Append(true)
			o.WriteToParquet(observablesValB)
		}
	} else {
		observablesB.AppendNull()
	}

	// Field 3: product_uid.
	productUIDB := sb.FieldBuilder(3).(*array.StringBuilder)
	if r.ProductUID != nil {
		productUIDB.Append(*r.ProductUID)
	} else {
		productUIDB.AppendNull()
	}

	// Field 4: type.
	typeB := sb.FieldBuilder(4).(*array.StringBuilder)
	if r.Type != nil {
		typeB.Append(*r.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 5: type_uid.
	typeUIDB := sb.FieldBuilder(5).(*array.Int64Builder)
	if r.TypeUID != nil {
		typeUIDB.Append(*r.TypeUID)
	} else {
		typeUIDB.AppendNull()
	}

	// Field 6: uid (required).
	uidB := sb.FieldBuilder(6).(*array.StringBuilder)
	uidB.Append(r.UID)
}
