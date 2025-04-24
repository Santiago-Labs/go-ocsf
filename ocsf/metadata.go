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
	CorrelationUID *string            `json:"correlation_uid,omitempty" parquet:"correlation_uid,optional" ch:"correlation_uid"`
	EventCode      *string            `json:"event_code,omitempty" parquet:"event_code,optional" ch:"event_code"`
	Extension      *SchemaExtension   `json:"extension,omitempty" parquet:"extension,optional" ch:"extension"`
	Extensions     []*SchemaExtension `json:"extensions,omitempty" parquet:"extensions,list,optional" ch:"extensions"`
	Labels         []string           `json:"labels,omitempty" parquet:"labels,list,optional" ch:"labels"`
	LogLevel       *string            `json:"log_level,omitempty" parquet:"log_level,optional" ch:"log_level"`
	LogName        *string            `json:"log_name,omitempty" parquet:"log_name,optional" ch:"log_name"`
	LogProvider    *string            `json:"log_provider,omitempty" parquet:"log_provider,optional" ch:"log_provider"`
	LogVersion     *string            `json:"log_version,omitempty" parquet:"log_version,optional" ch:"log_version"`
	LoggedTime     *int64             `json:"logged_time,omitempty" parquet:"logged_time,optional" ch:"logged_time"`
	Loggers        []*Logger          `json:"loggers,omitempty" parquet:"loggers,list,optional" ch:"loggers"`
	ModifiedTime   *int64             `json:"modified_time,omitempty" parquet:"modified_time,optional" ch:"modified_time"`
	OriginalTime   *int64             `json:"original_time,omitempty" parquet:"original_time,optional" ch:"original_time"`
	ProcessedTime  *int64             `json:"processed_time,omitempty" parquet:"processed_time,optional" ch:"processed_time"`
	Product        Product            `json:"product" parquet:"product" ch:"product"`
	Profiles       []string           `json:"profiles,omitempty" parquet:"profiles,list,optional" ch:"profiles"`
	Sequence       *int               `json:"sequence,omitempty" parquet:"sequence,optional" ch:"sequence"`
	TenantUID      *string            `json:"tenant_uid,omitempty" parquet:"tenant_uid,optional" ch:"tenant_uid"`
	UID            *string            `json:"uid,omitempty" parquet:"uid,optional" ch:"uid"`
	Version        string             `json:"version" parquet:"version" ch:"version"`
}
