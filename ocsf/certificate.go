package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// DigitalCertificateFields defines the Arrow fields for Digital Certificate.
var DigitalCertificateFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "expiration_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "fingerprints", Type: arrow.ListOf(FingerprintStruct), Nullable: true},
	{Name: "is_self_signed", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "issuer", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "sans", Type: arrow.ListOf(SubjectAlternativeNameStruct), Nullable: true},
	{Name: "serial_number", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "subject", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DigitalCertificateStruct = arrow.StructOf(DigitalCertificateFields...)
var DigitalCertificateClassname = "certificate"

type DigitalCertificate struct {
	CreatedTime    *int64                    `json:"created_time,omitempty" parquet:"created_time,optional"`
	ExpirationTime *int64                    `json:"expiration_time,omitempty" parquet:"expiration_time,optional"`
	Fingerprints   []*Fingerprint            `json:"fingerprints,omitempty" parquet:"fingerprints,list,optional"`
	IsSelfSigned   *bool                     `json:"is_self_signed,omitempty" parquet:"is_self_signed,optional"`
	Issuer         string                    `json:"issuer" parquet:"issuer"`
	SANs           []*SubjectAlternativeName `json:"sans,omitempty" parquet:"sans,list,optional"`
	SerialNumber   string                    `json:"serial_number" parquet:"serial_number"`
	Subject        *string                   `json:"subject,omitempty" parquet:"subject,optional"`
	UID            *string                   `json:"uid,omitempty" parquet:"uid,optional"`
	Version        *string                   `json:"version,omitempty" parquet:"version,optional"`
}
