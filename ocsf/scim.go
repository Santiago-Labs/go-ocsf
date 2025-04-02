package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// SCIMFields defines the Arrow fields for SCIM.
var SCIMFields = []arrow.Field{
	{Name: "auth_protocol", Type: arrow.BinaryTypes.String},
	{Name: "auth_protocol_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "error_message", Type: arrow.BinaryTypes.String},
	{Name: "is_group_provisioning_enabled", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "is_user_provisioning_enabled", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "last_run_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String},
	{Name: "rate_limit", Type: arrow.PrimitiveTypes.Int32},
	{Name: "scim_group_schema", Type: arrow.BinaryTypes.String},
	{Name: "scim_user_schema", Type: arrow.BinaryTypes.String},
	{Name: "state", Type: arrow.BinaryTypes.String},
	{Name: "state_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
	{Name: "url_string", Type: arrow.BinaryTypes.String},
	{Name: "vendor_name", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var SCIMStruct = arrow.StructOf(SCIMFields...)

type SCIM struct {
	AuthProtocol               *string    `json:"auth_protocol,omitempty" parquet:"auth_protocol"`
	AuthProtocolID             *int32     `json:"auth_protocol_id,omitempty" parquet:"auth_protocol_id"`
	CreatedTime                *time.Time `json:"created_time,omitempty" parquet:"created_time"`
	ErrorMessage               *string    `json:"error_message,omitempty" parquet:"error_message"`
	IsGroupProvisioningEnabled *bool      `json:"is_group_provisioning_enabled,omitempty" parquet:"is_group_provisioning_enabled"`
	IsUserProvisioningEnabled  *bool      `json:"is_user_provisioning_enabled,omitempty" parquet:"is_user_provisioning_enabled"`
	LastRunTime                *time.Time `json:"last_run_time,omitempty" parquet:"last_run_time"`
	ModifiedTime               *time.Time `json:"modified_time,omitempty" parquet:"modified_time"`
	Name                       *string    `json:"name,omitempty" parquet:"name"`
	ProtocolName               *string    `json:"protocol_name,omitempty" parquet:"protocol_name"`
	RateLimit                  *int32     `json:"rate_limit,omitempty" parquet:"rate_limit"`
	SCIMGroupSchema            *string    `json:"scim_group_schema,omitempty" parquet:"scim_group_schema"`
	SCIMUserSchema             *string    `json:"scim_user_schema,omitempty" parquet:"scim_user_schema"`
	State                      *string    `json:"state,omitempty" parquet:"state"`
	StateID                    *int32     `json:"state_id,omitempty" parquet:"state_id"`
	UID                        *string    `json:"uid,omitempty" parquet:"uid"`
	UIDAlt                     *string    `json:"uid_alt,omitempty" parquet:"uid_alt"`
	URLString                  *string    `json:"url_string,omitempty" parquet:"url_string"`
	VendorName                 *string    `json:"vendor_name,omitempty" parquet:"vendor_name"`
	Version                    *string    `json:"version,omitempty" parquet:"version"`
}
