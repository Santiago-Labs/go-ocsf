package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AuthorizationFields defines the Arrow fields for Authorization.
var AuthorizationFields = []arrow.Field{
	{Name: "decision", Type: arrow.BinaryTypes.String},
	{Name: "policy", Type: PolicyStruct},
}

var AuthorizationStruct = arrow.StructOf(AuthorizationFields...)
var AuthorizationClassname = "authorization"

type Authorization struct {
	Decision *string `json:"decision,omitempty" parquet:"decision"`
	Policy   *Policy `json:"policy,omitempty" parquet:"policy"`
}
