package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// APIFields defines the Arrow fields for API.
var APIFields = []arrow.Field{
	{Name: "group", Type: GroupStruct, Nullable: true},
	{Name: "operation", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "request", Type: RequestStruct, Nullable: true},
	{Name: "response", Type: ResponseStruct, Nullable: true},
	{Name: "service", Type: ServiceStruct, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var APIStruct = arrow.StructOf(APIFields...)
var APIClassname = "api"

type API struct {
	Group     *Group    `json:"group,omitempty" parquet:"group,optional" ch:"group,omitempty"`
	Operation string    `json:"operation" parquet:"operation" ch:"operation"`
	Request   *Request  `json:"request,omitempty" parquet:"request,optional" ch:"request,omitempty"`
	Response  *Response `json:"response,omitempty" parquet:"response,optional" ch:"response,omitempty"`
	Service   *Service  `json:"service,omitempty" parquet:"service,optional" ch:"service,omitempty"`
	Version   *string   `json:"version,omitempty" parquet:"version,optional" ch:"version,omitempty"`
}
