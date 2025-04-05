package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// APIFields defines the Arrow fields for API.
var APIFields = []arrow.Field{
	{Name: "group", Type: GroupStruct},
	{Name: "operation", Type: arrow.BinaryTypes.String},
	{Name: "request", Type: RequestStruct},
	{Name: "response", Type: ResponseStruct},
	{Name: "service", Type: ServiceStruct},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var APIStruct = arrow.StructOf(APIFields...)
var APIClassname = "api"

type API struct {
	Group     *Group    `json:"group,omitempty" parquet:"group"`
	Operation string    `json:"operation" parquet:"operation"`
	Request   *Request  `json:"request,omitempty" parquet:"request"`
	Response  *Response `json:"response,omitempty" parquet:"response"`
	Service   *Service  `json:"service,omitempty" parquet:"service"`
	Version   *string   `json:"version,omitempty" parquet:"version"`
}
