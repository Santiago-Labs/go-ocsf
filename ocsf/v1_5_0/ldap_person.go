// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type LDAPPerson struct {

	// Cost Center: The cost center associated with the user.
	CostCenter *string `json:"cost_center,omitempty" parquet:"cost_center,optional"`

	// Created Time: The timestamp when the user was created.
	CreatedTime *int64 `json:"created_time,omitempty" parquet:"created_time,optional"`

	// Deleted Time: The timestamp when the user was deleted. In Active Directory (AD), when a user is deleted they are moved to a temporary container and then removed after 30 days. So, this field can be populated even after a user is deleted for the next 30 days.
	DeletedTime *int64 `json:"deleted_time,omitempty" parquet:"deleted_time,optional"`

	// Display Name: The display name of the LDAP person. According to RFC 2798, this is the preferred name of a person to be used when displaying entries.
	DisplayName *string `json:"display_name,omitempty" parquet:"display_name,optional"`

	// Email Addresses: A list of additional email addresses for the user.
	EmailAddrs []string `json:"email_addrs,omitempty" parquet:"email_addrs,optional,list"`

	// Employee ID: The employee identifier assigned to the user by the organization.
	EmployeeUid *string `json:"employee_uid,omitempty" parquet:"employee_uid,optional"`

	// Given Name: The given or first name of the user.
	GivenName *string `json:"given_name,omitempty" parquet:"given_name,optional"`

	// Hire Time: The timestamp when the user was or will be hired by the organization.
	HireTime *int64 `json:"hire_time,omitempty" parquet:"hire_time,optional"`

	// Job Title: The user's job title.
	JobTitle *string `json:"job_title,omitempty" parquet:"job_title,optional"`

	// Labels: The labels associated with the user. For example in AD this could be the <code>userType</code>, <code>employeeType</code>. For example: <code>Member, Employee</code>.
	Labels []string `json:"labels,omitempty" parquet:"labels,optional,list"`

	// Last Login: The last time when the user logged in.
	LastLoginTime *int64 `json:"last_login_time,omitempty" parquet:"last_login_time,optional"`

	// LDAP Common Name: The LDAP and X.500 <code>commonName</code> attribute, typically the full name of the person. For example, <code>John Doe</code>.
	LdapCn *string `json:"ldap_cn,omitempty" parquet:"ldap_cn,optional"`

	// LDAP Distinguished Name: The X.500 Distinguished Name (DN) is a structured string that uniquely identifies an entry, such as a user, in an X.500 directory service For example, <code>cn=John Doe,ou=People,dc=example,dc=com</code>.
	LdapDn *string `json:"ldap_dn,omitempty" parquet:"ldap_dn,optional"`

	// Leave Time: The timestamp when the user left or will be leaving the organization.
	LeaveTime *int64 `json:"leave_time,omitempty" parquet:"leave_time,optional"`

	// Geo Location: The geographical location associated with a user. This is typically the user's usual work location.
	Location *GeoLocation `json:"location,omitempty" parquet:"location,optional"`

	// Manager: The user's manager. This helps in understanding an org hierarchy. This should only ever be populated once in an event. I.e. there should not be a manager's manager in an event.
	Manager *UserRef `json:"manager,omitempty" parquet:"manager,optional"`

	// Modified Time: The timestamp when the user entry was last modified.
	ModifiedTime *int64 `json:"modified_time,omitempty" parquet:"modified_time,optional"`

	// Office Location: The primary office location associated with the user. This could be any string and isn't a specific address. For example, <code>South East Virtual</code>.
	OfficeLocation *string `json:"office_location,omitempty" parquet:"office_location,optional"`

	// Telephone Number: The telephone number of the user. Corresponds to the LDAP <code>Telephone-Number</code> CN.
	PhoneNumber *string `json:"phone_number,omitempty" parquet:"phone_number,optional"`

	// Surname: The last or family name for the user.
	Surname *string `json:"surname,omitempty" parquet:"surname,optional"`

	// Tags: The list of tags; <code>{key:value}</code> pairs associated to the user.
	Tags []*KeyValueobject `json:"tags,omitempty" parquet:"tags,optional,list"`
}

var LDAPPersonFields = []arrow.Field{
	{Name: "cost_center", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "deleted_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "display_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "email_addrs", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "employee_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "given_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hire_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "job_title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "last_login_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "ldap_cn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ldap_dn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "leave_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "location", Type: GeoLocationStruct, Nullable: true},
	{Name: "manager", Type: UserRefStruct, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "office_location", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "phone_number", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "surname", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "tags", Type: arrow.ListOf(KeyValueobjectStruct), Nullable: true},
}

