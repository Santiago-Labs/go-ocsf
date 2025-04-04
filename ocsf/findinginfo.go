package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

var FindingInfoFields = []arrow.Field{
	{Name: "analytic", Type: AnalyticStruct},
	{Name: "attacks", Type: arrow.ListOf(MITREATTCKStruct)},
	{Name: "created_time", Type: arrow.BinaryTypes.String},
	{Name: "data_sources", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "first_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "kill_chain", Type: arrow.ListOf(KillChainPhaseStruct)},
	{Name: "last_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "product_uid", Type: arrow.BinaryTypes.String},
	{Name: "related_analytics", Type: arrow.ListOf(AnalyticStruct)},
	{Name: "related_events", Type: arrow.ListOf(RelatedEventStruct)},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "title", Type: arrow.BinaryTypes.String},
	{Name: "types", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var FindingInfoStruct = arrow.StructOf(FindingInfoFields...)
var FindingInfoClassname = "finding_info"

type FindingInfo struct {
	Analytic         *Analytic        `json:"analytic,omitempty" parquet:"analytic"`
	Attacks          []MITREATTCK     `json:"attacks,omitempty" parquet:"attacks"`
	CreatedTime      *time.Time       `json:"created_time,omitempty" parquet:"created_time"`
	DataSources      []string         `json:"data_sources,omitempty" parquet:"data_sources"`
	Desc             *string          `json:"desc,omitempty" parquet:"desc"`
	FirstSeenTime    *time.Time       `json:"first_seen_time,omitempty" parquet:"first_seen_time"`
	KillChain        []KillChainPhase `json:"kill_chain,omitempty" parquet:"kill_chain"`
	LastSeenTime     *time.Time       `json:"last_seen_time,omitempty" parquet:"last_seen_time"`
	ModifiedTime     *time.Time       `json:"modified_time,omitempty" parquet:"modified_time"`
	ProductUID       *string          `json:"product_uid,omitempty" parquet:"product_uid"`
	RelatedAnalytics []Analytic       `json:"related_analytics,omitempty" parquet:"related_analytics"`
	RelatedEvents    []RelatedEvent   `json:"related_events,omitempty" parquet:"related_events"`
	SrcURL           *string          `json:"src_url,omitempty" parquet:"src_url"`
	Title            string           `json:"title" parquet:"title"`
	Types            []string         `json:"types,omitempty" parquet:"types"`
	UID              string           `json:"uid" parquet:"uid"`
}
