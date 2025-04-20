package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var DisplayFields = []arrow.Field{
	{Name: "color_depth", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "physical_height", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "physical_orientation", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "physical_width", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "scale_factor", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
}

var DisplayStruct = arrow.StructOf(DisplayFields...)
var DisplayClassname = "display"

// Display represents display specifications.
type Display struct {
	ColorDepth          *int `json:"color_depth,omitempty" parquet:"color_depth,optional"`
	PhysicalHeight      *int `json:"physical_height,omitempty" parquet:"physical_height,optional"`
	PhysicalOrientation *int `json:"physical_orientation,omitempty" parquet:"physical_orientation,optional"`
	PhysicalWidth       *int `json:"physical_width,omitempty" parquet:"physical_width,optional"`
	ScaleFactor         *int `json:"scale_factor,omitempty" parquet:"scale_factor,optional"`
}
