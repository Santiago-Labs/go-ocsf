package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var FindingInfoFields = []arrow.Field{
	{Name: "analytic", Type: AnalyticStruct, Nullable: true},
	{Name: "attacks", Type: arrow.ListOf(MITREATTCKStruct), Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "data_sources", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "first_seen_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "kill_chain", Type: arrow.ListOf(KillChainPhaseStruct), Nullable: true},
	{Name: "last_seen_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "product_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "related_analytics", Type: arrow.ListOf(AnalyticStruct), Nullable: true},
	{Name: "related_events", Type: arrow.ListOf(RelatedEventStruct), Nullable: true},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "title", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "types", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var FindingInfoStruct = arrow.StructOf(FindingInfoFields...)
var FindingInfoClassname = "finding_info"

type FindingInfo struct {
	Analytic           *Analytic         `json:"analytic,omitempty" parquet:"analytic,optional"`
	Attacks            []*MITREATTCK     `json:"attacks,omitempty" parquet:"attacks,list,optional"`
	CreatedTime        *int64            `json:"created_time,omitempty" parquet:"created_time,optional"`
	DataSources        []string          `json:"data_sources,omitempty" parquet:"data_sources,list,optional"`
	Desc               *string           `json:"desc,omitempty" parquet:"desc,optional"`
	FirstSeenTime      *int64            `json:"first_seen_time,omitempty" parquet:"first_seen_time,optional"`
	KillChain          []*KillChainPhase `json:"kill_chain,omitempty" parquet:"kill_chain,list,optional"`
	LastSeenTime       *int64            `json:"last_seen_time,omitempty" parquet:"last_seen_time,optional"`
	ModifiedTime       *int64            `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Product            *Product          `json:"product,omitempty" parquet:"product,optional"`
	RelatedAnalytics   []*Analytic       `json:"related_analytics,omitempty" parquet:"related_analytics,list,optional"`
	RelatedEvents      []*RelatedEvent   `json:"related_events,omitempty" parquet:"related_events,list,optional"`
	RelatedEventsCount *int32            `json:"related_events_count,omitempty" parquet:"related_events_count,optional"`
	SrcURL             *string           `json:"src_url,omitempty" parquet:"src_url,optional"`
	Tags               *KeyValueObject   `json:"tags,omitempty" parquet:"tags,optional"`
	Title              string            `json:"title" parquet:"title"`
	Types              []string          `json:"types,omitempty" parquet:"types,list,optional"`
	UID                string            `json:"uid" parquet:"uid"`
	UIDAlt             *string           `json:"uid_alt,omitempty" parquet:"uid_alt,optional"`
}
