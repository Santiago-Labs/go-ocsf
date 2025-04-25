package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var EPSSFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "percentile", Type: arrow.PrimitiveTypes.Float64, Nullable: true},
	{Name: "score", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var EPSSStruct = arrow.StructOf(EPSSFields...)
var EPSSClassname = "epss"

type EPSS struct {
	CreatedTime   *int64   `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	CreatedTimeDt *string  `json:"created_time_dt" parquet:"created_time_dt,optional" ch:"created_time_dt"`
	Percentile    *float64 `json:"percentile" parquet:"percentile,optional" ch:"percentile"`
	Score         string   `json:"score" parquet:"score" ch:"score"`
	Version       *string  `json:"version" parquet:"version,optional" ch:"version"`
}
