package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// Assume the following schemas are defined elsewhere:
// SchemaExtensionSchema, LoggerSchema, and ProductSchema.

// MetadataFields defines the Arrow fields for Metadata.
var MetadataFields = []arrow.Field{
	{Name: "correlation_uid", Type: arrow.BinaryTypes.String},
	{Name: "event_code", Type: arrow.BinaryTypes.String},
	{Name: "extension", Type: arrow.StructOf(SchemaExtensionFields...)},
	{Name: "extensions", Type: arrow.ListOf(arrow.StructOf(SchemaExtensionFields...))},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "log_level", Type: arrow.BinaryTypes.String},
	{Name: "log_name", Type: arrow.BinaryTypes.String},
	{Name: "log_provider", Type: arrow.BinaryTypes.String},
	{Name: "log_version", Type: arrow.BinaryTypes.String},
	{Name: "logged_time", Type: arrow.BinaryTypes.String},
	{Name: "loggers", Type: arrow.ListOf(arrow.StructOf(LoggerFields...))},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "original_time", Type: arrow.BinaryTypes.String},
	{Name: "processed_time", Type: arrow.BinaryTypes.String},
	{Name: "product", Type: arrow.StructOf(ProductFields...)},
	{Name: "profiles", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "sequence", Type: arrow.PrimitiveTypes.Int32},
	{Name: "tenant_uid", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// MetadataSchema is the Arrow schema for Metadata.
var MetadataSchema = arrow.NewSchema(MetadataFields, nil)

// Metadata represents the metadata field.
type Metadata struct {
	CorrelationUID *string           `json:"correlation_uid,omitempty"`
	EventCode      *string           `json:"event_code,omitempty"`
	Extension      *SchemaExtension  `json:"extension,omitempty"`
	Extensions     []SchemaExtension `json:"extensions,omitempty"`
	Labels         []string          `json:"labels,omitempty"`
	LogLevel       *string           `json:"log_level,omitempty"`
	LogName        *string           `json:"log_name,omitempty"`
	LogProvider    *string           `json:"log_provider,omitempty"`
	LogVersion     *string           `json:"log_version,omitempty"`
	LoggedTime     *time.Time        `json:"logged_time,omitempty"`
	Loggers        []Logger          `json:"loggers,omitempty"`
	ModifiedTime   *time.Time        `json:"modified_time,omitempty"`
	OriginalTime   *time.Time        `json:"original_time,omitempty"`
	ProcessedTime  *time.Time        `json:"processed_time,omitempty"`
	Product        Product           `json:"product"`
	Profiles       []string          `json:"profiles,omitempty"`
	Sequence       *int              `json:"sequence,omitempty"`
	TenantUID      *string           `json:"tenant_uid,omitempty"`
	UID            *string           `json:"uid,omitempty"`
	Version        string            `json:"version"`
}

// WriteToParquet writes the Metadata fields to the provided Arrow StructBuilder.
func (m *Metadata) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: correlation_uid.
	corrB := sb.FieldBuilder(0).(*array.StringBuilder)
	if m.CorrelationUID != nil {
		corrB.Append(*m.CorrelationUID)
	} else {
		corrB.AppendNull()
	}

	// Field 1: event_code.
	eventCodeB := sb.FieldBuilder(1).(*array.StringBuilder)
	if m.EventCode != nil {
		eventCodeB.Append(*m.EventCode)
	} else {
		eventCodeB.AppendNull()
	}

	// Field 2: extension (nested struct).
	extB := sb.FieldBuilder(2).(*array.StructBuilder)
	if m.Extension != nil {
		extB.Append(true)
		m.Extension.WriteToParquet(extB)
	} else {
		extB.AppendNull()
	}

	// Field 3: extensions (list of SchemaExtension).
	extListB := sb.FieldBuilder(3).(*array.ListBuilder)
	if len(m.Extensions) > 0 {
		extListB.Append(true)
		extValB := extListB.ValueBuilder().(*array.StructBuilder)
		for _, ext := range m.Extensions {
			extValB.Append(true)
			ext.WriteToParquet(extValB)
		}
	} else {
		extListB.AppendNull()
	}

	// Field 4: labels (list of strings).
	labelsB := sb.FieldBuilder(4).(*array.ListBuilder)
	if len(m.Labels) > 0 {
		labelsB.Append(true)
		labelsValB := labelsB.ValueBuilder().(*array.StringBuilder)
		for _, label := range m.Labels {
			labelsValB.Append(label)
		}
	} else {
		labelsB.AppendNull()
	}

	// Field 5: log_level.
	logLevelB := sb.FieldBuilder(5).(*array.StringBuilder)
	if m.LogLevel != nil {
		logLevelB.Append(*m.LogLevel)
	} else {
		logLevelB.AppendNull()
	}

	// Field 6: log_name.
	logNameB := sb.FieldBuilder(6).(*array.StringBuilder)
	if m.LogName != nil {
		logNameB.Append(*m.LogName)
	} else {
		logNameB.AppendNull()
	}

	// Field 7: log_provider.
	logProviderB := sb.FieldBuilder(7).(*array.StringBuilder)
	if m.LogProvider != nil {
		logProviderB.Append(*m.LogProvider)
	} else {
		logProviderB.AppendNull()
	}

	// Field 8: log_version.
	logVersionB := sb.FieldBuilder(8).(*array.StringBuilder)
	if m.LogVersion != nil {
		logVersionB.Append(*m.LogVersion)
	} else {
		logVersionB.AppendNull()
	}

	// Field 9: logged_time.
	loggedTimeB := sb.FieldBuilder(9).(*array.StringBuilder)
	if m.LoggedTime != nil {
		loggedTimeB.Append(m.LoggedTime.Format(time.RFC3339))
	} else {
		loggedTimeB.AppendNull()
	}

	// Field 10: loggers (list of Logger structs).
	loggersB := sb.FieldBuilder(10).(*array.ListBuilder)
	if len(m.Loggers) > 0 {
		loggersB.Append(true)
		loggersValB := loggersB.ValueBuilder().(*array.StructBuilder)
		for _, logger := range m.Loggers {
			loggersValB.Append(true)
			logger.WriteToParquet(loggersValB)
		}
	} else {
		loggersB.AppendNull()
	}

	// Field 11: modified_time.
	modTimeB := sb.FieldBuilder(11).(*array.StringBuilder)
	if m.ModifiedTime != nil {
		modTimeB.Append(m.ModifiedTime.Format(time.RFC3339))
	} else {
		modTimeB.AppendNull()
	}

	// Field 12: original_time.
	origTimeB := sb.FieldBuilder(12).(*array.StringBuilder)
	if m.OriginalTime != nil {
		origTimeB.Append(m.OriginalTime.Format(time.RFC3339))
	} else {
		origTimeB.AppendNull()
	}

	// Field 13: processed_time.
	procTimeB := sb.FieldBuilder(13).(*array.StringBuilder)
	if m.ProcessedTime != nil {
		procTimeB.Append(m.ProcessedTime.Format(time.RFC3339))
	} else {
		procTimeB.AppendNull()
	}

	// Field 14: product (nested struct).
	prodB := sb.FieldBuilder(14).(*array.StructBuilder)
	prodB.Append(true)
	m.Product.WriteToParquet(prodB) // Product is non-pointer.

	// Field 15: profiles (list of strings).
	profilesB := sb.FieldBuilder(15).(*array.ListBuilder)
	if len(m.Profiles) > 0 {
		profilesB.Append(true)
		profilesValB := profilesB.ValueBuilder().(*array.StringBuilder)
		for _, profile := range m.Profiles {
			profilesValB.Append(profile)
		}
	} else {
		profilesB.AppendNull()
	}

	// Field 16: sequence.
	seqB := sb.FieldBuilder(16).(*array.Int32Builder)
	if m.Sequence != nil {
		seqB.Append(int32(*m.Sequence))
	} else {
		seqB.AppendNull()
	}

	// Field 17: tenant_uid.
	tenantUIDB := sb.FieldBuilder(17).(*array.StringBuilder)
	if m.TenantUID != nil {
		tenantUIDB.Append(*m.TenantUID)
	} else {
		tenantUIDB.AppendNull()
	}

	// Field 18: uid.
	uidB := sb.FieldBuilder(18).(*array.StringBuilder)
	if m.UID != nil {
		uidB.Append(*m.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 19: version.
	versionB := sb.FieldBuilder(19).(*array.StringBuilder)
	versionB.Append(m.Version)
}
