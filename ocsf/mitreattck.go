package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
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
	SubTechnique *SubTechnique `json:"sub_technique,omitempty" parquet:"sub_technique"`
	Tactic       *Tactic       `json:"tactic,omitempty" parquet:"tactic"`
	Tactics      []Tactic      `json:"tactics,omitempty" parquet:"tactics"`
	Technique    *Technique    `json:"technique,omitempty" parquet:"technique"`
	Version      string        `json:"version" parquet:"version"`
}
