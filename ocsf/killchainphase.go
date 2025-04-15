package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// KillChainPhaseFields defines the Arrow fields for KillChainPhase.
var KillChainPhaseFields = []arrow.Field{
	{Name: "phase", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "phase_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
}

var KillChainPhaseStruct = arrow.StructOf(KillChainPhaseFields...)
var KillChainPhaseClassname = "kill_chain_phase"

// KillChainPhase represents a kill chain phase.
type KillChainPhase struct {
	Phase   *string `json:"phase,omitempty" parquet:"phase"`
	PhaseID int     `json:"phase_id" parquet:"phase_id"`
}
