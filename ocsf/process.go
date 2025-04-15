package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
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
	Ancestry         *ProcessEntity     `json:"ancestry,omitempty" parquet:"ancestry"`
	CmdLine          *string            `json:"cmd_line,omitempty" parquet:"cmd_line"`
	CreatedTime      *int64             `json:"created_time,omitempty" parquet:"created_time"`
	EnvironmentVars  []EnvironmentVar   `json:"environment_variables,omitempty" parquet:"environment_variables"`
	File             *File              `json:"file,omitempty" parquet:"file"`
	Integrity        *string            `json:"integrity,omitempty" parquet:"integrity"`
	IntegrityID      *int               `json:"integrity_id,omitempty" parquet:"integrity_id"`
	Lineage          []string           `json:"lineage,omitempty" parquet:"lineage"`
	LoadedModules    []string           `json:"loaded_modules,omitempty" parquet:"loaded_modules"`
	Name             *string            `json:"name,omitempty" parquet:"name"`
	ParentProcess    *ParentProcess     `json:"parent_process,omitempty" parquet:"parent_process"`
	Path             *string            `json:"path,omitempty" parquet:"path"`
	PID              *int               `json:"pid,omitempty" parquet:"pid"`
	Sandbox          *string            `json:"sandbox,omitempty" parquet:"sandbox"`
	Session          *Session           `json:"session,omitempty" parquet:"session"`
	TerminatedTime   *time.Time         `json:"terminated_time,omitempty" parquet:"terminated_time"`
	TID              *int               `json:"tid,omitempty" parquet:"tid"`
	UID              *string            `json:"uid,omitempty" parquet:"uid"`
	User             *User              `json:"user,omitempty" parquet:"user"`
	WorkingDirectory *string            `json:"working_directory,omitempty" parquet:"working_directory"`
	XAttributes      *map[string]string `json:"xattributes,omitempty" parquet:"xattributes"`
}

type ParentProcess struct {
	Ancestry         *ProcessEntity     `json:"ancestry,omitempty" parquet:"ancestry"`
	CmdLine          *string            `json:"cmd_line,omitempty" parquet:"cmd_line"`
	CreatedTime      *int64             `json:"created_time,omitempty" parquet:"created_time"`
	EnvironmentVars  []EnvironmentVar   `json:"environment_variables,omitempty" parquet:"environment_variables"`
	File             *File              `json:"file,omitempty" parquet:"file"`
	Integrity        *string            `json:"integrity,omitempty" parquet:"integrity"`
	IntegrityID      *int               `json:"integrity_id,omitempty" parquet:"integrity_id"`
	Lineage          []string           `json:"lineage,omitempty" parquet:"lineage"`
	LoadedModules    []string           `json:"loaded_modules,omitempty" parquet:"loaded_modules"`
	Name             *string            `json:"name,omitempty" parquet:"name"`
	Path             *string            `json:"path,omitempty" parquet:"path"`
	PID              *int               `json:"pid,omitempty" parquet:"pid"`
	Sandbox          *string            `json:"sandbox,omitempty" parquet:"sandbox"`
	Session          *Session           `json:"session,omitempty" parquet:"session"`
	TerminatedTime   *time.Time         `json:"terminated_time,omitempty" parquet:"terminated_time"`
	TID              *int               `json:"tid,omitempty" parquet:"tid"`
	UID              *string            `json:"uid,omitempty" parquet:"uid"`
	User             *User              `json:"user,omitempty" parquet:"user"`
	WorkingDirectory *string            `json:"working_directory,omitempty" parquet:"working_directory"`
	XAttributes      *map[string]string `json:"xattributes,omitempty" parquet:"xattributes"`
}
