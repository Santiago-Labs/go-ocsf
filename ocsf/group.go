package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var GroupSchema = arrow.NewSchema(GroupFields, nil)

type Group struct {
	Desc       *string  `json:"desc,omitempty" parquet:"desc"`
	Domain     *string  `json:"domain,omitempty" parquet:"domain"`
	Name       *string  `json:"name,omitempty" parquet:"name"`
	Privileges []string `json:"privileges,omitempty" parquet:"privileges"`
	Type       *string  `json:"type,omitempty" parquet:"type"`
	UID        *string  `json:"uid,omitempty" parquet:"uid"`
}

// GroupFields defines the fields for the Group Arrow schema.
var GroupFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "privileges", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}
