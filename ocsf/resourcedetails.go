package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ResourceDetailsFields defines the Arrow fields for ResourceDetails.
var ResourceDetailsFields = []arrow.Field{
	{Name: "criticality", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "group", Type: GroupStruct, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "namespace", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "owner", Type: UserStruct, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ResourceDetailsStruct = arrow.StructOf(ResourceDetailsFields...)
var ResourceDetailsClassname = "resource_details"

type ResourceDetails struct {
	Criticality *string  `json:"criticality,omitempty" parquet:"criticality,optional" ch:"criticality,omitempty"`
	Data        *string  `json:"data,omitempty" parquet:"data,optional" ch:"data,omitempty"` // JSON blob
	Group       *Group   `json:"group,omitempty" parquet:"group,optional" ch:"group,omitempty"`
	Labels      []string `json:"labels,omitempty" parquet:"labels,list,optional" ch:"labels,omitempty"`
	UID         *string  `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
	Name        *string  `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	Namespace   *string  `json:"namespace,omitempty" parquet:"namespace,optional" ch:"namespace,omitempty"`
	Owner       *User    `json:"owner,omitempty" parquet:"owner,optional" ch:"owner,omitempty"`
	Type        *string  `json:"type,omitempty" parquet:"type,optional" ch:"type,omitempty"`
	Version     *string  `json:"version,omitempty" parquet:"version,optional" ch:"version,omitempty"`
}
