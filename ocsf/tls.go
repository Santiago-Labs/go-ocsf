package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var TLSFields = []arrow.Field{
	{Name: "alert", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "certificate", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "certification_chain", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "cipher", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "client_ciphers", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "handshake_dur", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "ja3_hash", Type: FingerprintStruct, Nullable: true},
	{Name: "ja3s_hash", Type: FingerprintStruct, Nullable: true},
	{Name: "key_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "server_ciphers", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "tls_extension_list", Type: arrow.ListOf(TLSExtensionStruct), Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var TLSStruct = arrow.StructOf(TLSFields...)
var TLSClassname = "tls"

type TLS struct {
	Alert            *int32              `json:"alert,omitempty" parquet:"alert,optional"`
	Certificate      *DigitalCertificate `json:"certificate,omitempty" parquet:"certificate,optional"`
	CertificateChain []*string           `json:"certificate_chain,omitempty" parquet:"certification_chain,list,optional"`
	Cipher           *string             `json:"cipher,omitempty" parquet:"cipher,optional"`
	ClientCiphers    []*string           `json:"client_ciphers,omitempty" parquet:"client_ciphers,list,optional"`
	HandshakeDur     *int64              `json:"handshake_dur,omitempty" parquet:"handshake_dur,optional"`
	JA3Hash          *Fingerprint        `json:"ja3_hash,omitempty" parquet:"ja3_hash,optional"`
	JA3SHash         *Fingerprint        `json:"ja3s_hash,omitempty" parquet:"ja3s_hash,optional"`
	KeyLength        *int32              `json:"key_length,omitempty" parquet:"key_length,optional"`
	ServerCiphers    []*string           `json:"server_ciphers,omitempty" parquet:"server_ciphers,list,optional"`
	TLSExtensionList []*TLSExtension     `json:"tls_extension_list,omitempty" parquet:"tls_extension_list,list,optional"`
	Version          *string             `json:"version,omitempty" parquet:"version,optional"`
}
