package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// UserFields defines the Arrow fields for User.
var UserFields = []arrow.Field{
	{Name: "account", Type: AccountStruct, Nullable: true},
	{Name: "credential_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "email_addr", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "full_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "groups", Type: arrow.ListOf(GroupStruct), Nullable: true},
	{Name: "ldap_person", Type: LdapPersonStruct, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "org", Type: OrganizationStruct, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String, Nullable: true},
}

var UserStruct = arrow.StructOf(UserFields...)
var UserClassname = "user"

type User struct {
	Account       *Account      `json:"account" parquet:"account,optional" ch:"account"`
	CredentialUID *string       `json:"credential_uid" parquet:"credential_uid,optional" ch:"credential_uid"`
	Domain        *string       `json:"domain" parquet:"domain,optional" ch:"domain"`
	EmailAddr     *string       `json:"email_addr" parquet:"email_addr,optional" ch:"email_addr"`
	FullName      *string       `json:"full_name" parquet:"full_name,optional" ch:"full_name"`
	Groups        []*Group      `json:"groups" parquet:"groups,list,optional" ch:"groups"`
	LDAPPerson    *LdapPerson   `json:"ldap_person" parquet:"ldap_person,optional" ch:"ldap_person"`
	Name          *string       `json:"name" parquet:"name,optional" ch:"name"`
	Org           *Organization `json:"org" parquet:"org,optional" ch:"org"`
	Type          *string       `json:"type" parquet:"type,optional" ch:"type"`
	// TypeID enum: [3,99,0,1,2]
	TypeID *int64  `json:"type_id" parquet:"type_id,optional" ch:"type_id"`
	UID    *string `json:"uid" parquet:"uid,optional" ch:"uid"`
	UIDAlt *string `json:"uid_alt" parquet:"uid_alt,optional" ch:"uid_alt"`
}
