package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AuthFactorFields defines the Arrow fields for AuthFactor.
var AuthFactorFields = []arrow.Field{
	{Name: "device", Type: DeviceStruct},
	{Name: "email", Type: arrow.BinaryTypes.String},
	{Name: "factor_type", Type: arrow.BinaryTypes.String},
	{Name: "factor_type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "is_hotp", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "is_totp", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "phone_number", Type: arrow.BinaryTypes.String},
	{Name: "provider", Type: arrow.BinaryTypes.String},
	{Name: "security_questions", Type: arrow.ListOf(arrow.BinaryTypes.String)},
}

var AuthFactorStruct = arrow.StructOf(AuthFactorFields...)
var AuthFactorClassname = "auth_factor"

type AuthFactor struct {
	Device            *Device  `json:"device,omitempty" parquet:"device"`
	Email             *string  `json:"email,omitempty" parquet:"email"`
	FactorType        *string  `json:"factor_type,omitempty" parquet:"factor_type"`
	FactorTypeID      int32    `json:"factor_type_id" parquet:"factor_type_id"`
	IsHOTP            *bool    `json:"is_hotp,omitempty" parquet:"is_hotp"`
	IsTOTP            *bool    `json:"is_totp,omitempty" parquet:"is_totp"`
	PhoneNumber       *string  `json:"phone_number,omitempty" parquet:"phone_number"`
	Provider          *string  `json:"provider,omitempty" parquet:"provider"`
	SecurityQuestions []string `json:"security_questions,omitempty" parquet:"security_questions"`
}
