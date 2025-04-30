package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// VendorAttributesFields defines the Arrow fields for VendorAttributes.
var VendorAttributesFields = []arrow.Field{
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
}

var VendorAttributesStruct = arrow.StructOf(VendorAttributesFields...)
var VendorAttributesClassname = "vendor_attributes"

type VendorAttributes struct {
	Severity   *string `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID *int32  `json:"severity_id,omitempty" parquet:"severity_id,optional"`
}
