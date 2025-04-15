package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

var LoggerFields = []arrow.Field{
	{Name: "device", Type: arrow.StructOf(DeviceFields...), Nullable: true},
	{Name: "log_level", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_provider", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "log_version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "logged_time", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "transmit_time", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var LoggerStruct = arrow.StructOf(LoggerFields...)
var LoggerClassname = "logger"

type Logger struct {
	Device       *Device    `json:"device,omitempty" parquet:"device"`
	LogLevel     *string    `json:"log_level,omitempty" parquet:"log_level"`
	LogName      *string    `json:"log_name,omitempty" parquet:"log_name"`
	LogProvider  *string    `json:"log_provider,omitempty" parquet:"log_provider"`
	LogVersion   *string    `json:"log_version,omitempty" parquet:"log_version"`
	LoggedTime   *time.Time `json:"logged_time,omitempty" parquet:"logged_time"`
	Name         *string    `json:"name,omitempty" parquet:"name"`
	Product      *Product   `json:"product,omitempty" parquet:"product"`
	TransmitTime *time.Time `json:"transmit_time,omitempty" parquet:"transmit_time"`
	UID          *string    `json:"uid,omitempty" parquet:"uid"`
	Version      *string    `json:"version,omitempty" parquet:"version"`
}
