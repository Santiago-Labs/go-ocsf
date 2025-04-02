package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// FingerprintFields defines the Arrow fields for Fingerprint.
var FingerprintFields = []arrow.Field{
	{Name: "algorithm", Type: arrow.BinaryTypes.String},
	{Name: "algorithm_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

var FingerprintStruct = arrow.StructOf(FingerprintFields...)

type Fingerprint struct {
	Algorithm   *string `json:"algorithm,omitempty" parquet:"algorithm"`
	AlgorithmID int32   `json:"algorithm_id" parquet:"algorithm_id"`
	Value       string  `json:"value" parquet:"value"`
}
