package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AccountFields defines the Arrow fields for Account.
var AccountFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var AccountStruct = arrow.StructOf(AccountFields...)

// AccountSchema is the Arrow schema for Account.
var AccountSchema = arrow.NewSchema(AccountFields, nil)

type Account struct {
	Name *string `json:"name,omitempty" parquet:"name"`
	Type *string `json:"type,omitempty" parquet:"type"`
	// TypeID enum: [3,6,99,0,1,2,10,4,5,7,8,9]
	TypeID *int    `json:"type_id,omitempty" parquet:"type_id"`
	UID    *string `json:"uid,omitempty" parquet:"uid"`
}
