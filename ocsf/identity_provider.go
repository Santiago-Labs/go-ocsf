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
	AuthFactors  []*AuthFactor `json:"auth_factors" parquet:"auth_factors,list,optional" ch:"auth_factors"`
	Domain       *string       `json:"domain" parquet:"domain,optional" ch:"domain"`
	Fingerprint  *Fingerprint  `json:"fingerprint" parquet:"fingerprint,optional" ch:"fingerprint"`
	HasMFA       *bool         `json:"has_mfa" parquet:"has_mfa,optional" ch:"has_mfa"`
	Issuer       *string       `json:"issuer" parquet:"issuer,optional" ch:"issuer"`
	Name         *string       `json:"name" parquet:"name,optional" ch:"name"`
	ProtocolName *string       `json:"protocol_name" parquet:"protocol_name,optional" ch:"protocol_name"`
	SCIM         *SCIM         `json:"scim" parquet:"scim,optional" ch:"scim"`
	SSO          *SSO          `json:"sso" parquet:"sso,optional" ch:"sso"`
	State        *string       `json:"state" parquet:"state,optional" ch:"state"`
	StateID      *int64        `json:"state_id" parquet:"state_id,optional" ch:"state_id"`
	TenantUID    *string       `json:"tenant_uid" parquet:"tenant_uid,optional" ch:"tenant_uid"`
	UID          *string       `json:"uid" parquet:"uid,optional" ch:"uid"`
	URLString    *string       `json:"url_string" parquet:"url_string,optional" ch:"url_string"`
}
