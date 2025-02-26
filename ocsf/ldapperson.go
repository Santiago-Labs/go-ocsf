package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

type LdapPerson struct {
	CostCenter      *string      `json:"cost_center,omitempty"`
	CreatedTime     *int         `json:"created_time,omitempty"`
	CreatedTimeDt   *string      `json:"created_time_dt,omitempty"`
	DeletedTime     *int         `json:"deleted_time,omitempty"`
	DeletedTimeDt   *string      `json:"deleted_time_dt,omitempty"`
	EmailAddrs      []string     `json:"email_addrs,omitempty"`
	EmployeeUID     *string      `json:"employee_uid,omitempty"`
	GivenName       *string      `json:"given_name,omitempty"`
	HireTime        *int         `json:"hire_time,omitempty"`
	HireTimeDt      *string      `json:"hire_time_dt,omitempty"`
	JobTitle        *string      `json:"job_title,omitempty"`
	Labels          []string     `json:"labels,omitempty"`
	LastLoginTime   *int         `json:"last_login_time,omitempty"`
	LastLoginTimeDt *string      `json:"last_login_time_dt,omitempty"`
	LDAPCn          *string      `json:"ldap_cn,omitempty"`
	LDAPDn          *string      `json:"ldap_dn,omitempty"`
	LeaveTime       *int         `json:"leave_time,omitempty"`
	LeaveTimeDt     *string      `json:"leave_time_dt,omitempty"`
	Location        *GeoLocation `json:"location,omitempty"`
	ModifiedTime    *int         `json:"modified_time,omitempty"`
	ModifiedTimeDt  *string      `json:"modified_time_dt,omitempty"`
	OfficeLocation  *string      `json:"office_location,omitempty"`
	Surname         *string      `json:"surname,omitempty"`
}

// LdapPersonFields defines the Arrow fields for LdapPerson.
var LdapPersonFields = []arrow.Field{
	{Name: "cost_center", Type: arrow.BinaryTypes.String},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "created_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "deleted_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "deleted_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "email_addrs", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "employee_uid", Type: arrow.BinaryTypes.String},
	{Name: "given_name", Type: arrow.BinaryTypes.String},
	{Name: "hire_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "hire_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "job_title", Type: arrow.BinaryTypes.String},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "last_login_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "last_login_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "ldap_cn", Type: arrow.BinaryTypes.String},
	{Name: "ldap_dn", Type: arrow.BinaryTypes.String},
	{Name: "leave_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "leave_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "location", Type: arrow.StructOf(GeoLocationFields...)},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int32},
	{Name: "modified_time_dt", Type: arrow.BinaryTypes.String},
	{Name: "office_location", Type: arrow.BinaryTypes.String},
	{Name: "surname", Type: arrow.BinaryTypes.String},
}

// LdapPersonSchema is the Arrow schema for LdapPerson.
var LdapPersonSchema = arrow.NewSchema(LdapPersonFields, nil)

