package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// MITREATTCKFields defines the Arrow fields for MITREATTCK.
var MITREATTCKFields = []arrow.Field{
	{Name: "sub_technique", Type: SubTechniqueStruct, Nullable: true},
	{Name: "tactic", Type: TacticStruct, Nullable: true},
	{Name: "technique", Type: TechniqueStruct, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var MITREATTCKStruct = arrow.StructOf(MITREATTCKFields...)
var MITREATTCKClassname = "attack"

// MITREATTCK represents MITRE ATT&CK® details.
type MITREATTCK struct {
	SubTechnique *SubTechnique `json:"sub_technique,omitempty" parquet:"sub_technique,optional"`
	Tactic       *Tactic       `json:"tactic,omitempty" parquet:"tactic,optional"`
	Technique    *Technique    `json:"technique,omitempty" parquet:"technique,optional"`
	Version      string        `json:"version" parquet:"version"`
}
