package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// Assume the following schemas are defined elsewhere:
// SchemaExtensionSchema, LoggerSchema, and ProductSchema.

// MetadataFields defines the Arrow fields for Metadata.
var MetadataFields = []arrow.Field{
	{Name: "correlation_uid", Type: arrow.BinaryTypes.String},
	{Name: "event_code", Type: arrow.BinaryTypes.String},
	{Name: "extension", Type: SchemaExtensionStruct},
	{Name: "extensions", Type: arrow.ListOf(SchemaExtensionStruct)},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "log_level", Type: arrow.BinaryTypes.String},
	{Name: "log_name", Type: arrow.BinaryTypes.String},
	{Name: "log_provider", Type: arrow.BinaryTypes.String},
	{Name: "log_version", Type: arrow.BinaryTypes.String},
	{Name: "logged_time", Type: arrow.BinaryTypes.String},
	{Name: "loggers", Type: arrow.ListOf(LoggerStruct)},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "original_time", Type: arrow.BinaryTypes.String},
	{Name: "processed_time", Type: arrow.BinaryTypes.String},
	{Name: "product", Type: ProductStruct},
	{Name: "profiles", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "sequence", Type: arrow.PrimitiveTypes.Int32},
	{Name: "tenant_uid", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var MetadataStruct = arrow.StructOf(MetadataFields...)
var MetadataClassname = "metadata"

type Metadata struct {
	CorrelationUID *string           `json:"correlation_uid,omitempty" parquet:"correlation_uid"`
	EventCode      *string           `json:"event_code,omitempty" parquet:"event_code"`
	Extension      *SchemaExtension  `json:"extension,omitempty" parquet:"extension"`
	Extensions     []SchemaExtension `json:"extensions,omitempty" parquet:"extensions"`
	Labels         []string          `json:"labels,omitempty" parquet:"labels"`
	LogLevel       *string           `json:"log_level,omitempty" parquet:"log_level"`
	LogName        *string           `json:"log_name,omitempty" parquet:"log_name"`
	LogProvider    *string           `json:"log_provider,omitempty" parquet:"log_provider"`
	LogVersion     *string           `json:"log_version,omitempty" parquet:"log_version"`
	LoggedTime     *time.Time        `json:"logged_time,omitempty" parquet:"logged_time"`
	Loggers        []Logger          `json:"loggers,omitempty" parquet:"loggers"`
	ModifiedTime   *time.Time        `json:"modified_time,omitempty" parquet:"modified_time"`
	OriginalTime   *time.Time        `json:"original_time,omitempty" parquet:"original_time"`
	ProcessedTime  *time.Time        `json:"processed_time,omitempty" parquet:"processed_time"`
	Product        Product           `json:"product" parquet:"product"`
	Profiles       []string          `json:"profiles,omitempty" parquet:"profiles"`
	Sequence       *int              `json:"sequence,omitempty" parquet:"sequence"`
	TenantUID      *string           `json:"tenant_uid,omitempty" parquet:"tenant_uid"`
	UID            *string           `json:"uid,omitempty" parquet:"uid"`
	Version        string            `json:"version" parquet:"version"`
}
