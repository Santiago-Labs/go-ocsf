package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// SubjectAlternativeNameFields defines the Arrow fields for Subject Alternative Name.
var SubjectAlternativeNameFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: false},
}

var SubjectAlternativeNameStruct = arrow.StructOf(SubjectAlternativeNameFields...)
var SubjectAlternativeNameClassname = "san"

type SubjectAlternativeName struct {
	Name string `json:"name" parquet:"name"`
	Type string `json:"type" parquet:"type"`
}
