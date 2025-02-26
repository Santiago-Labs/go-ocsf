package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var FeatureFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var FeatureSchema = arrow.NewSchema(FeatureFields, nil)

// Feature represents a feature with a name, UID, and version.
type Feature struct {
	Name    *string `json:"name,omitempty"`
	UID     *string `json:"uid,omitempty"`
	Version *string `json:"version,omitempty"`
}

// WriteToParquet writes the Feature fields to the provided Arrow StructBuilder.
func (f *Feature) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Name.
	nameB := sb.FieldBuilder(0).(*array.StringBuilder)
	if f.Name != nil {
		nameB.Append(*f.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 1: UID.
	uidB := sb.FieldBuilder(1).(*array.StringBuilder)
	if f.UID != nil {
		uidB.Append(*f.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 2: Version.
	versionB := sb.FieldBuilder(2).(*array.StringBuilder)
	if f.Version != nil {
		versionB.Append(*f.Version)
	} else {
		versionB.AppendNull()
	}
}
