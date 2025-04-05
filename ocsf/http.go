package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// HTTPRequestFields defines the Arrow fields for HTTPRequest.
var HTTPRequestFields = []arrow.Field{
	{Name: "args", Type: arrow.BinaryTypes.String},
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32},
	{Name: "http_method", Type: arrow.BinaryTypes.String},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32},
	{Name: "referrer", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "user_agent", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct)},
	{Name: "url", Type: URLStruct},
	{Name: "x_forwarded_for", Type: arrow.ListOf(arrow.BinaryTypes.String)},
}

var HTTPRequestStruct = arrow.StructOf(HTTPRequestFields...)
var HTTPRequestClassname = "http_request"

type HTTPRequest struct {
	Args          *string       `json:"args,omitempty" parquet:"args"`
	BodyLength    *int          `json:"body_length,omitempty" parquet:"body_length"`
	HTTPMethod    *string       `json:"http_method,omitempty" parquet:"http_method"`
	Length        *int          `json:"length,omitempty" parquet:"length"`
	Referrer      *string       `json:"referrer,omitempty" parquet:"referrer"`
	UID           *string       `json:"uid,omitempty" parquet:"uid"`
	UserAgent     *string       `json:"user_agent,omitempty" parquet:"user_agent"`
	Version       *string       `json:"version,omitempty" parquet:"version"`
	HTTPHeaders   []*HTTPHeader `json:"http_headers,omitempty" parquet:"http_headers"`
	URL           *URL          `json:"url,omitempty" parquet:"url"`
	XForwardedFor []*string     `json:"x_forwarded_for,omitempty" parquet:"x_forwarded_for"`
}

// HTTPResponseFields defines the Arrow fields for HTTPResponse.
var HTTPResponseFields = []arrow.Field{
	{Name: "body_length", Type: arrow.PrimitiveTypes.Int32},
	{Name: "code", Type: arrow.PrimitiveTypes.Int32},
	{Name: "content_type", Type: arrow.BinaryTypes.String},
	{Name: "latency", Type: arrow.PrimitiveTypes.Int32},
	{Name: "length", Type: arrow.PrimitiveTypes.Int32},
	{Name: "message", Type: arrow.BinaryTypes.String},
	{Name: "status", Type: arrow.BinaryTypes.String},
	{Name: "http_headers", Type: arrow.ListOf(HTTPHeaderStruct)},
}

var HTTPResponseStruct = arrow.StructOf(HTTPResponseFields...)
var HTTPResponseClassname = "http_response"

type HTTPResponse struct {
	BodyLength  *int          `json:"body_length,omitempty" parquet:"body_length"`
	Code        int           `json:"code" parquet:"code"`
	ContentType *string       `json:"content_type,omitempty" parquet:"content_type"`
	Latency     *int          `json:"latency,omitempty" parquet:"latency"`
	Length      *int          `json:"length,omitempty" parquet:"length"`
	Message     *string       `json:"message,omitempty" parquet:"message"`
	Status      *string       `json:"status,omitempty" parquet:"status"`
	HTTPHeaders []*HTTPHeader `json:"http_headers,omitempty" parquet:"http_headers"`
}

// HTTPHeaderFields defines the Arrow fields for HTTPHeader.
var HTTPHeaderFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "value", Type: arrow.BinaryTypes.String},
}

var HTTPHeaderStruct = arrow.StructOf(HTTPHeaderFields...)
var HTTPHeaderClassname = "http_header"

type HTTPHeader struct {
	Name  string `json:"name" parquet:"name"`
	Value string `json:"value" parquet:"value"`
}
