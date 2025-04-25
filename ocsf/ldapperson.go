package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type LdapPerson struct {
	CostCenter      *string      `json:"cost_center" parquet:"cost_center,optional" ch:"cost_center" ch:"cost_center"`
	CreatedTime     *int64       `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	CreatedTimeDt   *string      `json:"created_time_dt" parquet:"created_time_dt,optional" ch:"created_time_dt"`
	DeletedTime     *int64       `json:"deleted_time" parquet:"deleted_time,optional" ch:"deleted_time"`
	DeletedTimeDt   *string      `json:"deleted_time_dt" parquet:"deleted_time_dt,optional" ch:"deleted_time_dt"`
	EmailAddrs      []string     `json:"email_addrs" parquet:"email_addrs,list,optional" ch:"email_addrs"`
	EmployeeUID     *string      `json:"employee_uid" parquet:"employee_uid,optional" ch:"employee_uid"`
	GivenName       *string      `json:"given_name" parquet:"given_name,optional" ch:"given_name"`
	HireTime        *int64       `json:"hire_time" parquet:"hire_time,optional" ch:"hire_time"`
	HireTimeDt      *string      `json:"hire_time_dt" parquet:"hire_time_dt,optional" ch:"hire_time_dt"`
	JobTitle        *string      `json:"job_title" parquet:"job_title,optional" ch:"job_title"`
	Labels          []string     `json:"labels" parquet:"labels,list,optional" ch:"labels"`
	LastLoginTime   *int64       `json:"last_login_time" parquet:"last_login_time,optional" ch:"last_login_time"`
	LastLoginTimeDt *string      `json:"last_login_time_dt" parquet:"last_login_time_dt,optional" ch:"last_login_time_dt"`
	LDAPCn          *string      `json:"ldap_cn" parquet:"ldap_cn,optional" ch:"ldap_cn"`
	LDAPDn          *string      `json:"ldap_dn" parquet:"ldap_dn,optional" ch:"ldap_dn"`
	LeaveTime       *int64       `json:"leave_time" parquet:"leave_time,optional" ch:"leave_time"`
	LeaveTimeDt     *string      `json:"leave_time_dt" parquet:"leave_time_dt,optional" ch:"leave_time_dt"`
	Location        *GeoLocation `json:"location" parquet:"location,optional" ch:"location"`
	ModifiedTime    *int64       `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	ModifiedTimeDt  *string      `json:"modified_time_dt" parquet:"modified_time_dt,optional" ch:"modified_time_dt"`
	OfficeLocation  *string      `json:"office_location" parquet:"office_location,optional" ch:"office_location"`
	Surname         *string      `json:"surname" parquet:"surname,optional" ch:"surname"`
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
