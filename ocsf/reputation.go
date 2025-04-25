package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ReputationFields defines the Arrow fields for Reputation.
var ReputationFields = []arrow.Field{
	{Name: "base_score", Type: arrow.PrimitiveTypes.Float64, Nullable: false},
	{Name: "provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "score", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "score_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
}

var ReputationStruct = arrow.StructOf(ReputationFields...)
var ReputationClassname = "reputation"

type Reputation struct {
	BaseScore float64 `json:"base_score" parquet:"base_score" ch:"base_score"`
	Provider  *string `json:"provider" parquet:"provider,optional" ch:"provider"`
	Score     *string `json:"score" parquet:"score,optional" ch:"score"`
	ScoreID   int     `json:"score_id" parquet:"score_id" ch:"score_id"`
}
