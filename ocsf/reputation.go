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
	BaseScore float64 `json:"base_score" parquet:"base_score"`
	Provider  *string `json:"provider,omitempty" parquet:"provider,optional"`
	Score     *string `json:"score,omitempty" parquet:"score,optional"`
	ScoreID   int32   `json:"score_id" parquet:"score_id"`
}
