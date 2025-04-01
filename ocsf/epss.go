package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var EPSSFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "percentile", Type: arrow.PrimitiveTypes.Float64},
	{Name: "score", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var EPSSStruct = arrow.StructOf(EPSSFields...)
var EPSSSchema = arrow.NewSchema(EPSSFields, nil)

type EPSS struct {
	CreatedTime   *int     `json:"created_time,omitempty" parquet:"created_time"`
	CreatedTimeDt *string  `json:"created_time_dt,omitempty" parquet:"created_time_dt"`
	Percentile    *float64 `json:"percentile,omitempty" parquet:"percentile"`
	Score         string   `json:"score" parquet:"score"`
	Version       *string  `json:"version,omitempty" parquet:"version"`
}
