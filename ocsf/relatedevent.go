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
	Attacks       []*MITREATTCK     `json:"attacks,omitempty" parquet:"attacks,list,optional"`
	Count         *int32            `json:"count,omitempty" parquet:"count,optional"`
	CreatedTime   *int64            `json:"created_time,omitempty" parquet:"created_time,optional"`
	Description   *string           `json:"desc,omitempty" parquet:"desc,optional"`
	FirstSeenTime *int64            `json:"first_seen_time,omitempty" parquet:"first_seen_time,optional"`
	KillChain     []*KillChainPhase `json:"kill_chain,omitempty" parquet:"kill_chain,list,optional"`
	LastSeenTime  *int64            `json:"last_seen_time,omitempty" parquet:"last_seen_time,optional"`
	ModifiedTime  *int64            `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Observables   []*Observable     `json:"observables,omitempty" parquet:"observables,list,optional"`
	Product       *Product          `json:"product,omitempty" parquet:"product,optional"`
	Severity      *string           `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID    *int32            `json:"severity_id,omitempty" parquet:"severity_id,optional"`
	Tags          *KeyValueObject   `json:"tags,omitempty" parquet:"tags,optional"`
	Title         *string           `json:"title,omitempty" parquet:"title,optional"`
	Type          *string           `json:"type,omitempty" parquet:"type,optional"`
	TypeName      *string           `json:"type_name,omitempty" parquet:"type_name,optional"`
	TypeUID       *int64            `json:"type_uid,omitempty" parquet:"type_uid,optional"`
	UID           string            `json:"uid" parquet:"uid"` // required field
}
