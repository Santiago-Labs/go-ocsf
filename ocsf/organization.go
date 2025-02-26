package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// OrganizationFields defines the Arrow fields for Organization.
var OrganizationFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "ou_name", Type: arrow.BinaryTypes.String},
	{Name: "ou_uid", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// OrganizationSchema is the Arrow schema for Organization.
var OrganizationSchema = arrow.NewSchema(OrganizationFields, nil)

// Organization represents an organization.
type Organization struct {
	Name   *string `json:"name,omitempty"`
	OUName *string `json:"ou_name,omitempty"`
	OUID   *string `json:"ou_uid,omitempty"`
	UID    *string `json:"uid,omitempty"`
}

// WriteToParquet writes the Organization fields to the provided Arrow StructBuilder.
func (o *Organization) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	if o.Name != nil {
		nameB.Append(*o.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 1: ou_name.
	ouNameB := sb.FieldBuilder(1).(*array.StringBuilder)
	if o.OUName != nil {
		ouNameB.Append(*o.OUName)
	} else {
		ouNameB.AppendNull()
	}

	// Field 2: ou_uid.
	ouidB := sb.FieldBuilder(2).(*array.StringBuilder)
	if o.OUID != nil {
		ouidB.Append(*o.OUID)
	} else {
		ouidB.AppendNull()
	}

	// Field 3: uid.
	uidB := sb.FieldBuilder(3).(*array.StringBuilder)
	if o.UID != nil {
		uidB.Append(*o.UID)
	} else {
		uidB.AppendNull()
	}
}