// WriteToParquet writes the LdapPerson fields to the provided Arrow StructBuilder.
func (lp *LdapPerson) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: cost_center.
	costCenterB := sb.FieldBuilder(0).(*array.StringBuilder)
	if lp.CostCenter != nil {
		costCenterB.Append(*lp.CostCenter)
	} else {
		costCenterB.AppendNull()
	}

	// Field 1: created_time.
	createdTimeB := sb.FieldBuilder(1).(*array.Int32Builder)
	if lp.CreatedTime != nil {
		createdTimeB.Append(int32(*lp.CreatedTime))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 2: created_time_dt.
	createdTimeDtB := sb.FieldBuilder(2).(*array.StringBuilder)
	if lp.CreatedTimeDt != nil {
		createdTimeDtB.Append(*lp.CreatedTimeDt)
	} else {
		createdTimeDtB.AppendNull()
	}

	// Field 3: deleted_time.
	deletedTimeB := sb.FieldBuilder(3).(*array.Int32Builder)
	if lp.DeletedTime != nil {
		deletedTimeB.Append(int32(*lp.DeletedTime))
	} else {
		deletedTimeB.AppendNull()
	}

	// Field 4: deleted_time_dt.
	deletedTimeDtB := sb.FieldBuilder(4).(*array.StringBuilder)
	if lp.DeletedTimeDt != nil {
		deletedTimeDtB.Append(*lp.DeletedTimeDt)
	} else {
		deletedTimeDtB.AppendNull()
	}

	// Field 5: email_addrs (list of strings).
	emailAddrsB := sb.FieldBuilder(5).(*array.ListBuilder)
	if len(lp.EmailAddrs) > 0 {
		emailAddrsB.Append(true)
		emailValB := emailAddrsB.ValueBuilder().(*array.StringBuilder)
		for _, email := range lp.EmailAddrs {
			emailValB.Append(email)
		}
	} else {
		emailAddrsB.AppendNull()
	}

	// Field 6: employee_uid.
	employeeUIDB := sb.FieldBuilder(6).(*array.StringBuilder)
	if lp.EmployeeUID != nil {
		employeeUIDB.Append(*lp.EmployeeUID)
	} else {
		employeeUIDB.AppendNull()
	}

	// Field 7: given_name.
	givenNameB := sb.FieldBuilder(7).(*array.StringBuilder)
	if lp.GivenName != nil {
		givenNameB.Append(*lp.GivenName)
	} else {
		givenNameB.AppendNull()
	}

	// Field 8: hire_time.
	hireTimeB := sb.FieldBuilder(8).(*array.Int32Builder)
	if lp.HireTime != nil {
		hireTimeB.Append(int32(*lp.HireTime))
	} else {
		hireTimeB.AppendNull()
	}

	// Field 9: hire_time_dt.
	hireTimeDtB := sb.FieldBuilder(9).(*array.StringBuilder)
	if lp.HireTimeDt != nil {
		hireTimeDtB.Append(*lp.HireTimeDt)
	} else {
		hireTimeDtB.AppendNull()
	}

	// Field 10: job_title.
	jobTitleB := sb.FieldBuilder(10).(*array.StringBuilder)
	if lp.JobTitle != nil {
		jobTitleB.Append(*lp.JobTitle)
	} else {
		jobTitleB.AppendNull()
	}

	// Field 11: labels (list of strings).
	labelsB := sb.FieldBuilder(11).(*array.ListBuilder)
	if len(lp.Labels) > 0 {
		labelsB.Append(true)
		labelsValB := labelsB.ValueBuilder().(*array.StringBuilder)
		for _, label := range lp.Labels {
			labelsValB.Append(label)
		}
	} else {
		labelsB.AppendNull()
	}

	// Field 12: last_login_time.
	lastLoginTimeB := sb.FieldBuilder(12).(*array.Int32Builder)
	if lp.LastLoginTime != nil {
		lastLoginTimeB.Append(int32(*lp.LastLoginTime))
	} else {
		lastLoginTimeB.AppendNull()
	}

	// Field 13: last_login_time_dt.
	lastLoginTimeDtB := sb.FieldBuilder(13).(*array.StringBuilder)
	if lp.LastLoginTimeDt != nil {
		lastLoginTimeDtB.Append(*lp.LastLoginTimeDt)
	} else {
		lastLoginTimeDtB.AppendNull()
	}

	// Field 14: ldap_cn.
	ldapCnB := sb.FieldBuilder(14).(*array.StringBuilder)
	if lp.LDAPCn != nil {
		ldapCnB.Append(*lp.LDAPCn)
	} else {
		ldapCnB.AppendNull()
	}

	// Field 15: ldap_dn.
	ldapDnB := sb.FieldBuilder(15).(*array.StringBuilder)
	if lp.LDAPDn != nil {
		ldapDnB.Append(*lp.LDAPDn)
	} else {
		ldapDnB.AppendNull()
	}

	// Field 16: leave_time.
	leaveTimeB := sb.FieldBuilder(16).(*array.Int32Builder)
	if lp.LeaveTime != nil {
		leaveTimeB.Append(int32(*lp.LeaveTime))
	} else {
		leaveTimeB.AppendNull()
	}

	// Field 17: leave_time_dt.
	leaveTimeDtB := sb.FieldBuilder(17).(*array.StringBuilder)
	if lp.LeaveTimeDt != nil {
		leaveTimeDtB.Append(*lp.LeaveTimeDt)
	} else {
		leaveTimeDtB.AppendNull()
	}

	// Field 18: location (nested struct).
	locB := sb.FieldBuilder(18).(*array.StructBuilder)
	if lp.Location != nil {
		locB.Append(true)
		lp.Location.WriteToParquet(locB)
	} else {
		locB.AppendNull()
	}
	// Field 20: modified_time.
	modifiedTimeB := sb.FieldBuilder(20).(*array.Int32Builder)
	if lp.ModifiedTime != nil {
		modifiedTimeB.Append(int32(*lp.ModifiedTime))
	} else {
		modifiedTimeB.AppendNull()
	}

	// Field 21: modified_time_dt.
	modifiedTimeDtB := sb.FieldBuilder(21).(*array.StringBuilder)
	if lp.ModifiedTimeDt != nil {
		modifiedTimeDtB.Append(*lp.ModifiedTimeDt)
	} else {
		modifiedTimeDtB.AppendNull()
	}

	// Field 22: office_location.
	officeLocB := sb.FieldBuilder(22).(*array.StringBuilder)
	if lp.OfficeLocation != nil {
		officeLocB.Append(*lp.OfficeLocation)
	} else {
		officeLocB.AppendNull()
	}

	// Field 23: surname.
	surnameB := sb.FieldBuilder(23).(*array.StringBuilder)
	if lp.Surname != nil {
		surnameB.Append(*lp.Surname)
	} else {
		surnameB.AppendNull()
	}
}
