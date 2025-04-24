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
	City         *string    `json:"city,omitempty" parquet:"city,optional" ch:"city,omitempty" ch:"city,omitempty"`
	Continent    *string    `json:"continent,omitempty" parquet:"continent,optional" ch:"continent,omitempty"`
	Coordinates  []*float64 `json:"coordinates,omitempty" parquet:"coordinates,list,optional" ch:"coordinates,omitempty"`
	Country      *string    `json:"country,omitempty" parquet:"country,optional" ch:"country,omitempty"`
	Desc         *string    `json:"desc,omitempty" parquet:"desc,optional" ch:"desc,omitempty"`
	IsOnPremises *bool      `json:"is_on_premises,omitempty" parquet:"is_on_premises,optional" ch:"is_on_premises,omitempty"`
	ISP          *string    `json:"isp,omitempty" parquet:"isp,optional" ch:"isp,omitempty"`
	PostalCode   *string    `json:"postal_code,omitempty" parquet:"postal_code,optional" ch:"postal_code,omitempty"`
	Provider     *string    `json:"provider,omitempty" parquet:"provider,optional" ch:"provider,omitempty"`
	Region       *string    `json:"region,omitempty" parquet:"region,optional" ch:"region,omitempty"`
}
