package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var ScriptFields = []arrow.Field{
	{Name: "file", Type: FileStruct, Nullable: true},
	{Name: "hashes", Type: arrow.ListOf(FingerprintStruct), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "script_content", Type: arrow.BinaryTypes.LargeString, Nullable: false},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ScriptStruct = arrow.StructOf(ScriptFields...)
var ScriptClassname = "script"

type Script struct {
	File          *File          `json:"file,omitempty" parquet:"file,optional"`
	Hashes        []*Fingerprint `json:"hashes,omitempty" parquet:"hashes,list,optional"`
	Name          *string        `json:"name,omitempty" parquet:"name,optional"`
	ScriptContent *LongString    `json:"script_content,omitempty" parquet:"script_content,optional"`
	Type          *string        `json:"type,omitempty" parquet:"type,optional"`
	TypeID        int32          `json:"type_id" parquet:"type_id"`
	UID           *string        `json:"uid,omitempty" parquet:"uid,optional"`
}
