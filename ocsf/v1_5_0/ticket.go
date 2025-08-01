// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type Ticket struct {

	// Source URL: The url of a ticket in the ticket system.
	SrcUrl *string `json:"src_url,omitempty" parquet:"src_url,optional"`

	// Ticket Status: The status of the ticket normalized to the caption of the <code>status_id</code> value. In the case of <code>99</code>, this value should as defined by the source.
	Status *string `json:"status,omitempty" parquet:"status,optional"`

	// Status Details: A list of contextual descriptions of the <code>status, status_id</code> values.
	StatusDetails []string `json:"status_details,omitempty" parquet:"status_details,optional,list"`

	// Ticket Status ID: The normalized identifier for the ticket status.
	StatusId *int32 `json:"status_id,omitempty" parquet:"status_id,optional"`

	// Title: The title of the ticket.
	Title *string `json:"title,omitempty" parquet:"title,optional"`

	// Ticket Type: The linked ticket type determines whether the ticket is internal or in an external ticketing system.
	Type *string `json:"type,omitempty" parquet:"type,optional"`

	// Ticket Type ID: The normalized identifier for the ticket type.
	TypeId *int32 `json:"type_id,omitempty" parquet:"type_id,optional"`

	// Unique ID: Unique identifier of the ticket.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`
}

var TicketFields = []arrow.Field{
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_details", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var TicketStruct = arrow.StructOf(TicketFields...)

var TicketSchema = arrow.NewSchema(TicketFields, nil)
