package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AuthFactorFields defines the Arrow fields for AuthFactor.
var AuthFactorFields = []arrow.Field{
	{Name: "device", Type: DeviceStruct, Nullable: true},
	{Name: "email", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "factor_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "factor_type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "is_hotp", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_totp", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "phone_number", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "security_questions", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
}

var AuthFactorStruct = arrow.StructOf(AuthFactorFields...)
var AuthFactorClassname = "auth_factor"

type AuthFactor struct {
	Device            *Device   `json:"device,omitempty" parquet:"device,optional"`
	Email             *string   `json:"email,omitempty" parquet:"email,optional"`
	FactorType        *string   `json:"factor_type,omitempty" parquet:"factor_type,optional"`
	FactorTypeID      int32     `json:"factor_type_id" parquet:"factor_type_id"`
	IsHOTP            *bool     `json:"is_hotp,omitempty" parquet:"is_hotp,optional"`
	IsTOTP            *bool     `json:"is_totp,omitempty" parquet:"is_totp,optional"`
	PhoneNumber       *string   `json:"phone_number,omitempty" parquet:"phone_number,optional"`
	Provider          *string   `json:"provider,omitempty" parquet:"provider,optional"`
	SecurityQuestions []*string `json:"security_questions,omitempty" parquet:"security_questions,list,optional"`
}
