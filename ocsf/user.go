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
	Account       *Account      `json:"account,omitempty" parquet:"account,optional"`
	CredentialUID *string       `json:"credential_uid,omitempty" parquet:"credential_uid,optional"`
	Domain        *string       `json:"domain,omitempty" parquet:"domain,optional"`
	EmailAddr     *string       `json:"email_addr,omitempty" parquet:"email_addr,optional"`
	FullName      *string       `json:"full_name,omitempty" parquet:"full_name,optional"`
	Groups        []*Group      `json:"groups,omitempty" parquet:"groups,list,optional"`
	LDAPPerson    *LdapPerson   `json:"ldap_person,omitempty" parquet:"ldap_person,optional"`
	Name          *string       `json:"name,omitempty" parquet:"name,optional"`
	Org           *Organization `json:"org,omitempty" parquet:"org,optional"`
	Type          *string       `json:"type,omitempty" parquet:"type,optional"`
	// TypeID enum: [3,99,0,1,2]
	TypeID *int    `json:"type_id,omitempty" parquet:"type_id,optional"`
	UID    *string `json:"uid,omitempty" parquet:"uid,optional"`
	UIDAlt *string `json:"uid_alt,omitempty" parquet:"uid_alt,optional"`
}
