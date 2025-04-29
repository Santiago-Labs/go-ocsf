package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var LoggerFields = []arrow.Field{
	{Name: "device", Type: arrow.StructOf(DeviceFields...), Nullable: true},
	{Name: "log_level", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "logged_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "transmit_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var LoggerStruct = arrow.StructOf(LoggerFields...)
var LoggerClassname = "logger"

type Logger struct {
	Device       *Device  `json:"device,omitempty" parquet:"device,optional"`
	EventUID     *string  `json:"event_uid,omitempty" parquet:"event_uid,optional"`
	LogLevel     *string  `json:"log_level,omitempty" parquet:"log_level,optional"`
	LogName      *string  `json:"log_name,omitempty" parquet:"log_name,optional"`
	LogProvider  *string  `json:"log_provider,omitempty" parquet:"log_provider,optional"`
	LogVersion   *string  `json:"log_version,omitempty" parquet:"log_version,optional"`
	LoggedTime   *int64   `json:"logged_time,omitempty" parquet:"logged_time,optional"`
	Name         *string  `json:"name,omitempty" parquet:"name,optional"`
	Product      *Product `json:"product,omitempty" parquet:"product,optional"`
	TransmitTime *int64   `json:"transmit_time,omitempty" parquet:"transmit_time,optional"`
	UID          *string  `json:"uid,omitempty" parquet:"uid,optional"`
	Version      *string  `json:"version,omitempty" parquet:"version,optional"`
}
