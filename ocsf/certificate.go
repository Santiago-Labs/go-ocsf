package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// DigitalCertificateFields defines the Arrow fields for Digital Certificate.
var DigitalCertificateFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "expiration_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "fingerprints", Type: arrow.ListOf(FingerprintStruct)},
	{Name: "is_self_signed", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "issuer", Type: arrow.BinaryTypes.String},
	{Name: "sans", Type: arrow.ListOf(SubjectAlternativeNameStruct)},
	{Name: "serial_number", Type: arrow.BinaryTypes.String},
	{Name: "subject", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var DigitalCertificateStruct = arrow.StructOf(DigitalCertificateFields...)
var DigitalCertificateClassname = "certificate"

type DigitalCertificate struct {
	CreatedTime    *time.Time                `json:"created_time,omitempty" parquet:"created_time"`
	ExpirationTime *time.Time                `json:"expiration_time,omitempty" parquet:"expiration_time"`
	Fingerprints   []*Fingerprint            `json:"fingerprints,omitempty" parquet:"fingerprints"`
	IsSelfSigned   *bool                     `json:"is_self_signed,omitempty" parquet:"is_self_signed"`
	Issuer         string                    `json:"issuer" parquet:"issuer"`
	SANs           []*SubjectAlternativeName `json:"sans,omitempty" parquet:"sans"`
	SerialNumber   string                    `json:"serial_number" parquet:"serial_number"`
	Subject        *string                   `json:"subject,omitempty" parquet:"subject"`
	UID            *string                   `json:"uid,omitempty" parquet:"uid"`
	Version        *string                   `json:"version,omitempty" parquet:"version"`
}
