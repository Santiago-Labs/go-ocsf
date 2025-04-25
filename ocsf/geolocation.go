package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
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
	City         *string    `json:"city" parquet:"city,optional" ch:"city" ch:"city"`
	Continent    *string    `json:"continent" parquet:"continent,optional" ch:"continent"`
	Coordinates  []*float64 `json:"coordinates" parquet:"coordinates,list,optional" ch:"coordinates"`
	Country      *string    `json:"country" parquet:"country,optional" ch:"country"`
	Desc         *string    `json:"desc" parquet:"desc,optional" ch:"desc"`
	IsOnPremises *bool      `json:"is_on_premises" parquet:"is_on_premises,optional" ch:"is_on_premises"`
	ISP          *string    `json:"isp" parquet:"isp,optional" ch:"isp"`
	PostalCode   *string    `json:"postal_code" parquet:"postal_code,optional" ch:"postal_code"`
	Provider     *string    `json:"provider" parquet:"provider,optional" ch:"provider"`
	Region       *string    `json:"region" parquet:"region,optional" ch:"region"`
}
