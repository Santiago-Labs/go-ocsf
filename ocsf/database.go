package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var DatabaseFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "groups", Type: arrow.ListOf(GroupStruct), Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "size", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DatabaseStruct = arrow.StructOf(DatabaseFields...)
var DatabaseClassname = "database"

type Database struct {
	CreatedTime  *int64   `json:"created_time,omitempty" parquet:"created_time,optional"`
	Desc         *string  `json:"desc,omitempty" parquet:"desc,optional"`
	Groups       []*Group `json:"groups,omitempty" parquet:"groups,optional"`
	ModifiedTime *int64   `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Name         *string  `json:"name,omitempty" parquet:"name,optional"`
	Size         *int64   `json:"size,omitempty" parquet:"size,optional"`
	Type         *string  `json:"type,omitempty" parquet:"type,optional"`
	TypeID       int32    `json:"type_id" parquet:"type_id"`
	UID          *string  `json:"uid,omitempty" parquet:"uid,optional"`
}
