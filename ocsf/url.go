package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// URLFields defines the Arrow fields for URL.
var URLFields = []arrow.Field{
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "hostname", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "port", Type: arrow.PrimitiveTypes.Int32},
	{Name: "query_string", Type: arrow.BinaryTypes.String},
	{Name: "resource_type", Type: arrow.BinaryTypes.String},
	{Name: "scheme", Type: arrow.BinaryTypes.String},
	{Name: "subdomain", Type: arrow.BinaryTypes.String},
	{Name: "url_string", Type: arrow.BinaryTypes.String},
	{Name: "categories", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "category_ids", Type: arrow.ListOf(arrow.PrimitiveTypes.Int32)},
}

var URLStruct = arrow.StructOf(URLFields...)

type URL struct {
	Domain       *string   `json:"domain,omitempty" parquet:"domain"`
	Hostname     *string   `json:"hostname,omitempty" parquet:"hostname"`
	Path         *string   `json:"path,omitempty" parquet:"path"`
	Port         *int      `json:"port,omitempty" parquet:"port"`
	QueryString  *string   `json:"query_string,omitempty" parquet:"query_string"`
	ResourceType *string   `json:"resource_type,omitempty" parquet:"resource_type"`
	Scheme       *string   `json:"scheme,omitempty" parquet:"scheme"`
	Subdomain    *string   `json:"subdomain,omitempty" parquet:"subdomain"`
	URLString    *string   `json:"url_string,omitempty" parquet:"url_string"`
	Categories   []*string `json:"categories,omitempty" parquet:"categories"`
	CategoryIDs  []*int    `json:"category_ids,omitempty" parquet:"category_ids"`
}
