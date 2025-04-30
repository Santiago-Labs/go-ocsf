package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var DatabucketFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "encryption_details", Type: EncryptionDetailsStruct, Nullable: true},
	{Name: "file", Type: FileStruct, Nullable: true},
	{Name: "groups", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "is_encrypted", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_public", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "size", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DatabucketStruct = arrow.StructOf(DatabucketFields...)
var DatabucketClassname = "databucket"

type Databucket struct {
	CreatedTime       *int64             `json:"created_time,omitempty" parquet:"created_time,optional"`
	Desc              *string            `json:"desc,omitempty" parquet:"desc,optional"`
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty" parquet:"encryption_details,optional"`
	File              *File              `json:"file,omitempty" parquet:"file,optional"`
	Groups            []*string          `json:"groups,omitempty" parquet:"groups,list,optional"`
	IsEncrypted       *bool              `json:"is_encrypted,omitempty" parquet:"is_encrypted,optional"`
	IsPublic          *bool              `json:"is_public,omitempty" parquet:"is_public,optional"`
	ModifiedTime      *int64             `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Name              *string            `json:"name,omitempty" parquet:"name,optional"`
	Size              *int64             `json:"size,omitempty" parquet:"size,optional"`
	Type              *string            `json:"type,omitempty" parquet:"type,optional"`
	TypeID            int32              `json:"type_id" parquet:"type_id"`
	UID               *string            `json:"uid,omitempty" parquet:"uid,optional"`
}
