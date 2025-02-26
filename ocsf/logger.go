package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var LoggerFields = []arrow.Field{
	{Name: "device", Type: arrow.StructOf(DeviceFields...)},
	{Name: "log_level", Type: arrow.BinaryTypes.String},
	{Name: "log_name", Type: arrow.BinaryTypes.String},
	{Name: "log_provider", Type: arrow.BinaryTypes.String},
	{Name: "log_version", Type: arrow.BinaryTypes.String},
	{Name: "logged_time", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "product", Type: arrow.StructOf(ProductFields...)},
	{Name: "transmit_time", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// LoggerSchema is the Arrow schema for Logger.
var LoggerSchema = arrow.NewSchema(LoggerFields, nil)

// Logger represents logger details.
type Logger struct {
	Device       *Device    `json:"device,omitempty"`
	LogLevel     *string    `json:"log_level,omitempty"`
	LogName      *string    `json:"log_name,omitempty"`
	LogProvider  *string    `json:"log_provider,omitempty"`
	LogVersion   *string    `json:"log_version,omitempty"`
	LoggedTime   *time.Time `json:"logged_time,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Product      *Product   `json:"product,omitempty"`
	TransmitTime *time.Time `json:"transmit_time,omitempty"`
	UID          *string    `json:"uid,omitempty"`
	Version      *string    `json:"version,omitempty"`
}

// WriteToParquet writes the Logger fields to the provided Arrow StructBuilder.
func (l *Logger) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Device (nested struct)
	deviceB := sb.FieldBuilder(0).(*array.StructBuilder)
	if l.Device != nil {
		deviceB.Append(true)
		l.Device.WriteToParquet(deviceB)
	} else {
		deviceB.AppendNull()
	}

	// Field 1: LogLevel.
	logLevelB := sb.FieldBuilder(1).(*array.StringBuilder)
	if l.LogLevel != nil {
		logLevelB.Append(*l.LogLevel)
	} else {
		logLevelB.AppendNull()
	}

	// Field 2: LogName.
	logNameB := sb.FieldBuilder(2).(*array.StringBuilder)
	if l.LogName != nil {
		logNameB.Append(*l.LogName)
	} else {
		logNameB.AppendNull()
	}

	// Field 3: LogProvider.
	logProviderB := sb.FieldBuilder(3).(*array.StringBuilder)
	if l.LogProvider != nil {
		logProviderB.Append(*l.LogProvider)
	} else {
		logProviderB.AppendNull()
	}

	// Field 4: LogVersion.
	logVersionB := sb.FieldBuilder(4).(*array.StringBuilder)
	if l.LogVersion != nil {
		logVersionB.Append(*l.LogVersion)
	} else {
		logVersionB.AppendNull()
	}

	// Field 5: LoggedTime (formatted as RFC3339).
	loggedTimeB := sb.FieldBuilder(5).(*array.StringBuilder)
	if l.LoggedTime != nil {
		loggedTimeB.Append(l.LoggedTime.Format(time.RFC3339))
	} else {
		loggedTimeB.AppendNull()
	}

	// Field 6: Name.
	nameB := sb.FieldBuilder(6).(*array.StringBuilder)
	if l.Name != nil {
		nameB.Append(*l.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 7: Product (nested struct).
	productB := sb.FieldBuilder(7).(*array.StructBuilder)
	if l.Product != nil {
		productB.Append(true)
		l.Product.WriteToParquet(productB)
	} else {
		productB.AppendNull()
	}

	// Field 8: TransmitTime.
	transmitTimeB := sb.FieldBuilder(8).(*array.StringBuilder)
	if l.TransmitTime != nil {
		transmitTimeB.Append(l.TransmitTime.Format(time.RFC3339))
	} else {
		transmitTimeB.AppendNull()
	}

	// Field 9: UID.
	uidB := sb.FieldBuilder(9).(*array.StringBuilder)
	if l.UID != nil {
		uidB.Append(*l.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 10: Version.
	versionB := sb.FieldBuilder(10).(*array.StringBuilder)
	if l.Version != nil {
		versionB.Append(*l.Version)
	} else {
		versionB.AppendNull()
	}
}
