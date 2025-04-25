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
	CPEName    *string  `json:"cpe_name" parquet:"cpe_name,optional" ch:"cpe_name" ch:"cpe_name"`
	Feature    *Feature `json:"feature" parquet:"feature,optional" ch:"feature"`
	Lang       *string  `json:"lang" parquet:"lang,optional" ch:"lang"`
	Name       *string  `json:"name" parquet:"name,optional" ch:"name"`
	Path       *string  `json:"path" parquet:"path,optional" ch:"path"`
	UID        *string  `json:"uid" parquet:"uid,optional" ch:"uid"`
	URLString  *string  `json:"url_string" parquet:"url_string,optional" ch:"url_string"`
	VendorName string   `json:"vendor_name" parquet:"vendor_name" ch:"vendor_name"` // required field
	Version    *string  `json:"version" parquet:"version,optional" ch:"version"`
}
