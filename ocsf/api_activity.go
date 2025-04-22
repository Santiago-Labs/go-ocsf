package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// APIActivityFields defines the Arrow fields for APIActivity.
var APIActivityFields = []arrow.Field{
	{Name: "event_day", Type: arrow.FixedWidthTypes.Date32, Nullable: false},
	{Name: "activity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "activity_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "actor", Type: ActorStruct, Nullable: false},
	{Name: "api", Type: APIStruct, Nullable: false},
	{Name: "category_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "category_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "class_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "class_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "count", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "duration", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "end_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "raw_data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "start_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_code", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_detail", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "time", Type: arrow.FixedWidthTypes.Timestamp_s, Nullable: false},
	{Name: "timezone_offset", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "type_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "dst_endpoint", Type: NetworkEndpointStruct, Nullable: true},
	{Name: "enrichments", Type: arrow.ListOf(EnrichmentStruct), Nullable: true},
	{Name: "http_request", Type: HTTPRequestStruct, Nullable: true},
	{Name: "http_response", Type: HTTPResponseStruct, Nullable: true},
	{Name: "metadata", Type: MetadataStruct, Nullable: false},
	{Name: "observables", Type: arrow.ListOf(ObservableStruct), Nullable: true},
	{Name: "resources", Type: arrow.ListOf(ResourceDetailsStruct), Nullable: true},
	{Name: "src_endpoint", Type: NetworkEndpointStruct, Nullable: false},
	{Name: "unmapped", Type: arrow.BinaryTypes.String, Nullable: true},
}

var APIActivityStruct = arrow.StructOf(APIActivityFields...)
var APIActivityClassname = "api_activity"
var APIActivitySchema = arrow.NewSchema(APIActivityFields, nil)

type APIActivity struct {
	EventDay       int32              `json:"event_day" parquet:"event_day,date"` // Used for partitioning
	ActivityID     int                `json:"activity_id" parquet:"activity_id"`
	ActivityName   *string            `json:"activity_name,omitempty" parquet:"activity_name,optional"`
	Actor          Actor              `json:"actor,omitempty" parquet:"actor"`
	API            API                `json:"api" parquet:"api"`
	CategoryName   *string            `json:"category_name,omitempty" parquet:"category_name,optional"`
	CategoryUID    int                `json:"category_uid" parquet:"category_uid"`
	ClassName      *string            `json:"class_name,omitempty" parquet:"class_name,optional"`
	ClassUID       int                `json:"class_uid" parquet:"class_uid"`
	Count          *int               `json:"count,omitempty" parquet:"count,optional"`
	Duration       *int64             `json:"duration,omitempty" parquet:"duration,optional"`
	EndTime        *int64             `json:"end_time,omitempty" parquet:"end_time,optional"`
	Message        *string            `json:"message,omitempty" parquet:"message,optional"`
	RawData        *string            `json:"raw_data,omitempty" parquet:"raw_data,optional"`
	Severity       *string            `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID     int                `json:"severity_id" parquet:"severity_id"`
	StartTime      *int64             `json:"start_time,omitempty" parquet:"start_time,optional"`
	Status         *string            `json:"status,omitempty" parquet:"status,optional"`
	StatusCode     *string            `json:"status_code,omitempty" parquet:"status_code,optional"`
	StatusDetail   *string            `json:"status_detail,omitempty" parquet:"status_detail,optional"`
	StatusID       int                `json:"status_id" parquet:"status_id"`
	Time           int64              `json:"time" parquet:"time,timestamp"`
	TimezoneOffset int                `json:"timezone_offset" parquet:"timezone_offset"`
	TypeName       *string            `json:"type_name,omitempty" parquet:"type_name,optional"`
	TypeUID        int                `json:"type_uid" parquet:"type_uid"`
	DstEndpoint    *NetworkEndpoint   `json:"dst_endpoint,omitempty" parquet:"dst_endpoint,optional"`
	Enrichments    []*Enrichment      `json:"enrichments,omitempty" parquet:"enrichments,list,optional"`
	HTTPRequest    *HTTPRequest       `json:"http_request,omitempty" parquet:"http_request,optional"`
	HTTPResponse   *HTTPResponse      `json:"http_response,omitempty" parquet:"http_response,optional"`
	Metadata       Metadata           `json:"metadata" parquet:"metadata"`
	Observables    []*Observable      `json:"observables,omitempty" parquet:"observables,list,optional"`
	Resources      []*ResourceDetails `json:"resources,omitempty" parquet:"resources,list,optional"`
	SrcEndpoint    NetworkEndpoint    `json:"src_endpoint" parquet:"src_endpoint"`
	Unmapped       *string            `json:"unmapped,omitempty" parquet:"unmapped,optional"`
}

type APIActivityDTO struct {
	EventDay       int32   `json:"event_day" parquet:"event_day,date"` // Used for partitioning
	ActivityID     int     `json:"activity_id" parquet:"activity_id"`
	ActivityName   *string `json:"activity_name,omitempty" parquet:"activity_name,optional"`
	Actor          JSONB   `json:"actor,omitempty" parquet:"actor"`
	API            JSONB   `json:"api" parquet:"api"`
	CategoryName   *string `json:"category_name,omitempty" parquet:"category_name,optional"`
	CategoryUID    int     `json:"category_uid" parquet:"category_uid"`
	ClassName      *string `json:"class_name,omitempty" parquet:"class_name,optional"`
	ClassUID       int     `json:"class_uid" parquet:"class_uid"`
	Count          *int    `json:"count,omitempty" parquet:"count,optional"`
	Duration       *int64  `json:"duration,omitempty" parquet:"duration,optional"`
	EndTime        *int64  `json:"end_time,omitempty" parquet:"end_time,optional"`
	Message        *string `json:"message,omitempty" parquet:"message,optional"`
	RawData        *string `json:"raw_data,omitempty" parquet:"raw_data,optional"`
	Severity       *string `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID     int     `json:"severity_id" parquet:"severity_id"`
	StartTime      *int64  `json:"start_time,omitempty" parquet:"start_time,optional"`
	Status         *string `json:"status,omitempty" parquet:"status,optional"`
	StatusCode     *string `json:"status_code,omitempty" parquet:"status_code,optional"`
	StatusDetail   *string `json:"status_detail,omitempty" parquet:"status_detail,optional"`
	StatusID       int     `json:"status_id" parquet:"status_id"`
	Time           int64   `json:"time" parquet:"time"`
	TimezoneOffset int     `json:"timezone_offset" parquet:"timezone_offset"`
	TypeName       *string `json:"type_name,omitempty" parquet:"type_name,optional"`
	TypeUID        int     `json:"type_uid" parquet:"type_uid"`
	DstEndpoint    JSONB   `json:"dst_endpoint,omitempty" parquet:"dst_endpoint,optional"`
	Enrichments    JSONB   `json:"enrichments,omitempty" parquet:"enrichments,list,optional"`
	HTTPRequest    JSONB   `json:"http_request,omitempty" parquet:"http_request,optional"`
	HTTPResponse   JSONB   `json:"http_response,omitempty" parquet:"http_response,optional"`
	Metadata       JSONB   `json:"metadata" parquet:"metadata"`
	Observables    JSONB   `json:"observables,omitempty" parquet:"observables,list,optional"`
	Resources      JSONB   `json:"resources,omitempty" parquet:"resources,list,optional"`
	SrcEndpoint    JSONB   `json:"src_endpoint" parquet:"src_endpoint"`
	Unmapped       *string `json:"unmapped,omitempty" parquet:"unmapped,optional"`
}
