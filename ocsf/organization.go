package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// OrganizationFields defines the Arrow fields for Organization.
var OrganizationFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "ou_name", Type: arrow.BinaryTypes.String},
	{Name: "ou_uid", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

var OrganizationStruct = arrow.StructOf(OrganizationFields...)
var OrganizationClassname = "organization"

type Organization struct {
	Name   *string `json:"name,omitempty" parquet:"name"`
	OUName *string `json:"ou_name,omitempty" parquet:"ou_name"`
	OUID   *string `json:"ou_uid,omitempty" parquet:"ou_uid"`
	UID    *string `json:"uid,omitempty" parquet:"uid"`
}
