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
	CreatedTime    *int64                    `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	ExpirationTime *int64                    `json:"expiration_time" parquet:"expiration_time,optional" ch:"expiration_time"`
	Fingerprints   []*Fingerprint            `json:"fingerprints" parquet:"fingerprints,list,optional" ch:"fingerprints"`
	IsSelfSigned   *bool                     `json:"is_self_signed" parquet:"is_self_signed,optional" ch:"is_self_signed"`
	Issuer         string                    `json:"issuer" parquet:"issuer" ch:"issuer"`
	SANs           []*SubjectAlternativeName `json:"sans" parquet:"sans,list,optional" ch:"sans"`
	SerialNumber   string                    `json:"serial_number" parquet:"serial_number" ch:"serial_number"`
	Subject        *string                   `json:"subject" parquet:"subject,optional" ch:"subject"`
	UID            *string                   `json:"uid" parquet:"uid,optional" ch:"uid"`
	Version        *string                   `json:"version" parquet:"version,optional" ch:"version"`
}
