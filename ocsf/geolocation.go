package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var GeoLocationFields = []arrow.Field{
	{Name: "city", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "continent", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "coordinates", Type: arrow.ListOf(arrow.PrimitiveTypes.Float64), Nullable: true},
	{Name: "country", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "is_on_premises", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "isp", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "postal_code", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "region", Type: arrow.BinaryTypes.String, Nullable: true},
}

var GeoLocationStruct = arrow.StructOf(GeoLocationFields...)
var GeoLocationClassname = "location"

type GeoLocation struct {
	City         *string   `json:"city,omitempty" parquet:"city"`
	Continent    *string   `json:"continent,omitempty" parquet:"continent"`
	Coordinates  []float64 `json:"coordinates,omitempty" parquet:"coordinates"`
	Country      *string   `json:"country,omitempty" parquet:"country"`
	Desc         *string   `json:"desc,omitempty" parquet:"desc"`
	IsOnPremises *bool     `json:"is_on_premises,omitempty" parquet:"is_on_premises"`
	ISP          *string   `json:"isp,omitempty" parquet:"isp"`
	PostalCode   *string   `json:"postal_code,omitempty" parquet:"postal_code"`
	Provider     *string   `json:"provider,omitempty" parquet:"provider"`
	Region       *string   `json:"region,omitempty" parquet:"region"`
}
