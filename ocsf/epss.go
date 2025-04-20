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
	CreatedTime   *int     `json:"created_time,omitempty" parquet:"created_time,optional"`
	CreatedTimeDt *string  `json:"created_time_dt,omitempty" parquet:"created_time_dt,optional"`
	Percentile    *float64 `json:"percentile,omitempty" parquet:"percentile,optional"`
	Score         string   `json:"score" parquet:"score"`
	Version       *string  `json:"version,omitempty" parquet:"version,optional"`
}
