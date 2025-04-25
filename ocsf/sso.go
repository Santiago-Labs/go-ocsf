package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// SSOFields defines the Arrow fields for SSO.
var SSOFields = []arrow.Field{
	{Name: "auth_protocol", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "auth_protocol_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "certificate", Type: DigitalCertificateStruct, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "duration_mins", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "idle_timeout", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "login_endpoint", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "logout_endpoint", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "metadata_endpoint", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "scopes", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String, Nullable: true},
}

var SSOStruct = arrow.StructOf(SSOFields...)
var SSOClassname = "sso"

type SSO struct {
	AuthProtocol     *string             `json:"auth_protocol" parquet:"auth_protocol,optional" ch:"auth_protocol"`
	AuthProtocolID   *int32              `json:"auth_protocol_id" parquet:"auth_protocol_id,optional" ch:"auth_protocol_id"`
	Certificate      *DigitalCertificate `json:"certificate" parquet:"certificate,optional" ch:"certificate"`
	CreatedTime      *int64              `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	DurationMins     *int32              `json:"duration_mins" parquet:"duration_mins,optional" ch:"duration_mins"`
	IdleTimeout      *int32              `json:"idle_timeout" parquet:"idle_timeout,optional" ch:"idle_timeout"`
	LoginEndpoint    *string             `json:"login_endpoint" parquet:"login_endpoint,optional" ch:"login_endpoint"`
	LogoutEndpoint   *string             `json:"logout_endpoint" parquet:"logout_endpoint,optional" ch:"logout_endpoint"`
	MetadataEndpoint *string             `json:"metadata_endpoint" parquet:"metadata_endpoint,optional" ch:"metadata_endpoint"`
	ModifiedTime     *int64              `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	Name             *string             `json:"name" parquet:"name,optional" ch:"name"`
	ProtocolName     *string             `json:"protocol_name" parquet:"protocol_name,optional" ch:"protocol_name"`
	Scopes           []string            `json:"scopes" parquet:"scopes,list,optional" ch:"scopes"`
	UID              *string             `json:"uid" parquet:"uid,optional" ch:"uid"`
	VendorName       *string             `json:"vendor_name" parquet:"vendor_name,optional" ch:"vendor_name"`
}
