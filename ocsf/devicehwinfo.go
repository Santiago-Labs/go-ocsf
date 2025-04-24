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
	BiosDate         *string       `json:"bios_date,omitempty" parquet:"bios_date,optional" ch:"bios_date,omitempty" ch:"bios_date,omitempty"`
	BiosManufacturer *string       `json:"bios_manufacturer,omitempty" parquet:"bios_manufacturer,optional" ch:"bios_manufacturer,omitempty"`
	BiosVer          *string       `json:"bios_ver,omitempty" parquet:"bios_ver,optional" ch:"bios_ver,omitempty"`
	Chassis          *string       `json:"chassis,omitempty" parquet:"chassis,optional" ch:"chassis,omitempty"`
	CPUBits          *int          `json:"cpu_bits,omitempty" parquet:"cpu_bits,optional" ch:"cpu_bits,omitempty"`
	CPUCores         *int          `json:"cpu_cores,omitempty" parquet:"cpu_cores,optional" ch:"cpu_cores,omitempty"`
	CPUCount         *int          `json:"cpu_count,omitempty" parquet:"cpu_count,optional" ch:"cpu_count,omitempty"`
	CPUSpeed         *int          `json:"cpu_speed,omitempty" parquet:"cpu_speed,optional" ch:"cpu_speed,omitempty"`
	CPUType          *string       `json:"cpu_type,omitempty" parquet:"cpu_type,optional" ch:"cpu_type,omitempty"`
	DesktopDisplay   *Display      `json:"desktop_display,omitempty" parquet:"desktop_display,optional" ch:"desktop_display,omitempty"`
	KeyboardInfo     *KeyboardInfo `json:"keyboard_info,omitempty" parquet:"keyboard_info,optional" ch:"keyboard_info,omitempty"`
	RamSize          *int          `json:"ram_size,omitempty" parquet:"ram_size,optional" ch:"ram_size,omitempty"`
	SerialNumber     *string       `json:"serial_number,omitempty" parquet:"serial_number,optional" ch:"serial_number,omitempty"`
}
