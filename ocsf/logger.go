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
	Device       *Device  `json:"device" parquet:"device,optional" ch:"device"`
	LogLevel     *string  `json:"log_level" parquet:"log_level,optional" ch:"log_level"`
	LogName      *string  `json:"log_name" parquet:"log_name,optional" ch:"log_name"`
	LogProvider  *string  `json:"log_provider" parquet:"log_provider,optional" ch:"log_provider"`
	LogVersion   *string  `json:"log_version" parquet:"log_version,optional" ch:"log_version"`
	LoggedTime   *int64   `json:"logged_time" parquet:"logged_time,optional" ch:"logged_time"`
	Name         *string  `json:"name" parquet:"name,optional" ch:"name"`
	Product      *Product `json:"product" parquet:"product,optional" ch:"product"`
	TransmitTime *int64   `json:"transmit_time" parquet:"transmit_time,optional" ch:"transmit_time"`
	UID          *string  `json:"uid" parquet:"uid,optional" ch:"uid"`
	Version      *string  `json:"version" parquet:"version,optional" ch:"version"`
}
