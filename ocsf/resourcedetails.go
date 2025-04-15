package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
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
	Criticality *string  `json:"criticality,omitempty" parquet:"criticality"`
	Data        *string  `json:"data,omitempty" parquet:"data"` // JSON blob
	Group       *Group   `json:"group,omitempty" parquet:"group"`
	Labels      []string `json:"labels,omitempty" parquet:"labels"`
	UID         *string  `json:"uid,omitempty" parquet:"uid"`
	Name        *string  `json:"name,omitempty" parquet:"name"`
	Namespace   *string  `json:"namespace,omitempty" parquet:"namespace"`
	Owner       *User    `json:"owner,omitempty" parquet:"owner"`
	Type        *string  `json:"type,omitempty" parquet:"type"`
	Version     *string  `json:"version,omitempty" parquet:"version"`
}
