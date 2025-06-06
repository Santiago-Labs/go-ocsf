package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// FingerprintFields defines the Arrow fields for Fingerprint.
var FingerprintFields = []arrow.Field{
	{Name: "algorithm", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "algorithm_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var FingerprintStruct = arrow.StructOf(FingerprintFields...)
var FingerprintClassname = "fingerprint"

type Fingerprint struct {
	Algorithm   *string `json:"algorithm,omitempty" parquet:"algorithm,optional"`
	AlgorithmID int32   `json:"algorithm_id" parquet:"algorithm_id"`
	Value       string  `json:"value" parquet:"value"`
}
