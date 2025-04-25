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
	CmdLine     *string `json:"cmd_line" parquet:"cmd_line,optional" ch:"cmd_line"`
	CreatedTime *int64  `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	Name        *string `json:"name" parquet:"name,optional" ch:"name"`
	Path        *string `json:"path" parquet:"path,optional" ch:"path"`
	PID         *int64  `json:"pid" parquet:"pid,optional" ch:"pid"`
	UID         *string `json:"uid" parquet:"uid,optional" ch:"uid"`
}
