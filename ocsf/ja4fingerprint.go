package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var JA4FingerprintFields = []arrow.Field{
	{Name: "section_a", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "section_b", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "section_c", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "section_d", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var JA4FingerprintStruct = arrow.StructOf(JA4FingerprintFields...)
var JA4FingerprintClassname = "ja4_fingerprint"

type JA4Fingerprint struct {
	SectionA *string `json:"section_a,omitempty" parquet:"section_a,optional"`
	SectionB *string `json:"section_b,omitempty" parquet:"section_b,optional"`
	SectionC *string `json:"section_c,omitempty" parquet:"section_c,optional"`
	SectionD *string `json:"section_d,omitempty" parquet:"section_d,optional"`
	Type     *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID   int32   `json:"type_id" parquet:"type_id"`
	Value    string  `json:"value" parquet:"value"`
}
