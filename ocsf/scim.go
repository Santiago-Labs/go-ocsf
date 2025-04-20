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
	AuthProtocol               *string `json:"auth_protocol,omitempty" parquet:"auth_protocol,optional"`
	AuthProtocolID             *int32  `json:"auth_protocol_id,omitempty" parquet:"auth_protocol_id,optional"`
	CreatedTime                *int64  `json:"created_time,omitempty" parquet:"created_time,optional"`
	ErrorMessage               *string `json:"error_message,omitempty" parquet:"error_message,optional"`
	IsGroupProvisioningEnabled *bool   `json:"is_group_provisioning_enabled,omitempty" parquet:"is_group_provisioning_enabled,optional"`
	IsUserProvisioningEnabled  *bool   `json:"is_user_provisioning_enabled,omitempty" parquet:"is_user_provisioning_enabled,optional"`
	LastRunTime                *int64  `json:"last_run_time,omitempty" parquet:"last_run_time,optional"`
	ModifiedTime               *int64  `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	Name                       *string `json:"name,omitempty" parquet:"name,optional"`
	ProtocolName               *string `json:"protocol_name,omitempty" parquet:"protocol_name,optional"`
	RateLimit                  *int32  `json:"rate_limit,omitempty" parquet:"rate_limit,optional"`
	SCIMGroupSchema            *string `json:"scim_group_schema,omitempty" parquet:"scim_group_schema,optional"`
	SCIMUserSchema             *string `json:"scim_user_schema,omitempty" parquet:"scim_user_schema,optional"`
	State                      *string `json:"state,omitempty" parquet:"state,optional"`
	StateID                    *int32  `json:"state_id,omitempty" parquet:"state_id,optional"`
	UID                        *string `json:"uid,omitempty" parquet:"uid,optional"`
	UIDAlt                     *string `json:"uid_alt,omitempty" parquet:"uid_alt,optional"`
	URLString                  *string `json:"url_string,omitempty" parquet:"url_string,optional"`
	VendorName                 *string `json:"vendor_name,omitempty" parquet:"vendor_name,optional"`
	Version                    *string `json:"version,omitempty" parquet:"version,optional"`
}
