package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// ProductFields defines the Arrow fields for Product.
var ProductFields = []arrow.Field{
	{Name: "cpe_name", Type: arrow.BinaryTypes.String},
	{Name: "feature", Type: FeatureStruct},
	{Name: "lang", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "url_string", Type: arrow.BinaryTypes.String},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var ProductStruct = arrow.StructOf(ProductFields...)
var ProductClassname = "product"

type Product struct {
	CPEName    *string  `json:"cpe_name,omitempty" parquet:"cpe_name"`
	Feature    *Feature `json:"feature,omitempty" parquet:"feature"`
	Lang       *string  `json:"lang,omitempty" parquet:"lang"`
	Name       *string  `json:"name,omitempty" parquet:"name"`
	Path       *string  `json:"path,omitempty" parquet:"path"`
	UID        *string  `json:"uid,omitempty" parquet:"uid"`
	URLString  *string  `json:"url_string,omitempty" parquet:"url_string"`
	VendorName string   `json:"vendor_name" parquet:"vendor_name"` // required field
	Version    *string  `json:"version,omitempty" parquet:"version"`
}
