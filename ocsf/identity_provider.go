package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// IdentityProviderFields defines the Arrow fields for IdentityProvider.
var IdentityProviderFields = []arrow.Field{
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "fingerprint", Type: FingerprintStruct, Nullable: true},
	{Name: "has_mfa", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "issuer", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "auth_factors", Type: arrow.ListOf(AuthFactorStruct), Nullable: true},
	{Name: "scim", Type: SCIMStruct, Nullable: true},
	{Name: "sso", Type: SSOStruct, Nullable: true},
	{Name: "state", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "state_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "tenant_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "url_string", Type: arrow.BinaryTypes.String, Nullable: true},
}

var IdentityProviderStruct = arrow.StructOf(IdentityProviderFields...)
var IdentityProviderClassname = "idp"

type IdentityProvider struct {
	AuthFactors  []*AuthFactor `json:"auth_factors,omitempty" parquet:"auth_factors,list,optional"`
	Domain       *string       `json:"domain,omitempty" parquet:"domain,optional"`
	Fingerprint  *Fingerprint  `json:"fingerprint,omitempty" parquet:"fingerprint,optional"`
	HasMFA       *bool         `json:"has_mfa,omitempty" parquet:"has_mfa,optional"`
	Issuer       *string       `json:"issuer,omitempty" parquet:"issuer,optional"`
	Name         *string       `json:"name,omitempty" parquet:"name,optional"`
	ProtocolName *string       `json:"protocol_name,omitempty" parquet:"protocol_name,optional"`
	SCIM         *SCIM         `json:"scim,omitempty" parquet:"scim,optional"`
	SSO          *SSO          `json:"sso,omitempty" parquet:"sso,optional"`
	State        *string       `json:"state,omitempty" parquet:"state,optional"`
	StateID      *int          `json:"state_id,omitempty" parquet:"state_id,optional"`
	TenantUID    *string       `json:"tenant_uid,omitempty" parquet:"tenant_uid,optional"`
	UID          *string       `json:"uid,omitempty" parquet:"uid,optional"`
	URLString    *string       `json:"url_string,omitempty" parquet:"url_string,optional"`
}
