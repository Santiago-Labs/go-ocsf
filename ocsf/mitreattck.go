package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// MITREATTCKFields defines the Arrow fields for MITREATTCK.
var MITREATTCKFields = []arrow.Field{
	{Name: "sub_technique", Type: SubTechniqueStruct, Nullable: true},
	{Name: "tactic", Type: TacticStruct, Nullable: true},
	{Name: "tactics", Type: arrow.ListOf(TacticStruct), Nullable: true},
	{Name: "technique", Type: TechniqueStruct, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var MITREATTCKStruct = arrow.StructOf(MITREATTCKFields...)
var MITREATTCKClassname = "attack"

// MITREATTCK represents MITRE ATT&CK® details.
type MITREATTCK struct {
	SubTechnique *SubTechnique `json:"sub_technique" parquet:"sub_technique,optional" ch:"sub_technique"`
	Tactic       *Tactic       `json:"tactic" parquet:"tactic,optional" ch:"tactic"`
	Tactics      []*Tactic     `json:"tactics" parquet:"tactics,list,optional" ch:"tactics"`
	Technique    *Technique    `json:"technique" parquet:"technique,optional" ch:"technique"`
	Version      string        `json:"version" parquet:"version" ch:"version"`
}
