package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// SessionFields defines the Arrow fields for Session.
var SessionFields = []arrow.Field{
	{Name: "count", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "credential_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "expiration_reason", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "expiration_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "is_mfa", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_remote", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "is_vpn", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "issuer", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "terminal", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uuid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var SessionStruct = arrow.StructOf(SessionFields...)
var SessionClassname = "session"

type Session struct {
	Count            *int64  `json:"count" parquet:"count,optional" ch:"count" ch:"count"`
	CreatedTime      *int64  `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	CredentialUID    *string `json:"credential_uid" parquet:"credential_uid,optional" ch:"credential_uid"`
	ExpirationReason *string `json:"expiration_reason" parquet:"expiration_reason,optional" ch:"expiration_reason"`
	ExpirationTime   *int64  `json:"expiration_time" parquet:"expiration_time,optional" ch:"expiration_time"`
	IsMFA            *bool   `json:"is_mfa" parquet:"is_mfa,optional" ch:"is_mfa"`
	IsRemote         *bool   `json:"is_remote" parquet:"is_remote,optional" ch:"is_remote"`
	IsVPN            *bool   `json:"is_vpn" parquet:"is_vpn,optional" ch:"is_vpn"`
	Issuer           *string `json:"issuer" parquet:"issuer,optional" ch:"issuer"`
	Terminal         *string `json:"terminal" parquet:"terminal,optional" ch:"terminal"`
	UID              *string `json:"uid" parquet:"uid,optional" ch:"uid"`
	UIDAlt           *string `json:"uid_alt" parquet:"uid_alt,optional" ch:"uid_alt"`
	UUID             *string `json:"uuid" parquet:"uuid,optional" ch:"uuid"`
}
