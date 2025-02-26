package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// MITREATTCKFields defines the Arrow fields for MITREATTCK.
var MITREATTCKFields = []arrow.Field{
	{Name: "sub_technique", Type: arrow.StructOf(SubTechniqueFields...)},
	{Name: "tactic", Type: arrow.StructOf(TacticFields...)},
	{Name: "tactics", Type: arrow.ListOf(arrow.StructOf(TacticFields...))},
	{Name: "technique", Type: arrow.StructOf(TechniqueFields...)},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// MITREATTCKSchema is the Arrow schema for MITREATTCK.
var MITREATTCKSchema = arrow.NewSchema(MITREATTCKFields, nil)

// MITREATTCK represents MITRE ATT&CK® details.
type MITREATTCK struct {
	SubTechnique *SubTechnique `json:"sub_technique,omitempty"`
	Tactic       *Tactic       `json:"tactic,omitempty"`
	Tactics      []Tactic      `json:"tactics,omitempty"`
	Technique    *Technique    `json:"technique,omitempty"`
	Version      string        `json:"version"`
}

// WriteToParquet writes the MITREATTCK fields to the provided Arrow StructBuilder.
func (m *MITREATTCK) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: sub_technique.
	subTechB := sb.FieldBuilder(0).(*array.StructBuilder)
	if m.SubTechnique != nil {
		subTechB.Append(true)
		m.SubTechnique.WriteToParquet(subTechB)
	} else {
		subTechB.AppendNull()
	}

	// Field 1: tactic.
	tacticB := sb.FieldBuilder(1).(*array.StructBuilder)
	if m.Tactic != nil {
		tacticB.Append(true)
		m.Tactic.WriteToParquet(tacticB)
	} else {
		tacticB.AppendNull()
	}

	// Field 2: tactics (list of Tactic structs).
	tacticsB := sb.FieldBuilder(2).(*array.ListBuilder)
	if len(m.Tactics) > 0 {
		tacticsB.Append(true)
		tacticsValB := tacticsB.ValueBuilder().(*array.StructBuilder)
		for _, t := range m.Tactics {
			tacticsValB.Append(true)
			t.WriteToParquet(tacticsValB)
		}
	} else {
		tacticsB.AppendNull()
	}

	// Field 3: technique.
	techB := sb.FieldBuilder(3).(*array.StructBuilder)
	if m.Technique != nil {
		techB.Append(true)
		m.Technique.WriteToParquet(techB)
	} else {
		techB.AppendNull()
	}

	// Field 4: version.
	versionB := sb.FieldBuilder(4).(*array.StringBuilder)
	versionB.Append(m.Version)
}
