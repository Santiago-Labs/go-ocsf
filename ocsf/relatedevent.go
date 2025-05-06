package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// RelatedEventFields defines the Arrow fields for RelatedEvent.
var RelatedEventFields = []arrow.Field{
	{Name: "attacks", Type: arrow.ListOf(MITREATTCKStruct), Nullable: true},
	{Name: "kill_chain", Type: arrow.ListOf(KillChainPhaseStruct), Nullable: true},
	{Name: "observables", Type: arrow.ListOf(ObservableStruct), Nullable: true},
	{Name: "product_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var RelatedEventStruct = arrow.StructOf(RelatedEventFields...)
var RelatedEventClassname = "related_event"

type RelatedEvent struct {
	Attacks     []*MITREATTCK     `json:"attacks,omitempty" parquet:"attacks,list,optional"`
	KillChain   []*KillChainPhase `json:"kill_chain,omitempty" parquet:"kill_chain,list,optional"`
	Observables []*Observable     `json:"observables,omitempty" parquet:"observables,list,optional"`
	ProductUID  *string           `json:"product_uid,omitempty" parquet:"product_uid,optional"`
	Type        *string           `json:"type,omitempty" parquet:"type,optional"`
	TypeUID     *int64            `json:"type_uid,omitempty" parquet:"type_uid,optional"`
	UID         string            `json:"uid" parquet:"uid"` // required field
}
