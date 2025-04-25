package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// HTTPRequestFields defines the Arrow fields for HTTPRequest.
var HTTPRequestFields = []arrow.Field{
	{Name: "args", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "http_method", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "referrer", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "user_agent", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct), Nullable: true},
	{Name: "url", Type: URLStruct, Nullable: true},
	{Name: "x_forwarded_for", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
}

var HTTPRequestStruct = arrow.StructOf(HTTPRequestFields...)
var HTTPRequestClassname = "http_request"

type HTTPRequest struct {
	Args          *string       `json:"args" parquet:"args" ch:"args"`
	BodyLength    *int64        `json:"body_length" parquet:"body_length" ch:"body_length"`
	HTTPMethod    *string       `json:"http_method" parquet:"http_method" ch:"http_method"`
	Length        *int64        `json:"length" parquet:"length" ch:"length"`
	Referrer      *string       `json:"referrer" parquet:"referrer" ch:"referrer"`
	UID           *string       `json:"uid" parquet:"uid" ch:"uid"`
	UserAgent     *string       `json:"user_agent" parquet:"user_agent" ch:"user_agent"`
	Version       *string       `json:"version" parquet:"version,optional" ch:"version"`
	HTTPHeaders   []*HTTPHeader `json:"http_headers" parquet:"http_headers,list,optional" ch:"http_headers"`
	URL           *URL          `json:"url" parquet:"url" ch:"url"`
	XForwardedFor []string      `json:"x_forwarded_for" parquet:"x_forwarded_for,list,optional" ch:"x_forwarded_for"`
}

// HTTPResponseFields defines the Arrow fields for HTTPResponse.
var HTTPResponseFields = []arrow.Field{
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "code", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "content_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "latency", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct), Nullable: true},
}

var HTTPResponseStruct = arrow.StructOf(HTTPResponseFields...)
var HTTPResponseClassname = "http_response"

type HTTPResponse struct {
	BodyLength  *int64        `json:"body_length" parquet:"body_length,optional" ch:"body_length"`
	Code        int           `json:"code" parquet:"code" ch:"code"`
	ContentType *string       `json:"content_type" parquet:"content_type,optional" ch:"content_type"`
	Latency     *int64        `json:"latency" parquet:"latency,optional" ch:"latency"`
	Length      *int64        `json:"length" parquet:"length,optional" ch:"length"`
	Message     *string       `json:"message" parquet:"message,optional" ch:"message"`
	Status      *string       `json:"status" parquet:"status,optional" ch:"status"`
	HTTPHeaders []*HTTPHeader `json:"http_headers" parquet:"http_headers,list,optional" ch:"http_headers"`
}

// HTTPHeaderFields defines the Arrow fields for HTTPHeader.
var HTTPHeaderFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "value", Type: arrow.BinaryTypes.String, Nullable: false},
}

var HTTPHeaderStruct = arrow.StructOf(HTTPHeaderFields...)
var HTTPHeaderClassname = "http_header"

type HTTPHeader struct {
	Name  string `json:"name" parquet:"name" ch:"name"`
	Value string `json:"value" parquet:"value" ch:"value"`
}
