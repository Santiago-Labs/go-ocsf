package ocsf

import (
	"encoding/json"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// APIActivityFields defines the Arrow fields for APIActivity.
var APIActivityFields = []arrow.Field{
	{Name: "event_day", Type: arrow.FixedWidthTypes.Date64, Nullable: false},
	{Name: "filename", Type: arrow.BinaryTypes.String, Nullable: false},
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
	{Name: "time", Type: arrow.PrimitiveTypes.Date64, Nullable: false},
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

type APIActivity struct {
	EventDay       time.Time          `json:"event_day" parquet:"event_day"` // Used for partitioning
	Filename       string             `json:"filename" parquet:"filename"`   // Used for partitioning
	ActivityID     int                `json:"activity_id" parquet:"activity_id"`
	ActivityName   *string            `json:"activity_name,omitempty" parquet:"activity_name"`
	Actor          Actor              `json:"actor,omitempty" parquet:"actor"`
	API            API                `json:"api" parquet:"api"`
	CategoryName   *string            `json:"category_name,omitempty" parquet:"category_name"`
	CategoryUID    int                `json:"category_uid" parquet:"category_uid"`
	ClassName      *string            `json:"class_name,omitempty" parquet:"class_name"`
	ClassUID       int                `json:"class_uid" parquet:"class_uid"`
	Count          *int               `json:"count,omitempty" parquet:"count"`
	Duration       *int64             `json:"duration,omitempty" parquet:"duration"`
	EndTime        *time.Time         `json:"end_time,omitempty" parquet:"end_time"`
	Message        *string            `json:"message,omitempty" parquet:"message"`
	RawData        *string            `json:"raw_data,omitempty" parquet:"raw_data"`
	Severity       *string            `json:"severity,omitempty" parquet:"severity"`
	SeverityID     int                `json:"severity_id" parquet:"severity_id"`
	StartTime      *time.Time         `json:"start_time,omitempty" parquet:"start_time"`
	Status         *string            `json:"status,omitempty" parquet:"status"`
	StatusCode     *string            `json:"status_code,omitempty" parquet:"status_code"`
	StatusDetail   *string            `json:"status_detail,omitempty" parquet:"status_detail"`
	StatusID       int                `json:"status_id" parquet:"status_id"`
	Time           time.Time          `json:"time" parquet:"time"`
	TimezoneOffset int                `json:"timezone_offset" parquet:"timezone_offset"`
	TypeName       *string            `json:"type_name,omitempty" parquet:"type_name"`
	TypeUID        int                `json:"type_uid" parquet:"type_uid"`
	DstEndpoint    *NetworkEndpoint   `json:"dst_endpoint,omitempty" parquet:"dst_endpoint"`
	Enrichments    []*Enrichment      `json:"enrichments,omitempty" parquet:"enrichments"`
	HTTPRequest    *HTTPRequest       `json:"http_request,omitempty" parquet:"http_request"`
	HTTPResponse   *HTTPResponse      `json:"http_response,omitempty" parquet:"http_response"`
	Metadata       Metadata           `json:"metadata" parquet:"metadata"`
	Observables    []*Observable      `json:"observables,omitempty" parquet:"observables"`
	Resources      []*ResourceDetails `json:"resources,omitempty" parquet:"resources"`
	SrcEndpoint    NetworkEndpoint    `json:"src_endpoint" parquet:"src_endpoint"`
	Unmapped       string             `json:"unmapped,omitempty" parquet:"unmapped"`
}

type APIActivityDTO struct {
	EventDay       time.Time `json:"event_day" parquet:"event_day"` // Used for partitioning
	Filename       string    `json:"filename" parquet:"filename"`   // Used for partitioning
	ActivityID     int       `json:"activity_id" parquet:"activity_id"`
	ActivityName   *string   `json:"activity_name,omitempty" parquet:"activity_name"`
	Actor          JSONB     `json:"actor,omitempty" parquet:"actor"`
	API            JSONB     `json:"api" parquet:"api"`
	CategoryName   *string   `json:"category_name,omitempty" parquet:"category_name"`
	CategoryUID    int       `json:"category_uid" parquet:"category_uid"`
	ClassName      *string   `json:"class_name,omitempty" parquet:"class_name"`
	ClassUID       int       `json:"class_uid" parquet:"class_uid"`
	Count          *int      `json:"count,omitempty" parquet:"count"`
	Duration       *int64    `json:"duration,omitempty" parquet:"duration"`
	EndTime        *DBTime   `json:"end_time,omitempty" parquet:"end_time"`
	Message        *string   `json:"message,omitempty" parquet:"message"`
	RawData        *string   `json:"raw_data,omitempty" parquet:"raw_data"`
	Severity       *string   `json:"severity,omitempty" parquet:"severity"`
	SeverityID     int       `json:"severity_id" parquet:"severity_id"`
	StartTime      *DBTime   `json:"start_time,omitempty" parquet:"start_time"`
	Status         *string   `json:"status,omitempty" parquet:"status"`
	StatusCode     *string   `json:"status_code,omitempty" parquet:"status_code"`
	StatusDetail   *string   `json:"status_detail,omitempty" parquet:"status_detail"`
	StatusID       int       `json:"status_id" parquet:"status_id"`
	Time           time.Time `json:"time" parquet:"time"`
	TimezoneOffset int       `json:"timezone_offset" parquet:"timezone_offset"`
	TypeName       *string   `json:"type_name,omitempty" parquet:"type_name"`
	TypeUID        int       `json:"type_uid" parquet:"type_uid"`
	DstEndpoint    JSONB     `json:"dst_endpoint,omitempty" parquet:"dst_endpoint"`
	Enrichments    JSONB     `json:"enrichments,omitempty" parquet:"enrichments"`
	HTTPRequest    JSONB     `json:"http_request,omitempty" parquet:"http_request"`
	HTTPResponse   JSONB     `json:"http_response,omitempty" parquet:"http_response"`
	Metadata       JSONB     `json:"metadata" parquet:"metadata"`
	Observables    JSONB     `json:"observables,omitempty" parquet:"observables"`
	Resources      JSONB     `json:"resources,omitempty" parquet:"resources"`
	SrcEndpoint    JSONB     `json:"src_endpoint" parquet:"src_endpoint"`
	Unmapped       string    `json:"unmapped,omitempty" parquet:"unmapped"`
}

func (dto *APIActivityDTO) ToStruct() (*APIActivity, error) {
	var activity APIActivity

	activity.EventDay = dto.EventDay
	activity.Filename = dto.Filename
	activity.ActivityID = dto.ActivityID
	activity.ActivityName = dto.ActivityName
	activity.CategoryName = dto.CategoryName
	activity.CategoryUID = dto.CategoryUID
	activity.ClassName = dto.ClassName
	activity.ClassUID = dto.ClassUID
	activity.Count = dto.Count
	activity.Duration = dto.Duration
	activity.Message = dto.Message
	activity.RawData = dto.RawData
	activity.Severity = dto.Severity
	activity.SeverityID = dto.SeverityID
	activity.Status = dto.Status
	activity.StatusCode = dto.StatusCode
	activity.StatusDetail = dto.StatusDetail
	activity.StatusID = dto.StatusID
	activity.Time = dto.Time
	activity.TimezoneOffset = dto.TimezoneOffset
	activity.TypeName = dto.TypeName
	activity.TypeUID = dto.TypeUID
	activity.Unmapped = dto.Unmapped

	if dto.EndTime != nil && !dto.EndTime.IsZero() {
		activity.EndTime = &dto.EndTime.Time
	}

	if dto.StartTime != nil && !dto.StartTime.IsZero() {
		activity.StartTime = &dto.StartTime.Time
	}

	if err := json.Unmarshal(dto.Actor, &activity.Actor); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.API, &activity.API); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.DstEndpoint, &activity.DstEndpoint); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.Enrichments, &activity.Enrichments); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.HTTPRequest, &activity.HTTPRequest); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.HTTPResponse, &activity.HTTPResponse); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.Metadata, &activity.Metadata); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.Observables, &activity.Observables); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.Resources, &activity.Resources); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dto.SrcEndpoint, &activity.SrcEndpoint); err != nil {
		return nil, err
	}

	return &activity, nil
}
