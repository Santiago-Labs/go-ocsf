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
	Name       *string   `json:"name" parquet:"name,optional" ch:"name"`
	Type       *string   `json:"type" parquet:"type,optional" ch:"type"`
	TypeID     *int64    `json:"type_id" parquet:"type_id,optional" ch:"type_id"`
	UID        *string   `json:"uid" parquet:"uid,optional" ch:"uid"`
	UIDAlt     *string   `json:"uid_alt" parquet:"uid_alt,optional" ch:"uid_alt"`
	VendorName *string   `json:"vendor_name" parquet:"vendor_name,optional" ch:"vendor_name"`
	Version    *string   `json:"version" parquet:"version,optional" ch:"version"`
	Policies   []*Policy `json:"policies" parquet:"policies,list,optional" ch:"policies"`
}
