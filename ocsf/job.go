package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var JobFields = []arrow.Field{
	{Name: "cmd_line", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "file", Type: FileStruct, Nullable: false},
	{Name: "last_run_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "next_run_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "run_state", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "run_state_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var JobStruct = arrow.StructOf(JobFields...)
var JobClassname = "job"

type Job struct {
	CmdLine     *string `json:"cmd_line,omitempty" parquet:"cmd_line,optional"`
	CreatedTime *int64  `json:"created_time,omitempty" parquet:"created_time,optional"`
	Desc        *string `json:"desc,omitempty" parquet:"desc,optional"`
	Name        *string `json:"name,omitempty" parquet:"name,optional"`
	NextRunTime *int64  `json:"next_run_time,omitempty" parquet:"next_run_time,optional"`
	RunState    *string `json:"run_state,omitempty" parquet:"run_state,optional"`
	RunStateID  *int32  `json:"run_state_id,omitempty" parquet:"run_state_id,optional"`
	UID         *string `json:"uid,omitempty" parquet:"uid,optional"`
}
