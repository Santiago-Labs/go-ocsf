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
	City         *string    `json:"city,omitempty" parquet:"city,optional"`
	Continent    *string    `json:"continent,omitempty" parquet:"continent,optional"`
	Coordinates  []*float64 `json:"coordinates,omitempty" parquet:"coordinates,list,optional"`
	Country      *string    `json:"country,omitempty" parquet:"country,optional"`
	Desc         *string    `json:"desc,omitempty" parquet:"desc,optional"`
	IsOnPremises *bool      `json:"is_on_premises,omitempty" parquet:"is_on_premises,optional"`
	ISP          *string    `json:"isp,omitempty" parquet:"isp,optional"`
	PostalCode   *string    `json:"postal_code,omitempty" parquet:"postal_code,optional"`
	Provider     *string    `json:"provider,omitempty" parquet:"provider,optional"`
	Region       *string    `json:"region,omitempty" parquet:"region,optional"`
}
