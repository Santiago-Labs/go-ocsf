package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// ProductFields defines the Arrow fields for Product.
var ProductFields = []arrow.Field{
	{Name: "cpe_name", Type: arrow.BinaryTypes.String},
	{Name: "feature", Type: arrow.StructOf(FeatureFields...)},
	{Name: "lang", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "url_string", Type: arrow.BinaryTypes.String},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// ProductSchema is the Arrow schema for Product.
var ProductSchema = arrow.NewSchema(ProductFields, nil)

// Product represents a product.
type Product struct {
	CPEName    *string  `json:"cpe_name,omitempty"`
	Feature    *Feature `json:"feature,omitempty"`
	Lang       *string  `json:"lang,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Path       *string  `json:"path,omitempty"`
	UID        *string  `json:"uid,omitempty"`
	URLString  *string  `json:"url_string,omitempty"`
	VendorName string   `json:"vendor_name"` // required field
	Version    *string  `json:"version,omitempty"`
}

// WriteToParquet writes the Product fields to the provided Arrow StructBuilder.
func (p *Product) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: CPEName.
	cpeNameB := sb.FieldBuilder(0).(*array.StringBuilder)
	if p.CPEName != nil {
		cpeNameB.Append(*p.CPEName)
	} else {
		cpeNameB.AppendNull()
	}

	// Field 1: Feature (nested struct).
	featureB := sb.FieldBuilder(1).(*array.StructBuilder)
	if p.Feature != nil {
		featureB.Append(true)
		p.Feature.WriteToParquet(featureB)
	} else {
		featureB.AppendNull()
	}

	// Field 2: Lang.
	langB := sb.FieldBuilder(2).(*array.StringBuilder)
	if p.Lang != nil {
		langB.Append(*p.Lang)
	} else {
		langB.AppendNull()
	}

	// Field 3: Name.
	nameB := sb.FieldBuilder(3).(*array.StringBuilder)
	if p.Name != nil {
		nameB.Append(*p.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 4: Path.
	pathB := sb.FieldBuilder(4).(*array.StringBuilder)
	if p.Path != nil {
		pathB.Append(*p.Path)
	} else {
		pathB.AppendNull()
	}

	// Field 5: UID.
	uidB := sb.FieldBuilder(5).(*array.StringBuilder)
	if p.UID != nil {
		uidB.Append(*p.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 6: URLString.
	urlStringB := sb.FieldBuilder(6).(*array.StringBuilder)
	if p.URLString != nil {
		urlStringB.Append(*p.URLString)
	} else {
		urlStringB.AppendNull()
	}

	// Field 7: VendorName (required).
	vendorNameB := sb.FieldBuilder(7).(*array.StringBuilder)
	vendorNameB.Append(p.VendorName)

	// Field 8: Version.
	versionB := sb.FieldBuilder(8).(*array.StringBuilder)
	if p.Version != nil {
		versionB.Append(*p.Version)
	} else {
		versionB.AppendNull()
	}
}
