package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var EncryptionDetailsFields = []arrow.Field{
	{Name: "algorithm", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "algorithm_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "key_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "key_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var EncryptionDetailsStruct = arrow.StructOf(EncryptionDetailsFields...)
var EncryptionDetailsClassname = "encryption_details"

type EncryptionDetails struct {
	Algorithm   *string `json:"algorithm,omitempty" parquet:"algorithm,optional"`
	AlgorithmID *int32  `json:"algorithm_id,omitempty" parquet:"algorithm_id,optional"`
	KeyLength   *int32  `json:"key_length,omitempty" parquet:"key_length,optional"`
	KeyUID      *string `json:"key_uid,omitempty" parquet:"key_uid,optional"`
	Type        *string `json:"type,omitempty" parquet:"type,optional"`
}
