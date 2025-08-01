// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_4_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type WHOIS struct {

	// Autonomous System: The autonomous system information associated with a domain.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty" parquet:"autonomous_system,optional"`

	// Registered At: When the domain was registered or WHOIS entry was created.
	CreatedTime int64 `json:"created_time,omitempty" parquet:"created_time,timestamp_millis,timestamp(millisecond),optional"`

	// DNSSEC Status: The normalized value of dnssec_status_id.
	DnssecStatus *string `json:"dnssec_status,omitempty" parquet:"dnssec_status,optional"`

	// DNSSEC Status ID: Describes the normalized status of DNS Security Extensions (DNSSEC) for a domain.
	DnssecStatusId *int32 `json:"dnssec_status_id,omitempty" parquet:"dnssec_status_id,optional"`

	// Domain: The domain name corresponding to the WHOIS record.
	Domain *string `json:"domain,omitempty" parquet:"domain,optional"`

	// Domain Contacts: An array of <code>Domain Contact</code> objects.
	DomainContacts []DomainContact `json:"domain_contacts,omitempty" parquet:"domain_contacts,list,optional"`

	// Registrar Abuse Email Address: The email address for the registrar's abuse contact
	EmailAddr *string `json:"email_addr,omitempty" parquet:"email_addr,optional"`

	// Last Updated At: When the WHOIS record was last updated or seen at.
	LastSeenTime int64 `json:"last_seen_time,omitempty" parquet:"last_seen_time,timestamp_millis,timestamp(millisecond),optional"`

	// Name Servers: A collection of name servers related to a domain registration or other record.
	NameServers []string `json:"name_servers,omitempty" parquet:"name_servers,list,optional"`

	// Registrar Abuse Phone Number: The phone number for the registrar's abuse contact
	PhoneNumber *string `json:"phone_number,omitempty" parquet:"phone_number,optional"`

	// Domain Registrar: The domain registrar.
	Registrar *string `json:"registrar,omitempty" parquet:"registrar,optional"`

	// Domain Status: The status of a domain and its ability to be transferred, e.g., <code>clientTransferProhibited</code>.
	Status *string `json:"status,omitempty" parquet:"status,optional"`

	// Subdomains: An array of subdomain strings. Can be used to collect several subdomains such as those from Domain Generation Algorithms (DGAs).
	Subdomains []string `json:"subdomains,omitempty" parquet:"subdomains,list,optional"`

	// Subnet Block: The IP address block (CIDR) associated with a domain.
	Subnet *string `json:"subnet,omitempty" parquet:"subnet,optional"`
}

var WHOISFields = []arrow.Field{
	{Name: "autonomous_system", Type: AutonomousSystemStruct, Nullable: true},
	{Name: "created_time", Type: arrow.FixedWidthTypes.Timestamp_ms, Nullable: true},
	{Name: "dnssec_status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "dnssec_status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "domain_contacts", Type: arrow.ListOf(DomainContactStruct), Nullable: true},
	{Name: "email_addr", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "last_seen_time", Type: arrow.FixedWidthTypes.Timestamp_ms, Nullable: true},
	{Name: "name_servers", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "phone_number", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "registrar", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "subdomains", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "subnet", Type: arrow.BinaryTypes.String, Nullable: true},
}

var WHOISStruct = arrow.StructOf(WHOISFields...)

var WHOISSchema = arrow.NewSchema(WHOISFields, nil)
var WHOISClassname = "whois"
