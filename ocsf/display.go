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
	ColorDepth          *int64 `json:"color_depth" parquet:"color_depth,optional" ch:"color_depth"`
	PhysicalHeight      *int64 `json:"physical_height" parquet:"physical_height,optional" ch:"physical_height"`
	PhysicalOrientation *int64 `json:"physical_orientation" parquet:"physical_orientation,optional" ch:"physical_orientation"`
	PhysicalWidth       *int64 `json:"physical_width" parquet:"physical_width,optional" ch:"physical_width"`
	ScaleFactor         *int64 `json:"scale_factor" parquet:"scale_factor,optional" ch:"scale_factor"`
}
