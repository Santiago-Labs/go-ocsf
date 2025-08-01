// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_4_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type DNSAnswer struct {

	// Resource Record Class: The class of DNS data contained in this resource record. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: <code>IN</code>.
	Class *string `json:"class,omitempty" parquet:"class,optional"`

	// DNS Header Flags: The list of DNS answer header flag IDs.
	FlagIds []int32 `json:"flag_ids,omitempty" parquet:"flag_ids,list,optional"`

	// DNS Header Flags: The list of DNS answer header flags.
	Flags []string `json:"flags,omitempty" parquet:"flags,list,optional"`

	// Packet UID: The DNS packet identifier assigned by the program that generated the query. The identifier is copied to the response.
	PacketUid *int32 `json:"packet_uid,omitempty" parquet:"packet_uid,optional"`

	// DNS RData: The data describing the DNS resource. The meaning of this data depends on the type and class of the resource record.
	Rdata string `json:"rdata" parquet:"rdata"`

	// TTL: The time interval that the resource record may be cached. Zero value means that the resource record can only be used for the transaction in progress, and should not be cached.
	Ttl *int32 `json:"ttl,omitempty" parquet:"ttl,optional"`

	// Resource Record Type: The type of data contained in this resource record. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: <code>CNAME</code>.
	Type *string `json:"type,omitempty" parquet:"type,optional"`
}

var DNSAnswerFields = []arrow.Field{
	{Name: "class", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "flag_ids", Type: arrow.ListOf(arrow.PrimitiveTypes.Int32), Nullable: true},
	{Name: "flags", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "packet_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "rdata", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "ttl", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DNSAnswerStruct = arrow.StructOf(DNSAnswerFields...)

var DNSAnswerSchema = arrow.NewSchema(DNSAnswerFields, nil)
var DNSAnswerClassname = "dns_answer"
