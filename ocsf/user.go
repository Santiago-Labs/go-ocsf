package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// UserFields defines the Arrow fields for User.
var UserFields = []arrow.Field{
	{Name: "account", Type: AccountStruct},
	{Name: "credential_uid", Type: arrow.BinaryTypes.String},
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "email_addr", Type: arrow.BinaryTypes.String},
	{Name: "full_name", Type: arrow.BinaryTypes.String},
	{Name: "groups", Type: arrow.ListOf(GroupStruct)},
	{Name: "ldap_person", Type: LdapPersonStruct},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "org", Type: OrganizationStruct},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
}

var UserStruct = arrow.StructOf(UserFields...)
var UserClassname = "user"

type User struct {
	Account       *Account      `json:"account,omitempty" parquet:"account"`
	CredentialUID *string       `json:"credential_uid,omitempty" parquet:"credential_uid"`
	Domain        *string       `json:"domain,omitempty" parquet:"domain"`
	EmailAddr     *string       `json:"email_addr,omitempty" parquet:"email_addr"`
	FullName      *string       `json:"full_name,omitempty" parquet:"full_name"`
	Groups        []Group       `json:"groups,omitempty" parquet:"groups"`
	LDAPPerson    *LdapPerson   `json:"ldap_person,omitempty" parquet:"ldap_person"`
	Name          *string       `json:"name,omitempty" parquet:"name"`
	Org           *Organization `json:"org,omitempty" parquet:"org"`
	Type          *string       `json:"type,omitempty" parquet:"type"`
	// TypeID enum: [3,99,0,1,2]
	TypeID *int    `json:"type_id,omitempty" parquet:"type_id"`
	UID    *string `json:"uid,omitempty" parquet:"uid"`
	UIDAlt *string `json:"uid_alt,omitempty" parquet:"uid_alt"`
}
