package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var GeoLocationFields = []arrow.Field{
	{Name: "city", Type: arrow.BinaryTypes.String},
	{Name: "continent", Type: arrow.BinaryTypes.String},
	{Name: "coordinates", Type: arrow.ListOf(arrow.PrimitiveTypes.Float64)},
	{Name: "country", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "is_on_premises", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "isp", Type: arrow.BinaryTypes.String},
	{Name: "postal_code", Type: arrow.BinaryTypes.String},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "region", Type: arrow.BinaryTypes.String},
}

// GeoLocation represents geographic location details.
type GeoLocation struct {
	City         *string   `json:"city,omitempty"`
	Continent    *string   `json:"continent,omitempty"`
	Coordinates  []float64 `json:"coordinates,omitempty"`
	Country      *string   `json:"country,omitempty"`
	Desc         *string   `json:"desc,omitempty"`
	IsOnPremises *bool     `json:"is_on_premises,omitempty"`
	ISP          *string   `json:"isp,omitempty"`
	PostalCode   *string   `json:"postal_code,omitempty"`
	Provider     *string   `json:"provider,omitempty"`
	Region       *string   `json:"region,omitempty"`
}

// WriteToParquet writes the GeoLocation fields to the provided Arrow StructBuilder.
// It assumes the schema is defined as described above.
func (g *GeoLocation) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: City.
	cityB := sb.FieldBuilder(0).(*array.StringBuilder)
	if g.City != nil {
		cityB.Append(*g.City)
	} else {
		cityB.AppendNull()
	}

	// Field 1: Continent.
	continentB := sb.FieldBuilder(1).(*array.StringBuilder)
	if g.Continent != nil {
		continentB.Append(*g.Continent)
	} else {
		continentB.AppendNull()
	}

	// Field 2: Coordinates (list of float64).
	coordListB := sb.FieldBuilder(2).(*array.ListBuilder)
	if len(g.Coordinates) > 0 {
		coordListB.Append(true)
		coordValB := coordListB.ValueBuilder().(*array.Float64Builder)
		for _, coord := range g.Coordinates {
			coordValB.Append(coord)
		}
	} else {
		coordListB.AppendNull()
	}

	// Field 3: Country.
	countryB := sb.FieldBuilder(3).(*array.StringBuilder)
	if g.Country != nil {
		countryB.Append(*g.Country)
	} else {
		countryB.AppendNull()
	}

	// Field 4: Desc.
	descB := sb.FieldBuilder(4).(*array.StringBuilder)
	if g.Desc != nil {
		descB.Append(*g.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 5: IsOnPremises.
	onPremB := sb.FieldBuilder(5).(*array.BooleanBuilder)
	if g.IsOnPremises != nil {
		onPremB.Append(*g.IsOnPremises)
	} else {
		onPremB.AppendNull()
	}

	// Field 6: ISP.
	ispB := sb.FieldBuilder(6).(*array.StringBuilder)
	if g.ISP != nil {
		ispB.Append(*g.ISP)
	} else {
		ispB.AppendNull()
	}

	// Field 7: PostalCode.
	postalB := sb.FieldBuilder(7).(*array.StringBuilder)
	if g.PostalCode != nil {
		postalB.Append(*g.PostalCode)
	} else {
		postalB.AppendNull()
	}

	// Field 8: Provider.
	providerB := sb.FieldBuilder(8).(*array.StringBuilder)
	if g.Provider != nil {
		providerB.Append(*g.Provider)
	} else {
		providerB.AppendNull()
	}

	// Field 9: Region.
	regionB := sb.FieldBuilder(9).(*array.StringBuilder)
	if g.Region != nil {
		regionB.Append(*g.Region)
	} else {
		regionB.AppendNull()
	}
}
