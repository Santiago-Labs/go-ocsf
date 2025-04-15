package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

var ParentProcessFields = []arrow.Field{
	{Name: "ancestry", Type: ProcessEntityStruct},
	{Name: "cmd_line", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "environment_variables", Type: arrow.ListOf(EnvironmentVarStruct)},
	{Name: "file", Type: FileStruct},
	{Name: "integrity", Type: arrow.BinaryTypes.String},
	{Name: "integrity_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "lineage", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "loaded_modules", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "sandbox", Type: arrow.BinaryTypes.String},
	{Name: "session", Type: SessionStruct},
	{Name: "terminated_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "tid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "user", Type: UserStruct},
	{Name: "working_directory", Type: arrow.BinaryTypes.String},
}

// ProcessFields defines the Arrow fields for Process.
var ProcessFields = []arrow.Field{
	{Name: "ancestry", Type: ProcessEntityStruct},
	{Name: "cmd_line", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "environment_variables", Type: arrow.ListOf(EnvironmentVarStruct)},
	{Name: "file", Type: FileStruct},
	{Name: "integrity", Type: arrow.BinaryTypes.String},
	{Name: "integrity_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "lineage", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "loaded_modules", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "parent_process", Type: ParentProcessStruct},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "pid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "sandbox", Type: arrow.BinaryTypes.String},
	{Name: "session", Type: SessionStruct},
	{Name: "terminated_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "tid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "user", Type: UserStruct},
	{Name: "working_directory", Type: arrow.BinaryTypes.String},
	{Name: "xattributes", Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String)},
}

var ParentProcessStruct = arrow.StructOf(ParentProcessFields...)
var ProcessStruct = arrow.StructOf(ProcessFields...)

type Process struct {
	Ancestry         *ProcessEntity     `json:"ancestry,omitempty" parquet:"ancestry"`
	CmdLine          *string            `json:"cmd_line,omitempty" parquet:"cmd_line"`
	CreatedTime      *int64             `json:"created_time,omitempty" parquet:"created_time"`
	EnvironmentVars  []*EnvironmentVar  `json:"environment_variables,omitempty" parquet:"environment_variables"`
	File             *File              `json:"file,omitempty" parquet:"file"`
	Integrity        *string            `json:"integrity,omitempty" parquet:"integrity"`
	IntegrityID      *int               `json:"integrity_id,omitempty" parquet:"integrity_id"`
	Lineage          []*string          `json:"lineage,omitempty" parquet:"lineage"`
	LoadedModules    []*string          `json:"loaded_modules,omitempty" parquet:"loaded_modules"`
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

// ParentProcess is the parent process of a process, we use a different struct
// to avoid self reference because go-parquet does not support one type
// referencing itself.
type ParentProcess struct {
	Ancestry        *ProcessEntity    `json:"ancestry,omitempty" parquet:"ancestry"`
	CmdLine         *string           `json:"cmd_line,omitempty" parquet:"cmd_line"`
	CreatedTime     *int64            `json:"created_time,omitempty" parquet:"created_time"`
	EnvironmentVars []*EnvironmentVar `json:"environment_variables,omitempty" parquet:"environment_variables"`
	File            *File             `json:"file,omitempty" parquet:"file"`
	Integrity       *string           `json:"integrity,omitempty" parquet:"integrity"`
	IntegrityID     *int              `json:"integrity_id,omitempty" parquet:"integrity_id"`
	Lineage         []*string         `json:"lineage,omitempty" parquet:"lineage"`
	LoadedModules   []*string         `json:"loaded_modules,omitempty" parquet:"loaded_modules"`
	Name            *string           `json:"name,omitempty" parquet:"name"`
}
