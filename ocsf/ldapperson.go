package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

type LdapPerson struct {
	CostCenter      *string      `json:"cost_center,omitempty" parquet:"cost_center"`
	CreatedTime     *int         `json:"created_time,omitempty" parquet:"created_time"`
	CreatedTimeDt   *string      `json:"created_time_dt,omitempty" parquet:"created_time_dt"`
	DeletedTime     *int         `json:"deleted_time,omitempty" parquet:"deleted_time"`
	DeletedTimeDt   *string      `json:"deleted_time_dt,omitempty" parquet:"deleted_time_dt"`
	EmailAddrs      []string     `json:"email_addrs,omitempty" parquet:"email_addrs"`
	EmployeeUID     *string      `json:"employee_uid,omitempty" parquet:"employee_uid"`
	GivenName       *string      `json:"given_name,omitempty" parquet:"given_name"`
	HireTime        *int         `json:"hire_time,omitempty" parquet:"hire_time"`
	HireTimeDt      *string      `json:"hire_time_dt,omitempty" parquet:"hire_time_dt"`
	JobTitle        *string      `json:"job_title,omitempty" parquet:"job_title"`
	Labels          []string     `json:"labels,omitempty" parquet:"labels"`
	LastLoginTime   *int         `json:"last_login_time,omitempty" parquet:"last_login_time"`
	LastLoginTimeDt *string      `json:"last_login_time_dt,omitempty" parquet:"last_login_time_dt"`
	LDAPCn          *string      `json:"ldap_cn,omitempty" parquet:"ldap_cn"`
	LDAPDn          *string      `json:"ldap_dn,omitempty" parquet:"ldap_dn"`
	LeaveTime       *int         `json:"leave_time,omitempty" parquet:"leave_time"`
	LeaveTimeDt     *string      `json:"leave_time_dt,omitempty" parquet:"leave_time_dt"`
	Location        *GeoLocation `json:"location,omitempty" parquet:"location"`
	ModifiedTime    *int         `json:"modified_time,omitempty" parquet:"modified_time"`
	ModifiedTimeDt  *string      `json:"modified_time_dt,omitempty" parquet:"modified_time_dt"`
	OfficeLocation  *string      `json:"office_location,omitempty" parquet:"office_location"`
	Surname         *string      `json:"surname,omitempty" parquet:"surname"`
}

// LdapPersonFields defines the Arrow fields for LdapPerson.
var LdapPersonFields = []arrow.Field{
	{Name: "cost_center", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "deleted_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "deleted_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "email_addrs", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "employee_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "given_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hire_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "hire_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "job_title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "last_login_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "last_login_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ldap_cn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ldap_dn", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "leave_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "leave_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "location", Type: GeoLocationStruct, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "modified_time_dt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "office_location", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "surname", Type: arrow.BinaryTypes.String, Nullable: true},
}

var LdapPersonStruct = arrow.StructOf(LdapPersonFields...)
var LdapPersonClassname = "ldap_person"
