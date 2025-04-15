package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AgentFields defines the Arrow fields for Agent.
var AgentFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "policies", Type: arrow.ListOf(PolicyStruct), Nullable: true},
}

var AgentStruct = arrow.StructOf(AgentFields...)
var AgentClassname = "agent"

type Agent struct {
	Name       *string   `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	Type       *string   `json:"type,omitempty" parquet:"type,optional" ch:"type,omitempty"`
	TypeID     *int      `json:"type_id,omitempty" parquet:"type_id,optional" ch:"type_id,omitempty"`
	UID        *string   `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
	UIDAlt     *string   `json:"uid_alt,omitempty" parquet:"uid_alt,optional" ch:"uid_alt,omitempty"`
	VendorName *string   `json:"vendor_name,omitempty" parquet:"vendor_name,optional" ch:"vendor_name,omitempty"`
	Version    *string   `json:"version,omitempty" parquet:"version,optional" ch:"version,omitempty"`
	Policies   []*Policy `json:"policies,omitempty" parquet:"policies,list,optional" ch:"policies,omitempty"`
}
