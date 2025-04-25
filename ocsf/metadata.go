package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// Assume the following schemas are defined elsewhere:
// SchemaExtensionSchema, LoggerSchema, and ProductSchema.

// MetadataFields defines the Arrow fields for Metadata.
var MetadataFields = []arrow.Field{
	{Name: "correlation_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "event_code", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "extension", Type: SchemaExtensionStruct, Nullable: true},
	{Name: "extensions", Type: arrow.ListOf(SchemaExtensionStruct), Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "log_level", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "logged_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "loggers", Type: arrow.ListOf(LoggerStruct), Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "original_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "processed_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: false},
	{Name: "profiles", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "sequence", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "tenant_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var MetadataStruct = arrow.StructOf(MetadataFields...)
var MetadataClassname = "metadata"

type Metadata struct {
	CorrelationUID *string            `json:"correlation_uid" parquet:"correlation_uid,optional" ch:"correlation_uid"`
	EventCode      *string            `json:"event_code" parquet:"event_code,optional" ch:"event_code"`
	Extension      *SchemaExtension   `json:"extension" parquet:"extension,optional" ch:"extension"`
	Extensions     []*SchemaExtension `json:"extensions" parquet:"extensions,list,optional" ch:"extensions"`
	Labels         []string           `json:"labels" parquet:"labels,list,optional" ch:"labels"`
	LogLevel       *string            `json:"log_level" parquet:"log_level,optional" ch:"log_level"`
	LogName        *string            `json:"log_name" parquet:"log_name,optional" ch:"log_name"`
	LogProvider    *string            `json:"log_provider" parquet:"log_provider,optional" ch:"log_provider"`
	LogVersion     *string            `json:"log_version" parquet:"log_version,optional" ch:"log_version"`
	LoggedTime     *int64             `json:"logged_time" parquet:"logged_time,optional" ch:"logged_time"`
	Loggers        []*Logger          `json:"loggers" parquet:"loggers,list,optional" ch:"loggers"`
	ModifiedTime   *int64             `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	OriginalTime   *int64             `json:"original_time" parquet:"original_time,optional" ch:"original_time"`
	ProcessedTime  *int64             `json:"processed_time" parquet:"processed_time,optional" ch:"processed_time"`
	Product        Product            `json:"product" parquet:"product" ch:"product"`
	Profiles       []string           `json:"profiles" parquet:"profiles,list,optional" ch:"profiles"`
	Sequence       *int64             `json:"sequence" parquet:"sequence,optional" ch:"sequence"`
	TenantUID      *string            `json:"tenant_uid" parquet:"tenant_uid,optional" ch:"tenant_uid"`
	UID            *string            `json:"uid" parquet:"uid,optional" ch:"uid"`
	Version        string             `json:"version" parquet:"version" ch:"version"`
}
