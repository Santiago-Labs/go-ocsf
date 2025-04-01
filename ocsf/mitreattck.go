package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// MITREATTCKFields defines the Arrow fields for MITREATTCK.
var MITREATTCKFields = []arrow.Field{
	{Name: "sub_technique", Type: SubTechniqueStruct},
	{Name: "tactic", Type: TacticStruct},
	{Name: "tactics", Type: arrow.ListOf(TacticStruct)},
	{Name: "technique", Type: TechniqueStruct},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var MITREATTCKStruct = arrow.StructOf(MITREATTCKFields...)

// MITREATTCK represents MITRE ATT&CK® details.
type MITREATTCK struct {
	SubTechnique *SubTechnique `json:"sub_technique,omitempty" parquet:"sub_technique"`
	Tactic       *Tactic       `json:"tactic,omitempty" parquet:"tactic"`
	Tactics      []Tactic      `json:"tactics,omitempty" parquet:"tactics"`
	Technique    *Technique    `json:"technique,omitempty" parquet:"technique"`
	Version      string        `json:"version" parquet:"version"`
}
