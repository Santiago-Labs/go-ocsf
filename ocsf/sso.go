package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// SSOFields defines the Arrow fields for SSO.
var SSOFields = []arrow.Field{
	{Name: "auth_protocol", Type: arrow.BinaryTypes.String},
	{Name: "auth_protocol_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "certificate", Type: DigitalCertificateStruct},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "duration_mins", Type: arrow.PrimitiveTypes.Int32},
	{Name: "idle_timeout", Type: arrow.PrimitiveTypes.Int32},
	{Name: "login_endpoint", Type: arrow.BinaryTypes.String},
	{Name: "logout_endpoint", Type: arrow.BinaryTypes.String},
	{Name: "metadata_endpoint", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String},
	{Name: "scopes", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String},
}

var SSOStruct = arrow.StructOf(SSOFields...)
var SSOClassname = "sso"

type SSO struct {
	AuthProtocol     *string             `json:"auth_protocol,omitempty" parquet:"auth_protocol"`
	AuthProtocolID   *int32              `json:"auth_protocol_id,omitempty" parquet:"auth_protocol_id"`
	Certificate      *DigitalCertificate `json:"certificate,omitempty" parquet:"certificate"`
	CreatedTime      *time.Time          `json:"created_time,omitempty" parquet:"created_time"`
	DurationMins     *int32              `json:"duration_mins,omitempty" parquet:"duration_mins"`
	IdleTimeout      *int32              `json:"idle_timeout,omitempty" parquet:"idle_timeout"`
	LoginEndpoint    *string             `json:"login_endpoint,omitempty" parquet:"login_endpoint"`
	LogoutEndpoint   *string             `json:"logout_endpoint,omitempty" parquet:"logout_endpoint"`
	MetadataEndpoint *string             `json:"metadata_endpoint,omitempty" parquet:"metadata_endpoint"`
	ModifiedTime     *time.Time          `json:"modified_time,omitempty" parquet:"modified_time"`
	Name             *string             `json:"name,omitempty" parquet:"name"`
	ProtocolName     *string             `json:"protocol_name,omitempty" parquet:"protocol_name"`
	Scopes           []string            `json:"scopes,omitempty" parquet:"scopes"`
	UID              *string             `json:"uid,omitempty" parquet:"uid"`
	VendorName       *string             `json:"vendor_name,omitempty" parquet:"vendor_name"`
}
