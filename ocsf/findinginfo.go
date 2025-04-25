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
	Analytic         *Analytic         `json:"analytic" parquet:"analytic,optional" ch:"analytic" ch:"analytic"`
	Attacks          []*MITREATTCK     `json:"attacks" parquet:"attacks,list,optional" ch:"attacks"`
	CreatedTime      *int64            `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	DataSources      []string          `json:"data_sources" parquet:"data_sources,list,optional" ch:"data_sources"`
	Desc             *string           `json:"desc" parquet:"desc,optional" ch:"desc"`
	FirstSeenTime    *int64            `json:"first_seen_time" parquet:"first_seen_time,optional" ch:"first_seen_time"`
	KillChain        []*KillChainPhase `json:"kill_chain" parquet:"kill_chain,list,optional" ch:"kill_chain"`
	LastSeenTime     *int64            `json:"last_seen_time" parquet:"last_seen_time,optional" ch:"last_seen_time"`
	ModifiedTime     *int64            `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	ProductUID       *string           `json:"product_uid" parquet:"product_uid,optional" ch:"product_uid"`
	RelatedAnalytics []*Analytic       `json:"related_analytics" parquet:"related_analytics,list,optional" ch:"related_analytics"`
	RelatedEvents    []*RelatedEvent   `json:"related_events" parquet:"related_events,list,optional" ch:"related_events"`
	SrcURL           *string           `json:"src_url" parquet:"src_url,optional" ch:"src_url"`
	Title            string            `json:"title" parquet:"title" ch:"title"`
	Types            []string          `json:"types" parquet:"types,list,optional" ch:"types"`
	UID              string            `json:"uid" parquet:"uid" ch:"uid"`
}
