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
	Ancestry         *ProcessEntity     `json:"ancestry,omitempty" parquet:"ancestry,optional" ch:"ancestry,omitempty"`
	CmdLine          *string            `json:"cmd_line,omitempty" parquet:"cmd_line,optional" ch:"cmd_line,omitempty"`
	CreatedTime      *int64             `json:"created_time,omitempty" parquet:"created_time,optional" ch:"created_time,omitempty"`
	EnvironmentVars  []*EnvironmentVar  `json:"environment_variables,omitempty" parquet:"environment_variables,list,optional" ch:"environment_variables,omitempty"`
	File             *File              `json:"file,omitempty" parquet:"file,optional" ch:"file,omitempty"`
	Integrity        *string            `json:"integrity,omitempty" parquet:"integrity,optional" ch:"integrity,omitempty"`
	IntegrityID      *int               `json:"integrity_id,omitempty" parquet:"integrity_id,optional" ch:"integrity_id,omitempty"`
	Lineage          []string           `json:"lineage,omitempty" parquet:"lineage,list,optional" ch:"lineage,omitempty"`
	LoadedModules    []string           `json:"loaded_modules,omitempty" parquet:"loaded_modules,list,optional" ch:"loaded_modules,omitempty"`
	Name             *string            `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	ParentProcess    *ParentProcess     `json:"parent_process,omitempty" parquet:"parent_process,optional" ch:"parent_process,omitempty"`
	Path             *string            `json:"path,omitempty" parquet:"path,optional" ch:"path,omitempty"`
	PID              *int               `json:"pid,omitempty" parquet:"pid,optional" ch:"pid,omitempty"`
	Sandbox          *string            `json:"sandbox,omitempty" parquet:"sandbox,optional" ch:"sandbox,omitempty"`
	Session          *Session           `json:"session,omitempty" parquet:"session,optional" ch:"session,omitempty"`
	TerminatedTime   *int64             `json:"terminated_time,omitempty" parquet:"terminated_time,optional" ch:"terminated_time,omitempty"`
	TID              *int               `json:"tid,omitempty" parquet:"tid,optional" ch:"tid,omitempty"`
	UID              *string            `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
	User             *User              `json:"user,omitempty" parquet:"user,optional" ch:"user,omitempty"`
	WorkingDirectory *string            `json:"working_directory,omitempty" parquet:"working_directory,optional" ch:"working_directory,omitempty"`
	XAttributes      map[string]*string `json:"xattributes,omitempty" parquet:"xattributes,optional" ch:"xattributes,omitempty"`
}

type ParentProcess struct {
	Ancestry         *ProcessEntity     `json:"ancestry,omitempty" parquet:"ancestry,optional" ch:"ancestry,omitempty"`
	CmdLine          *string            `json:"cmd_line,omitempty" parquet:"cmd_line,optional" ch:"cmd_line,omitempty"`
	CreatedTime      *int64             `json:"created_time,omitempty" parquet:"created_time,optional" ch:"created_time,omitempty"`
	EnvironmentVars  []*EnvironmentVar  `json:"environment_variables,omitempty" parquet:"environment_variables,list,optional" ch:"environment_variables,omitempty"`
	File             *File              `json:"file,omitempty" parquet:"file,optional" ch:"file,omitempty"`
	Integrity        *string            `json:"integrity,omitempty" parquet:"integrity,optional" ch:"integrity,omitempty"`
	IntegrityID      *int               `json:"integrity_id,omitempty" parquet:"integrity_id,optional" ch:"integrity_id,omitempty"`
	Lineage          []string           `json:"lineage,omitempty" parquet:"lineage,list,optional" ch:"lineage,omitempty"`
	LoadedModules    []string           `json:"loaded_modules,omitempty" parquet:"loaded_modules,list,optional" ch:"loaded_modules,omitempty"`
	Name             *string            `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	Path             *string            `json:"path,omitempty" parquet:"path,optional" ch:"path,omitempty"`
	PID              *int               `json:"pid,omitempty" parquet:"pid,optional" ch:"pid,omitempty"`
	Sandbox          *string            `json:"sandbox,omitempty" parquet:"sandbox,optional" ch:"sandbox,omitempty"`
	Session          *Session           `json:"session,omitempty" parquet:"session,optional" ch:"session,omitempty"`
	TerminatedTime   *int64             `json:"terminated_time,omitempty" parquet:"terminated_time,optional" ch:"terminated_time,omitempty"`
	TID              *int               `json:"tid,omitempty" parquet:"tid,optional" ch:"tid,omitempty"`
	UID              *string            `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
	User             *User              `json:"user,omitempty" parquet:"user,optional" ch:"user,omitempty"`
	WorkingDirectory *string            `json:"working_directory,omitempty" parquet:"working_directory,optional" ch:"working_directory,omitempty"`
	XAttributes      map[string]*string `json:"xattributes,omitempty" parquet:"xattributes,optional" ch:"xattributes,omitempty"`
}
