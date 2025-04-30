package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var EmailFields = []arrow.Field{
	{Name: "from", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "message_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "size", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "subject", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "to", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var EmailStruct = arrow.StructOf(EmailFields...)
var EmailClassname = "email"

type Email struct {
	From       *string   `json:"from,omitempty" parquet:"from,optional"`
	MessageUID *string   `json:"message_uid,omitempty" parquet:"message_uid,optional"`
	Size       *int64    `json:"size,omitempty" parquet:"size,optional"`
	Subject    *string   `json:"subject,omitempty" parquet:"subject,optional"`
	To         []*string `json:"to,omitempty" parquet:"to,list,optional"`
	UID        *string   `json:"uid,omitempty" parquet:"uid,optional"`
}
