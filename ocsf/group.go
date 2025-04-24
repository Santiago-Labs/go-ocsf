package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var GroupClassname = "group"

type Group struct {
	Desc       *string  `json:"desc,omitempty" parquet:"desc,optional" ch:"desc,omitempty" ch:"desc,omitempty"`
	Domain     *string  `json:"domain,omitempty" parquet:"domain,optional" ch:"domain,omitempty"`
	Name       *string  `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	Privileges []string `json:"privileges,omitempty" parquet:"privileges,list,optional" ch:"privileges,omitempty"`
	Type       *string  `json:"type,omitempty" parquet:"type,optional" ch:"type,omitempty"`
	UID        *string  `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
}

// GroupFields defines the fields for the Group Arrow schema.
var GroupFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "privileges", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var GroupStruct = arrow.StructOf(GroupFields...)
