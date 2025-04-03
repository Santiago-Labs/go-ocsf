package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AgentFields defines the Arrow fields for Agent.
var AgentFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
	{Name: "policies", Type: arrow.ListOf(PolicyStruct)},
}

var AgentStruct = arrow.StructOf(AgentFields...)
var AgentClassname = "agent"

type Agent struct {
	Name       *string   `json:"name,omitempty" parquet:"name"`
	Type       *string   `json:"type,omitempty" parquet:"type"`
	TypeID     *int      `json:"type_id,omitempty" parquet:"type_id"`
	UID        *string   `json:"uid,omitempty" parquet:"uid"`
	UIDAlt     *string   `json:"uid_alt,omitempty" parquet:"uid_alt"`
	VendorName *string   `json:"vendor_name,omitempty" parquet:"vendor_name"`
	Version    *string   `json:"version,omitempty" parquet:"version"`
	Policies   []*Policy `json:"policies,omitempty" parquet:"policies"`
}
