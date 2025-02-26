package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var DisplayFields = []arrow.Field{
	{Name: "color_depth", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_height", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_orientation", Type: arrow.PrimitiveTypes.Int32},
	{Name: "physical_width", Type: arrow.PrimitiveTypes.Int32},
	{Name: "scale_factor", Type: arrow.PrimitiveTypes.Int32},
}

// Display represents display specifications.
type Display struct {
	ColorDepth          *int `json:"color_depth,omitempty"`
	PhysicalHeight      *int `json:"physical_height,omitempty"`
	PhysicalOrientation *int `json:"physical_orientation,omitempty"`
	PhysicalWidth       *int `json:"physical_width,omitempty"`
	ScaleFactor         *int `json:"scale_factor,omitempty"`
}

// WriteToParquet writes the Display fields to the provided Arrow StructBuilder.
func (d *Display) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: ColorDepth.
	colorDepthB := sb.FieldBuilder(0).(*array.Int32Builder)
	if d.ColorDepth != nil {
		colorDepthB.Append(int32(*d.ColorDepth))
	} else {
		colorDepthB.AppendNull()
	}

	// Field 1: PhysicalHeight.
	physHeightB := sb.FieldBuilder(1).(*array.Int32Builder)
	if d.PhysicalHeight != nil {
		physHeightB.Append(int32(*d.PhysicalHeight))
	} else {
		physHeightB.AppendNull()
	}

	// Field 2: PhysicalOrientation.
	physOrientB := sb.FieldBuilder(2).(*array.Int32Builder)
	if d.PhysicalOrientation != nil {
		physOrientB.Append(int32(*d.PhysicalOrientation))
	} else {
		physOrientB.AppendNull()
	}

	// Field 3: PhysicalWidth.
	physWidthB := sb.FieldBuilder(3).(*array.Int32Builder)
	if d.PhysicalWidth != nil {
		physWidthB.Append(int32(*d.PhysicalWidth))
	} else {
		physWidthB.AppendNull()
	}

	// Field 4: ScaleFactor.
	scaleFactorB := sb.FieldBuilder(4).(*array.Int32Builder)
	if d.ScaleFactor != nil {
		scaleFactorB.Append(int32(*d.ScaleFactor))
	} else {
		scaleFactorB.AppendNull()
	}
}
