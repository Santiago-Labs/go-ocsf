package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var EPSSFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "percentile", Type: arrow.PrimitiveTypes.Float64},
	{Name: "score", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var EPSSSchema = arrow.NewSchema(EPSSFields, nil)

// EPSS represents the EPSS object.
type EPSS struct {
	CreatedTime   *int     `json:"created_time,omitempty"`
	CreatedTimeDt *string  `json:"created_time_dt,omitempty"`
	Percentile    *float64 `json:"percentile,omitempty"`
	Score         string   `json:"score"`
	Version       *string  `json:"version,omitempty"`
}

// WriteToParquet writes the EPSS fields to the provided Arrow StructBuilder.
func (e *EPSS) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: CreatedTime.
	createdTimeB := sb.FieldBuilder(0).(*array.Int32Builder)
	if e.CreatedTime != nil {
		createdTimeB.Append(int32(*e.CreatedTime))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 1: CreatedTimeDt.
	createdTimeDtB := sb.FieldBuilder(1).(*array.StringBuilder)
	if e.CreatedTimeDt != nil {
		createdTimeDtB.Append(*e.CreatedTimeDt)
	} else {
		createdTimeDtB.AppendNull()
	}

	// Field 2: Percentile.
	percentileB := sb.FieldBuilder(2).(*array.Float64Builder)
	if e.Percentile != nil {
		percentileB.Append(*e.Percentile)
	} else {
		percentileB.AppendNull()
	}

	// Field 3: Score.
	scoreB := sb.FieldBuilder(3).(*array.StringBuilder)
	scoreB.Append(e.Score)

	// Field 4: Version.
	versionB := sb.FieldBuilder(4).(*array.StringBuilder)
	if e.Version != nil {
		versionB.Append(*e.Version)
	} else {
		versionB.AppendNull()
	}
}
