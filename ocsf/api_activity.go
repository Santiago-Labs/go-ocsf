package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// APIActivityFields defines the Arrow fields for APIActivity.
var APIActivityFields = []arrow.Field{
	{Name: "activity_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "activity_name", Type: arrow.BinaryTypes.String},
	{Name: "actor", Type: ActorStruct},
	{Name: "category_name", Type: arrow.BinaryTypes.String},
	{Name: "category_uid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "class_name", Type: arrow.BinaryTypes.String},
	{Name: "class_uid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "count", Type: arrow.PrimitiveTypes.Int32},
	{Name: "duration", Type: arrow.PrimitiveTypes.Int64},
	{Name: "end_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "message", Type: arrow.BinaryTypes.String},
	{Name: "raw_data", Type: arrow.BinaryTypes.String},
	{Name: "severity", Type: arrow.BinaryTypes.String},
	{Name: "severity_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "start_time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "status", Type: arrow.BinaryTypes.String},
	{Name: "status_code", Type: arrow.BinaryTypes.String},
	{Name: "status_detail", Type: arrow.BinaryTypes.String},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "time", Type: arrow.PrimitiveTypes.Int64},
	{Name: "timezone_offset", Type: arrow.PrimitiveTypes.Int32},
	{Name: "type_name", Type: arrow.BinaryTypes.String},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int32},
	{Name: "dst_endpoint", Type: NetworkEndpointStruct},
	{Name: "enrichments", Type: arrow.ListOf(EnrichmentStruct)},
	{Name: "http_request", Type: HTTPRequestStruct},
	{Name: "http_response", Type: HTTPResponseStruct},
	{Name: "metadata", Type: MetadataStruct},
	{Name: "observables", Type: arrow.ListOf(ObservableStruct)},
	{Name: "resources", Type: arrow.ListOf(ResourceDetailsStruct)},
	{Name: "src_endpoint", Type: NetworkEndpointStruct},
	{Name: "unmapped", Type: arrow.MapOf(arrow.BinaryTypes.String, arrow.BinaryTypes.String)},
}

var APIActivityStruct = arrow.StructOf(APIActivityFields...)
var APIActivityClassname = "api_activity"

type APIActivity struct {
	ActivityID     int                    `json:"activity_id" parquet:"activity_id"`
	ActivityName   *string                `json:"activity_name,omitempty" parquet:"activity_name"`
	Actor          Actor                  `json:"actor,omitempty" parquet:"actor"`
	API            API                    `json:"api" parquet:"api"`
	CategoryName   *string                `json:"category_name,omitempty" parquet:"category_name"`
	CategoryUID    int                    `json:"category_uid" parquet:"category_uid"`
	ClassName      *string                `json:"class_name,omitempty" parquet:"class_name"`
	ClassUID       int                    `json:"class_uid" parquet:"class_uid"`
	Count          *int                   `json:"count,omitempty" parquet:"count"`
	Duration       *int64                 `json:"duration,omitempty" parquet:"duration"`
	EndTime        *time.Time             `json:"end_time,omitempty" parquet:"end_time"`
	Message        *string                `json:"message,omitempty" parquet:"message"`
	RawData        *string                `json:"raw_data,omitempty" parquet:"raw_data"`
	Severity       *string                `json:"severity,omitempty" parquet:"severity"`
	SeverityID     int                    `json:"severity_id" parquet:"severity_id"`
	StartTime      *time.Time             `json:"start_time,omitempty" parquet:"start_time"`
	Status         *string                `json:"status,omitempty" parquet:"status"`
	StatusCode     *string                `json:"status_code,omitempty" parquet:"status_code"`
	StatusDetail   *string                `json:"status_detail,omitempty" parquet:"status_detail"`
	StatusID       int                    `json:"status_id" parquet:"status_id"`
	Time           time.Time              `json:"time" parquet:"time"`
	TimezoneOffset int                    `json:"timezone_offset" parquet:"timezone_offset"`
	TypeName       *string                `json:"type_name,omitempty" parquet:"type_name"`
	TypeUID        int                    `json:"type_uid" parquet:"type_uid"`
	DstEndpoint    *NetworkEndpoint       `json:"dst_endpoint,omitempty" parquet:"dst_endpoint"`
	Enrichments    []*Enrichment          `json:"enrichments,omitempty" parquet:"enrichments"`
	HTTPRequest    *HTTPRequest           `json:"http_request,omitempty" parquet:"http_request"`
	HTTPResponse   *HTTPResponse          `json:"http_response,omitempty" parquet:"http_response"`
	Metadata       Metadata               `json:"metadata" parquet:"metadata"`
	Observables    []*Observable          `json:"observables,omitempty" parquet:"observables"`
	Resources      []*ResourceDetails     `json:"resources,omitempty" parquet:"resources"`
	SrcEndpoint    NetworkEndpoint        `json:"src_endpoint" parquet:"src_endpoint"`
	Unmapped       map[string]interface{} `json:"unmapped,omitempty" parquet:"unmapped"`
}
