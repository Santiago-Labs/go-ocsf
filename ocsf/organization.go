package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// OrganizationFields defines the Arrow fields for Organization.
var OrganizationFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ou_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ou_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var OrganizationStruct = arrow.StructOf(OrganizationFields...)
var OrganizationClassname = "organization"

type Organization struct {
	Name   *string `json:"name,omitempty" parquet:"name,optional"`
	OUName *string `json:"ou_name,omitempty" parquet:"ou_name,optional"`
	OUID   *string `json:"ou_uid,omitempty" parquet:"ou_uid,optional"`
	UID    *string `json:"uid,omitempty" parquet:"uid,optional"`
}
