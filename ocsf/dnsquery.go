package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var DNSQueryFields = []arrow.Field{
	{Name: "class", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hostname", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "opcode", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "opcode_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "packet_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DNSQueryStruct = arrow.StructOf(DNSQueryFields...)
var DNSQueryClassname = "dns_query"

type DNSQuery struct {
	Class     *string `json:"class,omitempty" parquet:"class,optional"`
	Hostname  string  `json:"hostname" parquet:"hostname"`
	Opcode    *string `json:"opcode,omitempty" parquet:"opcode,optional"`
	OpcodeID  *int32  `json:"opcode_id,omitempty" parquet:"opcode_id,optional"`
	PacketUID *int32  `json:"packet_uid,omitempty" parquet:"packet_uid,optional"`
	Type      *string `json:"type,omitempty" parquet:"type,optional"`
}
