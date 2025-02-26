package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// ReputationFields defines the Arrow fields for Reputation.
var ReputationFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "score", Type: arrow.BinaryTypes.String},
	{Name: "score_id", Type: arrow.PrimitiveTypes.Int32},
}

// ReputationSchema is the Arrow schema for Reputation.
var ReputationSchema = arrow.NewSchema(ReputationFields, nil)

// Reputation represents reputation details.
type Reputation struct {
	BaseScore float64 `json:"base_score"`
	Provider  *string `json:"provider,omitempty"`
	Score     *string `json:"score,omitempty"`
	ScoreID   int     `json:"score_id"`
}

// WriteToParquet writes the Reputation fields to the provided Arrow StructBuilder.
func (r *Reputation) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: BaseScore.
	baseScoreB := sb.FieldBuilder(0).(*array.Float64Builder)
	baseScoreB.Append(r.BaseScore)

	// Field 1: Provider.
	providerB := sb.FieldBuilder(1).(*array.StringBuilder)
	if r.Provider != nil {
		providerB.Append(*r.Provider)
	} else {
		providerB.AppendNull()
	}

	// Field 2: Score.
	scoreB := sb.FieldBuilder(2).(*array.StringBuilder)
	if r.Score != nil {
		scoreB.Append(*r.Score)
	} else {
		scoreB.AppendNull()
	}

	// Field 3: ScoreID.
	scoreIDB := sb.FieldBuilder(3).(*array.Int32Builder)
	scoreIDB.Append(int32(r.ScoreID))
}
