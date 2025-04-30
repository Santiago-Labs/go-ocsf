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
	{Name: "original_time", Type: arrow.BinaryTypes.String, Nullable: true},
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
	CorrelationUID *string            `json:"correlation_uid,omitempty" parquet:"correlation_uid,optional"`
	EventCode      *string            `json:"event_code,omitempty" parquet:"event_code,optional"`
	Extension      *SchemaExtension   `json:"extension,omitempty" parquet:"extension,optional"`
	Extensions     []*SchemaExtension `json:"extensions,omitempty" parquet:"extensions,list,optional"`
	Labels         []string           `json:"labels,omitempty" parquet:"labels,list,optional"`
	LogLevel       *string            `json:"log_level,omitempty" parquet:"log_level,optional"`
	LogName        *string            `json:"log_name,omitempty" parquet:"log_name,optional"`
	LogProvider    *string            `json:"log_provider,omitempty" parquet:"log_provider,optional"`
	LogVersion     *string            `json:"log_version,omitempty" parquet:"log_version,optional"`
	LoggedTime     *int64             `json:"logged_time,omitempty" parquet:"logged_time,optional"`
	Loggers        []*Logger          `json:"loggers,omitempty" parquet:"loggers,list,optional"`
	ModifiedTime   *int64             `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	OriginalTime   *string            `json:"original_time,omitempty" parquet:"original_time,optional"`
	ProcessedTime  *int64             `json:"processed_time,omitempty" parquet:"processed_time,optional"`
	Product        Product            `json:"product" parquet:"product"`
	Profiles       []string           `json:"profiles,omitempty" parquet:"profiles,list,optional"`
	Sequence       *int               `json:"sequence,omitempty" parquet:"sequence,optional"`
	TenantUID      *string            `json:"tenant_uid,omitempty" parquet:"tenant_uid,optional"`
	UID            *string            `json:"uid,omitempty" parquet:"uid,optional"`
	Version        string             `json:"version" parquet:"version"`
}
