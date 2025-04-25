package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// SCIMFields defines the Arrow fields for SCIM.
var SCIMFields = []arrow.Field{
	{Name: "auth_protocol", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "auth_protocol_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "error_message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "is_group_provisioning_enabled", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_user_provisioning_enabled", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "last_run_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "rate_limit", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "scim_group_schema", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "scim_user_schema", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "state", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "state_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "url_string", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var SCIMStruct = arrow.StructOf(SCIMFields...)
var SCIMClassname = "scim"

type SCIM struct {
	AuthProtocol               *string `json:"auth_protocol" parquet:"auth_protocol,optional" ch:"auth_protocol" ch:"auth_protocol"`
	AuthProtocolID             *int32  `json:"auth_protocol_id" parquet:"auth_protocol_id,optional" ch:"auth_protocol_id"`
	CreatedTime                *int64  `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	ErrorMessage               *string `json:"error_message" parquet:"error_message,optional" ch:"error_message"`
	IsGroupProvisioningEnabled *bool   `json:"is_group_provisioning_enabled" parquet:"is_group_provisioning_enabled,optional" ch:"is_group_provisioning_enabled"`
	IsUserProvisioningEnabled  *bool   `json:"is_user_provisioning_enabled" parquet:"is_user_provisioning_enabled,optional" ch:"is_user_provisioning_enabled"`
	LastRunTime                *int64  `json:"last_run_time" parquet:"last_run_time,optional" ch:"last_run_time"`
	ModifiedTime               *int64  `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	Name                       *string `json:"name" parquet:"name,optional" ch:"name"`
	ProtocolName               *string `json:"protocol_name" parquet:"protocol_name,optional" ch:"protocol_name"`
	RateLimit                  *int32  `json:"rate_limit" parquet:"rate_limit,optional" ch:"rate_limit"`
	SCIMGroupSchema            *string `json:"scim_group_schema" parquet:"scim_group_schema,optional" ch:"scim_group_schema"`
	SCIMUserSchema             *string `json:"scim_user_schema" parquet:"scim_user_schema,optional" ch:"scim_user_schema"`
	State                      *string `json:"state" parquet:"state,optional" ch:"state"`
	StateID                    *int32  `json:"state_id" parquet:"state_id,optional" ch:"state_id"`
	UID                        *string `json:"uid" parquet:"uid,optional" ch:"uid"`
	UIDAlt                     *string `json:"uid_alt" parquet:"uid_alt,optional" ch:"uid_alt"`
	URLString                  *string `json:"url_string" parquet:"url_string,optional" ch:"url_string"`
	VendorName                 *string `json:"vendor_name" parquet:"vendor_name,optional" ch:"vendor_name"`
	Version                    *string `json:"version" parquet:"version,optional" ch:"version"`
}
