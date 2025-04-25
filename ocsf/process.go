package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var ParentProcessFields = []arrow.Field{
	{Name: "ancestry", Type: ProcessEntityStruct, Nullable: true},
	{Name: "cmd_line", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "environment_variables", Type: arrow.ListOf(EnvironmentVarStruct), Nullable: true},
	{Name: "file", Type: FileStruct, Nullable: true},
	{Name: "integrity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "integrity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "lineage", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "loaded_modules", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "sandbox", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "session", Type: SessionStruct, Nullable: true},
	{Name: "terminated_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "tid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "user", Type: UserStruct, Nullable: true},
	{Name: "working_directory", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "xattributes", Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String), Nullable: true},
}

// ProcessFields defines the Arrow fields for Process.
var ProcessFields = []arrow.Field{
	{Name: "ancestry", Type: ProcessEntityStruct, Nullable: true},
	{Name: "cmd_line", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "environment_variables", Type: arrow.ListOf(EnvironmentVarStruct), Nullable: true},
	{Name: "file", Type: FileStruct, Nullable: true},
	{Name: "integrity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "integrity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "lineage", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "loaded_modules", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "parent_process", Type: ParentProcessStruct, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "sandbox", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "session", Type: SessionStruct, Nullable: true},
	{Name: "terminated_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "tid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "user", Type: UserStruct, Nullable: true},
	{Name: "working_directory", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "xattributes", Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String), Nullable: true},
}

var ParentProcessStruct = arrow.StructOf(ParentProcessFields...)
var ParentProcessClassname = "parent_process"

var ProcessStruct = arrow.StructOf(ProcessFields...)
var ProcessClassname = "process"

type Process struct {
	Ancestry         *ProcessEntity    `json:"ancestry" parquet:"ancestry,optional" ch:"ancestry"`
	CmdLine          *string           `json:"cmd_line" parquet:"cmd_line,optional" ch:"cmd_line"`
	CreatedTime      *int64            `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	EnvironmentVars  []*EnvironmentVar `json:"environment_variables" parquet:"environment_variables,list,optional" ch:"environment_variables"`
	File             *File             `json:"file" parquet:"file,optional" ch:"file"`
	Integrity        *string           `json:"integrity" parquet:"integrity,optional" ch:"integrity"`
	IntegrityID      *int64            `json:"integrity_id" parquet:"integrity_id,optional" ch:"integrity_id"`
	Lineage          []string          `json:"lineage" parquet:"lineage,list,optional" ch:"lineage"`
	LoadedModules    []string          `json:"loaded_modules" parquet:"loaded_modules,list,optional" ch:"loaded_modules"`
	Name             *string           `json:"name" parquet:"name,optional" ch:"name"`
	ParentProcess    *ParentProcess    `json:"parent_process" parquet:"parent_process,optional" ch:"parent_process"`
	Path             *string           `json:"path" parquet:"path,optional" ch:"path"`
	PID              *int64            `json:"pid" parquet:"pid,optional" ch:"pid"`
	Sandbox          *string           `json:"sandbox" parquet:"sandbox,optional" ch:"sandbox"`
	Session          *Session          `json:"session" parquet:"session,optional" ch:"session"`
	TerminatedTime   *int64            `json:"terminated_time" parquet:"terminated_time,optional" ch:"terminated_time"`
	TID              *int64            `json:"tid" parquet:"tid,optional" ch:"tid"`
	UID              *string           `json:"uid" parquet:"uid,optional" ch:"uid"`
	User             *User             `json:"user" parquet:"user,optional" ch:"user"`
	WorkingDirectory *string           `json:"working_directory" parquet:"working_directory,optional" ch:"working_directory"`
	XAttributes      map[string]string `json:"xattributes" parquet:"xattributes,optional" ch:"xattributes"`
}

type ParentProcess struct {
	Ancestry         *ProcessEntity    `json:"ancestry" parquet:"ancestry,optional" ch:"ancestry"`
	CmdLine          *string           `json:"cmd_line" parquet:"cmd_line,optional" ch:"cmd_line"`
	CreatedTime      *int64            `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	EnvironmentVars  []*EnvironmentVar `json:"environment_variables" parquet:"environment_variables,list,optional" ch:"environment_variables"`
	File             *File             `json:"file" parquet:"file,optional" ch:"file"`
	Integrity        *string           `json:"integrity" parquet:"integrity,optional" ch:"integrity"`
	IntegrityID      *int64            `json:"integrity_id" parquet:"integrity_id,optional" ch:"integrity_id"`
	Lineage          []string          `json:"lineage" parquet:"lineage,list,optional" ch:"lineage"`
	LoadedModules    []string          `json:"loaded_modules" parquet:"loaded_modules,list,optional" ch:"loaded_modules"`
	Name             *string           `json:"name" parquet:"name,optional" ch:"name"`
	Path             *string           `json:"path" parquet:"path,optional" ch:"path"`
	PID              *int64            `json:"pid" parquet:"pid,optional" ch:"pid"`
	Sandbox          *string           `json:"sandbox" parquet:"sandbox,optional" ch:"sandbox"`
	Session          *Session          `json:"session" parquet:"session,optional" ch:"session"`
	TerminatedTime   *int64            `json:"terminated_time" parquet:"terminated_time,optional" ch:"terminated_time"`
	TID              *int64            `json:"tid" parquet:"tid,optional" ch:"tid"`
	UID              *string           `json:"uid" parquet:"uid,optional" ch:"uid"`
	User             *User             `json:"user" parquet:"user,optional" ch:"user"`
	WorkingDirectory *string           `json:"working_directory" parquet:"working_directory,optional" ch:"working_directory"`
	XAttributes      map[string]string `json:"xattributes" parquet:"xattributes,optional" ch:"xattributes"`
}
