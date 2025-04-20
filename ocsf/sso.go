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
	AuthProtocol     *string             `json:"auth_protocol,omitempty" parquet:"auth_protocol,optional"`
	AuthProtocolID   *int32              `json:"auth_protocol_id,omitempty" parquet:"auth_protocol_id,optional"`
	Certificate      *DigitalCertificate `json:"certificate,omitempty" parquet:"certificate,optional"`
	CreatedTime      *int64              `json:"created_time,omitempty" parquet:"created_time,optional"`
	DurationMins     *int32              `json:"duration_mins,omitempty" parquet:"duration_mins,optional"`
	IdleTimeout      *int32              `json:"idle_timeout,omitempty" parquet:"idle_timeout,optional"`
	LoginEndpoint    *string             `json:"login_endpoint,omitempty" parquet:"login_endpoint,optional"`
	LogoutEndpoint   *string             `json:"logout_endpoint,omitempty" parquet:"logout_endpoint,optional"`
	MetadataEndpoint *string             `json:"metadata_endpoint,omitempty" parquet:"metadata_endpoint,optional"`
	ModifiedTime     *int64              `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Name             *string             `json:"name,omitempty" parquet:"name,optional"`
	ProtocolName     *string             `json:"protocol_name,omitempty" parquet:"protocol_name,optional"`
	Scopes           []*string           `json:"scopes,omitempty" parquet:"scopes,list,optional"`
	UID              *string             `json:"uid,omitempty" parquet:"uid,optional"`
	VendorName       *string             `json:"vendor_name,omitempty" parquet:"vendor_name,optional"`
}
