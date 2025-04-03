package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// SessionFields defines the Arrow fields for Session.
var SessionFields = []arrow.Field{
	{Name: "count", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "credential_uid", Type: arrow.BinaryTypes.String},
	{Name: "expiration_reason", Type: arrow.BinaryTypes.String},
	{Name: "expiration_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "is_mfa", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "is_remote", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "is_vpn", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "issuer", Type: arrow.BinaryTypes.String},
	{Name: "terminal", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
	{Name: "uuid", Type: arrow.BinaryTypes.String},
}

var SessionStruct = arrow.StructOf(SessionFields...)
var SessionClassname = "session"

type Session struct {
	Count            *int       `json:"count,omitempty" parquet:"count"`
	CreatedTime      *time.Time `json:"created_time,omitempty" parquet:"created_time"`
	CredentialUID    *string    `json:"credential_uid,omitempty" parquet:"credential_uid"`
	ExpirationReason *string    `json:"expiration_reason,omitempty" parquet:"expiration_reason"`
	ExpirationTime   *time.Time `json:"expiration_time,omitempty" parquet:"expiration_time"`
	IsMFA            *bool      `json:"is_mfa,omitempty" parquet:"is_mfa"`
	IsRemote         *bool      `json:"is_remote,omitempty" parquet:"is_remote"`
	IsVPN            *bool      `json:"is_vpn,omitempty" parquet:"is_vpn"`
	Issuer           *string    `json:"issuer,omitempty" parquet:"issuer"`
	Terminal         *string    `json:"terminal,omitempty" parquet:"terminal"`
	UID              *string    `json:"uid,omitempty" parquet:"uid"`
	UIDAlt           *string    `json:"uid_alt,omitempty" parquet:"uid_alt"`
	UUID             *string    `json:"uuid,omitempty" parquet:"uuid"`
}