var LDAPPersonStruct = arrow.StructOf(LDAPPersonFields...)

var LDAPPersonSchema = arrow.NewSchema(LDAPPersonFields, nil)
var LDAPPersonRefFields = []arrow.Field{
	{Name: "cost_center", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "deleted_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "display_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "email_addrs", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "employee_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "given_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hire_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "job_title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "last_login_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "ldap_cn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ldap_dn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "leave_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},

	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "office_location", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "phone_number", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "surname", Type: arrow.BinaryTypes.String, Nullable: true},
}
var LDAPPersonRefStruct = arrow.StructOf(LDAPPersonRefFields...)

type LDAPPersonRef struct {

	// Cost Center: The cost center associated with the user.
	CostCenter *string `json:"cost_center,omitempty" parquet:"cost_center,optional"`

	// Created Time: The timestamp when the user was created.
	CreatedTime *int64 `json:"created_time,omitempty" parquet:"created_time,optional"`

	// Deleted Time: The timestamp when the user was deleted. In Active Directory (AD), when a user is deleted they are moved to a temporary container and then removed after 30 days. So, this field can be populated even after a user is deleted for the next 30 days.
	DeletedTime *int64 `json:"deleted_time,omitempty" parquet:"deleted_time,optional"`

	// Display Name: The display name of the LDAP person. According to RFC 2798, this is the preferred name of a person to be used when displaying entries.
	DisplayName *string `json:"display_name,omitempty" parquet:"display_name,optional"`

	// Email Addresses: A list of additional email addresses for the user.
	EmailAddrs []string `json:"email_addrs,omitempty" parquet:"email_addrs,optional,list"`

	// Employee ID: The employee identifier assigned to the user by the organization.
	EmployeeUid *string `json:"employee_uid,omitempty" parquet:"employee_uid,optional"`

	// Given Name: The given or first name of the user.
	GivenName *string `json:"given_name,omitempty" parquet:"given_name,optional"`

	// Hire Time: The timestamp when the user was or will be hired by the organization.
	HireTime *int64 `json:"hire_time,omitempty" parquet:"hire_time,optional"`

	// Job Title: The user's job title.
	JobTitle *string `json:"job_title,omitempty" parquet:"job_title,optional"`

	// Labels: The labels associated with the user. For example in AD this could be the <code>userType</code>, <code>employeeType</code>. For example: <code>Member, Employee</code>.
	Labels []string `json:"labels,omitempty" parquet:"labels,optional,list"`

	// Last Login: The last time when the user logged in.
	LastLoginTime *int64 `json:"last_login_time,omitempty" parquet:"last_login_time,optional"`

	// LDAP Common Name: The LDAP and X.500 <code>commonName</code> attribute, typically the full name of the person. For example, <code>John Doe</code>.
	LdapCn *string `json:"ldap_cn,omitempty" parquet:"ldap_cn,optional"`

	// LDAP Distinguished Name: The X.500 Distinguished Name (DN) is a structured string that uniquely identifies an entry, such as a user, in an X.500 directory service For example, <code>cn=John Doe,ou=People,dc=example,dc=com</code>.
	LdapDn *string `json:"ldap_dn,omitempty" parquet:"ldap_dn,optional"`

	// Leave Time: The timestamp when the user left or will be leaving the organization.
	LeaveTime *int64 `json:"leave_time,omitempty" parquet:"leave_time,optional"`

	// Geo Location: The geographical location associated with a user. This is typically the user's usual work location.

	// Manager: The user's manager. This helps in understanding an org hierarchy. This should only ever be populated once in an event. I.e. there should not be a manager's manager in an event.

	// Modified Time: The timestamp when the user entry was last modified.
	ModifiedTime *int64 `json:"modified_time,omitempty" parquet:"modified_time,optional"`

	// Office Location: The primary office location associated with the user. This could be any string and isn't a specific address. For example, <code>South East Virtual</code>.
	OfficeLocation *string `json:"office_location,omitempty" parquet:"office_location,optional"`

	// Telephone Number: The telephone number of the user. Corresponds to the LDAP <code>Telephone-Number</code> CN.
	PhoneNumber *string `json:"phone_number,omitempty" parquet:"phone_number,optional"`

	// Surname: The last or family name for the user.
	Surname *string `json:"surname,omitempty" parquet:"surname,optional"`

	// Tags: The list of tags; <code>{key:value}</code> pairs associated to the user.

}
