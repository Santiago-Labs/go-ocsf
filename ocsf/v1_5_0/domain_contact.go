// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type DomainContact struct {

	// Contact Email: The user's primary email address.
	EmailAddr *string `json:"email_addr,omitempty" parquet:"email_addr,optional"`

	// Contact Location Information: Location details for the contract such as the city, state/province, country, etc.
	Location *GeoLocation `json:"location,omitempty" parquet:"location,optional"`

	// Name: The individual or organization name for the contact.
	Name *string `json:"name,omitempty" parquet:"name,optional"`

	// Phone Number: The number associated with the phone.
	PhoneNumber *string `json:"phone_number,omitempty" parquet:"phone_number,optional"`

	// Domain Contact Type: The Domain Contact type, normalized to the caption of the <code>type_id</code> value. In the case of 'Other', it is defined by the source
	Type *string `json:"type,omitempty" parquet:"type,optional"`

	// Domain Contact Type ID: The normalized domain contact type ID.
	TypeId int32 `json:"type_id" parquet:"type_id"`

	// Unique ID: The unique identifier of the contact information, typically provided in WHOIS information.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`
}

var DomainContactFields = []arrow.Field{
	{Name: "email_addr", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "location", Type: GeoLocationStruct, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "phone_number", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DomainContactStruct = arrow.StructOf(DomainContactFields...)

var DomainContactSchema = arrow.NewSchema(DomainContactFields, nil)
