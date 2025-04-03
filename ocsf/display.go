package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var DisplayFields = []arrow.Field{
	{Name: "color_depth", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_height", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_orientation", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_width", Type: arrow.PrimitiveTypes.Int32},
	{Name: "scale_factor", Type: arrow.PrimitiveTypes.Int32},
}

var DisplayStruct = arrow.StructOf(DisplayFields...)
var DisplayClassname = "display"

// Display represents display specifications.
type Display struct {
	ColorDepth          *int `json:"color_depth,omitempty" parquet:"color_depth"`
	PhysicalHeight      *int `json:"physical_height,omitempty" parquet:"physical_height"`
	PhysicalOrientation *int `json:"physical_orientation,omitempty" parquet:"physical_orientation"`
	PhysicalWidth       *int `json:"physical_width,omitempty" parquet:"physical_width"`
	ScaleFactor         *int `json:"scale_factor,omitempty" parquet:"scale_factor"`
}
