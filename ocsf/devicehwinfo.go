package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var DeviceHWInfoFields = []arrow.Field{
	{Name: "bios_date", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "bios_manufacturer", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "bios_ver", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "chassis", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cpu_bits", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "cpu_cores", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "cpu_count", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "cpu_speed", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "cpu_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desktop_display", Type: DisplayStruct, Nullable: true},
	{Name: "keyboard_info", Type: KeyboardInfoStruct, Nullable: true},
	{Name: "ram_size", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "serial_number", Type: arrow.BinaryTypes.String, Nullable: true},
}

var DeviceHWInfoStruct = arrow.StructOf(DeviceHWInfoFields...)
var DeviceHWInfoClassname = "device_hw_info"

// DeviceHWInfo contains hardware information.
type DeviceHWInfo struct {
	BiosDate         *string       `json:"bios_date" parquet:"bios_date,optional" ch:"bios_date" ch:"bios_date"`
	BiosManufacturer *string       `json:"bios_manufacturer" parquet:"bios_manufacturer,optional" ch:"bios_manufacturer"`
	BiosVer          *string       `json:"bios_ver" parquet:"bios_ver,optional" ch:"bios_ver"`
	Chassis          *string       `json:"chassis" parquet:"chassis,optional" ch:"chassis"`
	CPUBits          *int64        `json:"cpu_bits" parquet:"cpu_bits,optional" ch:"cpu_bits"`
	CPUCores         *int64        `json:"cpu_cores" parquet:"cpu_cores,optional" ch:"cpu_cores"`
	CPUCount         *int64        `json:"cpu_count" parquet:"cpu_count,optional" ch:"cpu_count"`
	CPUSpeed         *int64        `json:"cpu_speed" parquet:"cpu_speed,optional" ch:"cpu_speed"`
	CPUType          *string       `json:"cpu_type" parquet:"cpu_type,optional" ch:"cpu_type"`
	DesktopDisplay   *Display      `json:"desktop_display" parquet:"desktop_display,optional" ch:"desktop_display"`
	KeyboardInfo     *KeyboardInfo `json:"keyboard_info" parquet:"keyboard_info,optional" ch:"keyboard_info"`
	RamSize          *int64        `json:"ram_size" parquet:"ram_size,optional" ch:"ram_size"`
	SerialNumber     *string       `json:"serial_number" parquet:"serial_number,optional" ch:"serial_number"`
}
