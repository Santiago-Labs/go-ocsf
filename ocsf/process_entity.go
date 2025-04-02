package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// ProcessEntityFields defines the Arrow fields for Process Entity.
var ProcessEntityFields = []arrow.Field{
	{Name: "cmd_line", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int64},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var ProcessEntityStruct = arrow.StructOf(ProcessEntityFields...)

type ProcessEntity struct {
	CmdLine     *string    `json:"cmd_line,omitempty" parquet:"cmd_line"`
	CreatedTime *time.Time `json:"created_time,omitempty" parquet:"created_time"`
	Name        *string    `json:"name,omitempty" parquet:"name"`
	Path        *string    `json:"path,omitempty" parquet:"path"`
	PID         *int64     `json:"pid,omitempty" parquet:"pid"`
	UID         *string    `json:"uid,omitempty" parquet:"uid"`
}
