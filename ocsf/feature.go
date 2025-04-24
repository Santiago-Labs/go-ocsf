package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var FeatureFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var FeatureStruct = arrow.StructOf(FeatureFields...)
var FeatureClassname = "feature"

type Feature struct {
	Name    *string `json:"name,omitempty" parquet:"name,optional" ch:"name"`
	UID     *string `json:"uid,omitempty" parquet:"uid,optional" ch:"uid"`
	Version *string `json:"version,omitempty" parquet:"version,optional" ch:"version"`
}
