package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// RelatedEventFields defines the Arrow fields for RelatedEvent.
var RelatedEventFields = []arrow.Field{
	{Name: "attacks", Type: arrow.ListOf(MITREATTCKStruct)},
	{Name: "kill_chain", Type: arrow.ListOf(KillChainPhaseStruct)},
	{Name: "observables", Type: arrow.ListOf(ObservableStruct)},
	{Name: "product_uid", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int64},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var RelatedEventStruct = arrow.StructOf(RelatedEventFields...)

type RelatedEvent struct {
	Attacks     []MITREATTCK     `json:"attacks,omitempty" parquet:"attacks"`
	KillChain   []KillChainPhase `json:"kill_chain,omitempty" parquet:"kill_chain"`
	Observables []Observable     `json:"observables,omitempty" parquet:"observables"`
	ProductUID  *string          `json:"product_uid,omitempty" parquet:"product_uid"`
	Type        *string          `json:"type,omitempty" parquet:"type"`
	TypeUID     *int64           `json:"type_uid,omitempty" parquet:"type_uid"`
	UID         string           `json:"uid" parquet:"uid"` // required field
}
