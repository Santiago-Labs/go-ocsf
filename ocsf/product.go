package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ProductFields defines the Arrow fields for Product.
var ProductFields = []arrow.Field{
	{Name: "cpe_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "feature", Type: FeatureStruct, Nullable: true},
	{Name: "lang", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "url_string", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ProductStruct = arrow.StructOf(ProductFields...)
var ProductClassname = "product"

type Product struct {
	CPEName    *string  `json:"cpe_name,omitempty" parquet:"cpe_name,optional" ch:"cpe_name,omitempty" ch:"cpe_name,omitempty"`
	Feature    *Feature `json:"feature,omitempty" parquet:"feature,optional" ch:"feature,omitempty"`
	Lang       *string  `json:"lang,omitempty" parquet:"lang,optional" ch:"lang,omitempty"`
	Name       *string  `json:"name,omitempty" parquet:"name,optional" ch:"name,omitempty"`
	Path       *string  `json:"path,omitempty" parquet:"path,optional" ch:"path,omitempty"`
	UID        *string  `json:"uid,omitempty" parquet:"uid,optional" ch:"uid,omitempty"`
	URLString  *string  `json:"url_string,omitempty" parquet:"url_string,optional" ch:"url_string,omitempty"`
	VendorName string   `json:"vendor_name" parquet:"vendor_name" ch:"vendor_name"` // required field
	Version    *string  `json:"version,omitempty" parquet:"version,optional" ch:"version,omitempty"`
}
