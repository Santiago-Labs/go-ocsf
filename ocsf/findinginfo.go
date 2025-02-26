package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var FindingInfoFields = []arrow.Field{
	{Name: "analytic", Type: arrow.StructOf(AnalyticFields...)},
	{Name: "attacks", Type: arrow.ListOf(arrow.StructOf(MITREATTCKFields...))},
	{Name: "created_time", Type: arrow.BinaryTypes.String},
	{Name: "data_sources", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "first_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "kill_chain", Type: arrow.ListOf(arrow.StructOf(KillChainPhaseFields...))},
	{Name: "last_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "product_uid", Type: arrow.BinaryTypes.String},
	{Name: "related_analytics", Type: arrow.ListOf(arrow.StructOf(AnalyticFields...))},
	{Name: "related_events", Type: arrow.ListOf(arrow.StructOf(RelatedEventFields...))},
	{Name: "src_url", Type: arrow.BinaryTypes.String},
	{Name: "title", Type: arrow.BinaryTypes.String},
	{Name: "types", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var FindingInfoSchema = arrow.NewSchema(FindingInfoFields, nil)

type FindingInfo struct {
	Analytic         *Analytic        `json:"analytic,omitempty"`
	Attacks          []MITREATTCK     `json:"attacks,omitempty"`
	CreatedTime      *time.Time       `json:"created_time,omitempty"`
	DataSources      []string         `json:"data_sources,omitempty"`
	Desc             *string          `json:"desc,omitempty"`
	FirstSeenTime    *time.Time       `json:"first_seen_time,omitempty"`
	KillChain        []KillChainPhase `json:"kill_chain,omitempty"`
	LastSeenTime     *time.Time       `json:"last_seen_time,omitempty"`
	ModifiedTime     *time.Time       `json:"modified_time,omitempty"`
	ProductUID       *string          `json:"product_uid,omitempty"`
	RelatedAnalytics []Analytic       `json:"related_analytics,omitempty"`
	RelatedEvents    []RelatedEvent   `json:"related_events,omitempty"`
	SrcURL           *string          `json:"src_url,omitempty"`
	Title            string           `json:"title"`
	Types            []string         `json:"types,omitempty"`
	UID              string           `json:"uid"`
}

// WriteToParquet writes the FindingInfo fields to the provided Arrow StructBuilder.
func (fi *FindingInfo) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: analytic (nested struct)
	analyticB := sb.FieldBuilder(0).(*array.StructBuilder)
	if fi.Analytic != nil {
		analyticB.Append(true)
		fi.Analytic.WriteToParquet(analyticB)
	} else {
		analyticB.AppendNull()
	}

	// Field 1: attacks (list of MITREATTCK structs)
	attacksB := sb.FieldBuilder(1).(*array.ListBuilder)
	if len(fi.Attacks) > 0 {
		attacksB.Append(true)
		attacksValB := attacksB.ValueBuilder().(*array.StructBuilder)
		for _, attack := range fi.Attacks {
			attacksValB.Append(true)
			attack.WriteToParquet(attacksValB)
		}
	} else {
		attacksB.AppendNull()
	}

	// Field 2: created_time (formatted as RFC3339 string)
	createdTimeB := sb.FieldBuilder(2).(*array.StringBuilder)
	if fi.CreatedTime != nil {
		createdTimeB.Append(fi.CreatedTime.Format(time.RFC3339))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 3: data_sources (list of strings)
	dataSourcesB := sb.FieldBuilder(3).(*array.ListBuilder)
	if len(fi.DataSources) > 0 {
		dataSourcesB.Append(true)
		dataSourcesValB := dataSourcesB.ValueBuilder().(*array.StringBuilder)
		for _, ds := range fi.DataSources {
			dataSourcesValB.Append(ds)
		}
	} else {
		dataSourcesB.AppendNull()
	}

	// Field 4: desc (string)
	descB := sb.FieldBuilder(4).(*array.StringBuilder)
	if fi.Desc != nil {
		descB.Append(*fi.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 5: first_seen_time (formatted as RFC3339 string)
	firstSeenB := sb.FieldBuilder(5).(*array.StringBuilder)
	if fi.FirstSeenTime != nil {
		firstSeenB.Append(fi.FirstSeenTime.Format(time.RFC3339))
	} else {
		firstSeenB.AppendNull()
	}

	// Field 6: kill_chain (list of KillChainPhase structs)
	killChainB := sb.FieldBuilder(6).(*array.ListBuilder)
	if len(fi.KillChain) > 0 {
		killChainB.Append(true)
		killChainValB := killChainB.ValueBuilder().(*array.StructBuilder)
		for _, kc := range fi.KillChain {
			killChainValB.Append(true)
			kc.WriteToParquet(killChainValB)
		}
	} else {
		killChainB.AppendNull()
	}

	// Field 7: last_seen_time (formatted as RFC3339 string)
	lastSeenB := sb.FieldBuilder(7).(*array.StringBuilder)
	if fi.LastSeenTime != nil {
		lastSeenB.Append(fi.LastSeenTime.Format(time.RFC3339))
	} else {
		lastSeenB.AppendNull()
	}

	// Field 8: modified_time (formatted as RFC3339 string)
	modTimeB := sb.FieldBuilder(8).(*array.StringBuilder)
	if fi.ModifiedTime != nil {
		modTimeB.Append(fi.ModifiedTime.Format(time.RFC3339))
	} else {
		modTimeB.AppendNull()
	}

	// Field 9: product_uid (string)
	productUIDB := sb.FieldBuilder(9).(*array.StringBuilder)
	if fi.ProductUID != nil {
		productUIDB.Append(*fi.ProductUID)
	} else {
		productUIDB.AppendNull()
	}

	// Field 10: related_analytics (list of Analytic structs)
	relAnalyticsB := sb.FieldBuilder(10).(*array.ListBuilder)
	if len(fi.RelatedAnalytics) > 0 {
		relAnalyticsB.Append(true)
		relAnalyticsValB := relAnalyticsB.ValueBuilder().(*array.StructBuilder)
		for _, ra := range fi.RelatedAnalytics {
			relAnalyticsValB.Append(true)
			ra.WriteToParquet(relAnalyticsValB)
		}
	} else {
		relAnalyticsB.AppendNull()
	}

	// Field 11: related_events (list of RelatedEvent structs)
	relEventsB := sb.FieldBuilder(11).(*array.ListBuilder)
	if len(fi.RelatedEvents) > 0 {
		relEventsB.Append(true)
		relEventsValB := relEventsB.ValueBuilder().(*array.StructBuilder)
		for _, re := range fi.RelatedEvents {
			relEventsValB.Append(true)
			re.WriteToParquet(relEventsValB)
		}
	} else {
		relEventsB.AppendNull()
	}

	// Field 12: src_url (string)
	srcURLB := sb.FieldBuilder(12).(*array.StringBuilder)
	if fi.SrcURL != nil {
		srcURLB.Append(*fi.SrcURL)
	} else {
		srcURLB.AppendNull()
	}

	// Field 13: title (string)
	titleB := sb.FieldBuilder(13).(*array.StringBuilder)
	titleB.Append(fi.Title)

	// Field 14: types (list of strings)
	typesB := sb.FieldBuilder(14).(*array.ListBuilder)
	if len(fi.Types) > 0 {
		typesB.Append(true)
		typesValB := typesB.ValueBuilder().(*array.StringBuilder)
		for _, t := range fi.Types {
			typesValB.Append(t)
		}
	} else {
		typesB.AppendNull()
	}

	// Field 15: uid (string)
	uidB := sb.FieldBuilder(15).(*array.StringBuilder)
	uidB.Append(fi.UID)
}
