package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var FeatureFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var FeatureSchema = arrow.NewSchema(FeatureFields, nil)

type Feature struct {
	Name    *string `json:"name,omitempty" parquet:"name"`
	UID     *string `json:"uid,omitempty" parquet:"uid"`
	Version *string `json:"version,omitempty" parquet:"version"`
}
