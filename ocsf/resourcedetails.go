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
	Criticality *string  `json:"criticality" parquet:"criticality,optional" ch:"criticality"`
	Data        *string  `json:"data" parquet:"data,optional" ch:"data"` // JSON blob
	Group       *Group   `json:"group" parquet:"group,optional" ch:"group"`
	Labels      []string `json:"labels" parquet:"labels,list,optional" ch:"labels"`
	UID         *string  `json:"uid" parquet:"uid,optional" ch:"uid"`
	Name        *string  `json:"name" parquet:"name,optional" ch:"name"`
	Namespace   *string  `json:"namespace" parquet:"namespace,optional" ch:"namespace"`
	Owner       *User    `json:"owner" parquet:"owner,optional" ch:"owner"`
	Type        *string  `json:"type" parquet:"type,optional" ch:"type"`
	Version     *string  `json:"version" parquet:"version,optional" ch:"version"`
}
