package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var GroupClassname = "group"

type Group struct {
	Desc       *string  `json:"desc" parquet:"desc,optional" ch:"desc" ch:"desc"`
	Domain     *string  `json:"domain" parquet:"domain,optional" ch:"domain"`
	Name       *string  `json:"name" parquet:"name,optional" ch:"name"`
	Privileges []string `json:"privileges" parquet:"privileges,list,optional" ch:"privileges"`
	Type       *string  `json:"type" parquet:"type,optional" ch:"type"`
	UID        *string  `json:"uid" parquet:"uid,optional" ch:"uid"`
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
