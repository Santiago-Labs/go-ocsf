package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AccountFields defines the Arrow fields for Account.
var AccountFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var AccountStruct = arrow.StructOf(AccountFields...)
var AccountClassname = "account"

type Account struct {
	Name *string `json:"name" parquet:"name,optional" ch:"name"`
	Type *string `json:"type" parquet:"type,optional" ch:"type"`
	// TypeID enum: [3,6,99,0,1,2,10,4,5,7,8,9]
	TypeID *int64  `json:"type_id" parquet:"type_id,optional" ch:"type_id"`
	UID    *string `json:"uid" parquet:"uid,optional" ch:"uid"`
}
