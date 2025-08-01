// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_4_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type EncryptionDetails struct {

	// Encryption Algorithm: The encryption algorithm used, normalized to the caption of 'algorithm_id
	Algorithm *string `json:"algorithm,omitempty" parquet:"algorithm,optional"`

	// Encryption Algorithm ID: The encryption algorithm used.
	AlgorithmId *int32 `json:"algorithm_id,omitempty" parquet:"algorithm_id,optional"`

	// Encryption Key Length: The length of the encryption key used.
	KeyLength *int32 `json:"key_length,omitempty" parquet:"key_length,optional"`

	// Key UID: The unique identifier of the key used for encrpytion. For example, AWS KMS Key ARN.
	KeyUid *string `json:"key_uid,omitempty" parquet:"key_uid,optional"`

	// Encryption Type: The type of the encryption used.
	Type *string `json:"type,omitempty" parquet:"type,optional"`
}

var EncryptionDetailsFields = []arrow.Field{
	{Name: "algorithm", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "algorithm_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "key_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "key_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var EncryptionDetailsStruct = arrow.StructOf(EncryptionDetailsFields...)

var EncryptionDetailsSchema = arrow.NewSchema(EncryptionDetailsFields, nil)
var EncryptionDetailsClassname = "encryption_details"
