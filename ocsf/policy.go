package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// PolicyFields defines the Arrow fields for Policy.
var PolicyFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "group", Type: GroupStruct, Nullable: true},
	{Name: "is_applied", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var PolicyStruct = arrow.StructOf(PolicyFields...)
var PolicyClassname = "policy"

type Policy struct {
	Desc      *string `json:"desc,omitempty" parquet:"desc,optional"`
	Group     *Group  `json:"group,omitempty" parquet:"group,optional"`
	IsApplied *bool   `json:"is_applied,omitempty" parquet:"is_applied,optional"`
	Name      *string `json:"name,omitempty" parquet:"name,optional"`
	UID       *string `json:"uid,omitempty" parquet:"uid,optional"`
	Version   *string `json:"version,omitempty" parquet:"version,optional"`
}
