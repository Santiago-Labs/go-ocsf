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
	Attacks     []*MITREATTCK     `json:"attacks" parquet:"attacks,list,optional" ch:"attacks"`
	KillChain   []*KillChainPhase `json:"kill_chain" parquet:"kill_chain,list,optional" ch:"kill_chain"`
	Observables []*Observable     `json:"observables" parquet:"observables,list,optional" ch:"observables"`
	ProductUID  *string           `json:"product_uid" parquet:"product_uid,optional" ch:"product_uid"`
	Type        *string           `json:"type" parquet:"type,optional" ch:"type"`
	TypeUID     *int64            `json:"type_uid" parquet:"type_uid,optional" ch:"type_uid"`
	UID         string            `json:"uid" parquet:"uid" ch:"uid"` // required field
}
