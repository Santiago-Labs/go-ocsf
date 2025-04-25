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
	Desc      *string `json:"desc" parquet:"desc,optional" ch:"desc"`
	Group     *Group  `json:"group" parquet:"group,optional" ch:"group"`
	IsApplied *bool   `json:"is_applied" parquet:"is_applied,optional" ch:"is_applied"`
	Name      *string `json:"name" parquet:"name,optional" ch:"name"`
	UID       *string `json:"uid" parquet:"uid,optional" ch:"uid"`
	Version   *string `json:"version" parquet:"version,optional" ch:"version"`
}
