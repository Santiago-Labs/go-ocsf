package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// ReputationFields defines the Arrow fields for Reputation.
var ReputationFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "score", Type: arrow.BinaryTypes.String},
	{Name: "score_id", Type: arrow.PrimitiveTypes.Int32},
}

var ReputationStruct = arrow.StructOf(ReputationFields...)

type Reputation struct {
	BaseScore float64 `json:"base_score" parquet:"base_score"`
	Provider  *string `json:"provider,omitempty" parquet:"provider"`
	Score     *string `json:"score,omitempty" parquet:"score"`
	ScoreID   int     `json:"score_id" parquet:"score_id"`
}
