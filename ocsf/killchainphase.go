package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// KillChainPhaseFields defines the Arrow fields for KillChainPhase.
var KillChainPhaseFields = []arrow.Field{
	{Name: "phase", Type: arrow.BinaryTypes.String},
	{Name: "phase_id", Type: arrow.PrimitiveTypes.Int32},
}

var KillChainPhaseStruct = arrow.StructOf(KillChainPhaseFields...)

// KillChainPhaseSchema is the Arrow schema for KillChainPhase.
var KillChainPhaseSchema = arrow.NewSchema(KillChainPhaseFields, nil)

// KillChainPhase represents a kill chain phase.
type KillChainPhase struct {
	Phase   *string `json:"phase,omitempty" parquet:"phase"`
	PhaseID int     `json:"phase_id" parquet:"phase_id"`
}
