package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// AccountFields defines the Arrow fields for Account.
var AccountFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// AccountSchema is the Arrow schema for Account.
var AccountSchema = arrow.NewSchema(AccountFields, nil)

type Account struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	// TypeID enum: [3,6,99,0,1,2,10,4,5,7,8,9]
	TypeID *int    `json:"type_id,omitempty"`
	UID    *string `json:"uid,omitempty"`
}

func (a *Account) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	if a.Name != nil {
		nameB.Append(*a.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 1: Type.
	typeB := sb.FieldBuilder(1).(*array.StringBuilder)
	if a.Type != nil {
		typeB.Append(*a.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 2: TypeID.
	typeIDB := sb.FieldBuilder(2).(*array.Int32Builder)
	if a.TypeID != nil {
		typeIDB.Append(int32(*a.TypeID))
	} else {
		typeIDB.AppendNull()
	}

	// Field 3: UID.
	uidB := sb.FieldBuilder(3).(*array.StringBuilder)
	if a.UID != nil {
		uidB.Append(*a.UID)
	} else {
		uidB.AppendNull()
	}
}
