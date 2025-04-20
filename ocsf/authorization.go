package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AuthorizationFields defines the Arrow fields for Authorization.
var AuthorizationFields = []arrow.Field{
	{Name: "decision", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "policy", Type: PolicyStruct, Nullable: true},
}

var AuthorizationStruct = arrow.StructOf(AuthorizationFields...)
var AuthorizationClassname = "authorization"

type Authorization struct {
	Decision *string `json:"decision,omitempty" parquet:"decision,optional"`
	Policy   *Policy `json:"policy,omitempty" parquet:"policy,optional"`
}
