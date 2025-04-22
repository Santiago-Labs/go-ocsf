package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ProcessEntityFields defines the Arrow fields for Process Entity.
var ProcessEntityFields = []arrow.Field{
	{Name: "cmd_line", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ProcessEntityStruct = arrow.StructOf(ProcessEntityFields...)
var ProcessEntityClassname = "process_entity"

type ProcessEntity struct {
	CmdLine     *string `json:"cmd_line,omitempty" parquet:"cmd_line,optional"`
	CreatedTime *int64  `json:"created_time,omitempty" parquet:"created_time,optional"`
	Name        *string `json:"name,omitempty" parquet:"name,optional"`
	Path        *string `json:"path,omitempty" parquet:"path,optional"`
	PID         *int64  `json:"pid,omitempty" parquet:"pid,optional"`
	UID         *string `json:"uid,omitempty" parquet:"uid,optional"`
}
