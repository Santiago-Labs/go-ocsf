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

type APIActivity struct {
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
	Unmapped       map[string]string  `json:"unmapped,omitempty" parquet:"unmapped"`
}

func (a *APIActivity) Equals(other *APIActivity) bool {
	if a == nil || other == nil {
		return a == other
	}

	// Compare primitive fields
	if a.ActivityID != other.ActivityID ||
		a.CategoryUID != other.CategoryUID ||
		a.ClassUID != other.ClassUID ||
		a.SeverityID != other.SeverityID ||
		a.StatusID != other.StatusID ||
		!a.Time.Equal(other.Time) ||
		a.TimezoneOffset != other.TimezoneOffset ||
		a.TypeUID != other.TypeUID {
		return false
	}

	// Compare pointer fields with safe dereferencing
	if !pointerStringsEqual(a.ActivityName, other.ActivityName) ||
		!pointerStringsEqual(a.CategoryName, other.CategoryName) ||
		!pointerStringsEqual(a.ClassName, other.ClassName) ||
		!pointerStringsEqual(a.Message, other.Message) ||
		!pointerStringsEqual(a.RawData, other.RawData) ||
		!pointerStringsEqual(a.Severity, other.Severity) ||
		!pointerStringsEqual(a.Status, other.Status) ||
		!pointerStringsEqual(a.StatusCode, other.StatusCode) ||
		!pointerStringsEqual(a.StatusDetail, other.StatusDetail) ||
		!pointerStringsEqual(a.TypeName, other.TypeName) {
		return false
	}

	// Compare pointer to int
	if !pointerIntsEqual(a.Count, other.Count) {
		return false
	}

	// Compare pointer to int64
	if !pointerInt64sEqual(a.Duration, other.Duration) {
		return false
	}

	// Compare pointer to time.Time
	if !pointerTimesEqual(a.EndTime, other.EndTime) ||
		!pointerTimesEqual(a.StartTime, other.StartTime) {
		return false
	}

	// Compare maps
	return mapsEqual(a.Unmapped, other.Unmapped)
}

// Helper functions for safe comparison

func pointerStringsEqual(a, b *string) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func pointerIntsEqual(a, b *int) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func pointerInt64sEqual(a, b *int64) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

func pointerTimesEqual(a, b *time.Time) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Equal(*b)
}

func pointerStructsEqual[T any](a, b *T, equals func(*T, *T) bool) bool {
	if a == nil || b == nil {
		return a == b
	}
	return equals(a, b)
}

func slicesEqual[T any](a, b []*T, equals func(*T, *T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] == nil || b[i] == nil {
			if a[i] != b[i] {
				return false
			}
			continue
		}
		if !equals(a[i], b[i]) {
			return false
		}
	}
	return true
}

func mapsEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}
