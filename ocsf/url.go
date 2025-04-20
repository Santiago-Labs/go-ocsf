package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// URLFields defines the Arrow fields for URL.
var URLFields = []arrow.Field{
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hostname", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "port", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "query_string", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "resource_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "scheme", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "subdomain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "url_string", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "categories", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "category_ids", Type: arrow.ListOf(arrow.PrimitiveTypes.Int32), Nullable: true},
}

var URLStruct = arrow.StructOf(URLFields...)
var URLClassname = "url"

type URL struct {
	Domain       *string   `json:"domain,omitempty" parquet:"domain,optional"`
	Hostname     *string   `json:"hostname,omitempty" parquet:"hostname,optional"`
	Path         *string   `json:"path,omitempty" parquet:"path,optional"`
	Port         *int      `json:"port,omitempty" parquet:"port,optional"`
	QueryString  *string   `json:"query_string,omitempty" parquet:"query_string,optional"`
	ResourceType *string   `json:"resource_type,omitempty" parquet:"resource_type,optional"`
	Scheme       *string   `json:"scheme,omitempty" parquet:"scheme,optional"`
	Subdomain    *string   `json:"subdomain,omitempty" parquet:"subdomain,optional"`
	URLString    *string   `json:"url_string,omitempty" parquet:"url_string,optional"`
	Categories   []*string `json:"categories,omitempty" parquet:"categories,list,optional"`
	CategoryIDs  []*int    `json:"category_ids,omitempty" parquet:"category_ids,list,optional"`
}
