package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// ResourceDetailsFields defines the Arrow fields for ResourceDetails.
var ResourceDetailsFields = []arrow.Field{
	{Name: "criticality", Type: arrow.BinaryTypes.String},
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "group", Type: arrow.StructOf(GroupFields...)},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "namespace", Type: arrow.BinaryTypes.String},
	{Name: "owner", Type: arrow.StructOf(UserFields...)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

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
