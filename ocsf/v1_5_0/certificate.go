// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type DigitalCertificate struct {

	// Created Time: The time when the certificate was created.
	CreatedTime *int64 `json:"created_time,omitempty" parquet:"created_time,optional"`

	// Expiration Time: The expiration time of the certificate.
	ExpirationTime *int64 `json:"expiration_time,omitempty" parquet:"expiration_time,optional"`

	// Fingerprints: The fingerprint list of the certificate.
	Fingerprints []*Fingerprint `json:"fingerprints,omitempty" parquet:"fingerprints,optional,list"`

	// Certificate Self-Signed: Denotes whether a digital certificate is self-signed or signed by a known certificate authority (CA).
	IsSelfSigned *bool `json:"is_self_signed,omitempty" parquet:"is_self_signed,optional"`

	// Issuer Distinguished Name: The certificate issuer distinguished name.
	Issuer string `json:"issuer" parquet:"issuer"`

	// Subject Alternative Names: The list of subject alternative names that are secured by a specific certificate.
	Sans []*SubjectAlternativeName `json:"sans,omitempty" parquet:"sans,optional,list"`

	// Certificate Serial Number: The serial number of the certificate used to create the digital signature.
	SerialNumber string `json:"serial_number" parquet:"serial_number"`

	// Subject Distinguished Name: The certificate subject distinguished name.
	Subject *string `json:"subject,omitempty" parquet:"subject,optional"`

	// Unique ID: The unique identifier of the certificate.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`

	// Version: The certificate version.
	Version *string `json:"version,omitempty" parquet:"version,optional"`
}

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

var DigitalCertificateSchema = arrow.NewSchema(DigitalCertificateFields, nil)
