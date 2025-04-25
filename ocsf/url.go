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
	Domain       *string  `json:"domain" parquet:"domain,optional" ch:"domain"	`
	Hostname     *string  `json:"hostname" parquet:"hostname,optional" ch:"hostname"`
	Path         *string  `json:"path" parquet:"path,optional" ch:"path"`
	Port         *int64   `json:"port" parquet:"port,optional" ch:"port"`
	QueryString  *string  `json:"query_string" parquet:"query_string,optional" ch:"query_string"`
	ResourceType *string  `json:"resource_type" parquet:"resource_type,optional" ch:"resource_type"`
	Scheme       *string  `json:"scheme" parquet:"scheme,optional" ch:"scheme"`
	Subdomain    *string  `json:"subdomain" parquet:"subdomain,optional" ch:"subdomain"`
	URLString    *string  `json:"url_string" parquet:"url_string,optional" ch:"url_string"`
	Categories   []string `json:"categories" parquet:"categories,list,optional" ch:"categories"`
	CategoryIDs  []int    `json:"category_ids" parquet:"category_ids,list,optional" ch:"category_ids"`
}
