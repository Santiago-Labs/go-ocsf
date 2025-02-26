package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// KillChainPhaseFields defines the Arrow fields for KillChainPhase.
var KillChainPhaseFields = []arrow.Field{
	{Name: "phase", Type: arrow.BinaryTypes.String},
	{Name: "phase_id", Type: arrow.PrimitiveTypes.Int32},
}

// KillChainPhaseSchema is the Arrow schema for KillChainPhase.
var KillChainPhaseSchema = arrow.NewSchema(KillChainPhaseFields, nil)

// KillChainPhase represents a kill chain phase.
type KillChainPhase struct {
	Phase   *string `json:"phase,omitempty"`
	PhaseID int     `json:"phase_id"`
}

// WriteToParquet writes the KillChainPhase fields to the provided Arrow StructBuilder.
func (k *KillChainPhase) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Phase.
	phaseB := sb.FieldBuilder(0).(*array.StringBuilder)
	if k.Phase != nil {
		phaseB.Append(*k.Phase)
	} else {
		phaseB.AppendNull()
	}

	// Field 1: PhaseID.
	phaseIDB := sb.FieldBuilder(1).(*array.Int32Builder)
	phaseIDB.Append(int32(k.PhaseID))
}
