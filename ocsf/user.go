package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// UserFields defines the Arrow fields for User.
var UserFields = []arrow.Field{
	{Name: "account", Type: arrow.StructOf(AccountFields...)},
	{Name: "credential_uid", Type: arrow.BinaryTypes.String},
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "email_addr", Type: arrow.BinaryTypes.String},
	{Name: "full_name", Type: arrow.BinaryTypes.String},
	{Name: "groups", Type: arrow.ListOf(arrow.StructOf(GroupFields...))},
	{Name: "ldap_person", Type: arrow.StructOf(LdapPersonFields...)},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "org", Type: arrow.StructOf(OrganizationFields...)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
}

// UserSchema is the Arrow schema for User.
var UserSchema = arrow.NewSchema(UserFields, nil)

// User represents a user.
type User struct {
	Account       *Account      `json:"account,omitempty"`
	CredentialUID *string       `json:"credential_uid,omitempty"`
	Domain        *string       `json:"domain,omitempty"`
	EmailAddr     *string       `json:"email_addr,omitempty"`
	FullName      *string       `json:"full_name,omitempty"`
	Groups        []Group       `json:"groups,omitempty"`
	LDAPPerson    *LdapPerson   `json:"ldap_person,omitempty"`
	Name          *string       `json:"name,omitempty"`
	Org           *Organization `json:"org,omitempty"`
	Type          *string       `json:"type,omitempty"`
	// TypeID enum: [3,99,0,1,2]
	TypeID *int    `json:"type_id,omitempty"`
	UID    *string `json:"uid,omitempty"`
	UIDAlt *string `json:"uid_alt,omitempty"`
}

// WriteToParquet writes the User fields to the provided Arrow StructBuilder.
func (u *User) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: account (nested struct).
	accountB := sb.FieldBuilder(0).(*array.StructBuilder)
	if u.Account != nil {
		accountB.Append(true)
		u.Account.WriteToParquet(accountB)
	} else {
		accountB.AppendNull()
	}

	// Field 1: credential_uid.
	credB := sb.FieldBuilder(1).(*array.StringBuilder)
	if u.CredentialUID != nil {
		credB.Append(*u.CredentialUID)
	} else {
		credB.AppendNull()
	}

	// Field 2: domain.
	domainB := sb.FieldBuilder(2).(*array.StringBuilder)
	if u.Domain != nil {
		domainB.Append(*u.Domain)
	} else {
		domainB.AppendNull()
	}

	// Field 3: email_addr.
	emailB := sb.FieldBuilder(3).(*array.StringBuilder)
	if u.EmailAddr != nil {
		emailB.Append(*u.EmailAddr)
	} else {
		emailB.AppendNull()
	}

	// Field 4: full_name.
	fullNameB := sb.FieldBuilder(4).(*array.StringBuilder)
	if u.FullName != nil {
		fullNameB.Append(*u.FullName)
	} else {
		fullNameB.AppendNull()
	}

	// Field 5: groups (list of nested Group structs).
	groupsB := sb.FieldBuilder(5).(*array.ListBuilder)
	if len(u.Groups) > 0 {
		groupsB.Append(true)
		groupsValB := groupsB.ValueBuilder().(*array.StructBuilder)
		for _, grp := range u.Groups {
			groupsValB.Append(true)
			grp.WriteToParquet(groupsValB)
		}
	} else {
		groupsB.AppendNull()
	}

	// Field 6: ldap_person (nested struct).
	ldapB := sb.FieldBuilder(6).(*array.StructBuilder)
	if u.LDAPPerson != nil {
		ldapB.Append(true)
		u.LDAPPerson.WriteToParquet(ldapB)
	} else {
		ldapB.AppendNull()
	}

	// Field 7: name.
	nameB := sb.FieldBuilder(7).(*array.StringBuilder)
	if u.Name != nil {
		nameB.Append(*u.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 8: org (nested struct).
	orgB := sb.FieldBuilder(8).(*array.StructBuilder)
	if u.Org != nil {
		orgB.Append(true)
		u.Org.WriteToParquet(orgB)
	} else {
		orgB.AppendNull()
	}

	// Field 9: type.
	typeB := sb.FieldBuilder(9).(*array.StringBuilder)
	if u.Type != nil {
		typeB.Append(*u.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 10: type_id.
	typeIDB := sb.FieldBuilder(10).(*array.Int32Builder)
	if u.TypeID != nil {
		typeIDB.Append(int32(*u.TypeID))
	} else {
		typeIDB.AppendNull()
	}

	// Field 11: uid.
	uidB := sb.FieldBuilder(11).(*array.StringBuilder)
	if u.UID != nil {
		uidB.Append(*u.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 12: uid_alt.
	uidAltB := sb.FieldBuilder(12).(*array.StringBuilder)
	if u.UIDAlt != nil {
		uidAltB.Append(*u.UIDAlt)
	} else {
		uidAltB.AppendNull()
	}
}
