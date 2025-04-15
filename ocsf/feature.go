package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var FeatureFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var FeatureStruct = arrow.StructOf(FeatureFields...)
var FeatureClassname = "feature"

type Feature struct {
	Name    *string `json:"name,omitempty" parquet:"name"`
	UID     *string `json:"uid,omitempty" parquet:"uid"`
	Version *string `json:"version,omitempty" parquet:"version"`
}
