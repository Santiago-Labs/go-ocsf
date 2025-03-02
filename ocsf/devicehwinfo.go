package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

var DeviceHWInfoFields = []arrow.Field{
	{Name: "bios_date", Type: arrow.BinaryTypes.String},
	{Name: "bios_manufacturer", Type: arrow.BinaryTypes.String},
	{Name: "bios_ver", Type: arrow.BinaryTypes.String},
	{Name: "chassis", Type: arrow.BinaryTypes.String},
	{Name: "cpu_bits", Type: arrow.PrimitiveTypes.Int32},
	{Name: "cpu_cores", Type: arrow.PrimitiveTypes.Int32},
	{Name: "cpu_count", Type: arrow.PrimitiveTypes.Int32},
	{Name: "cpu_speed", Type: arrow.PrimitiveTypes.Int32},
	{Name: "cpu_type", Type: arrow.BinaryTypes.String},
	{Name: "desktop_display", Type: arrow.StructOf(DisplayFields...)},
	{Name: "keyboard_info", Type: arrow.StructOf(KeyboardInfoFields...)},
	{Name: "ram_size", Type: arrow.PrimitiveTypes.Int32},
	{Name: "serial_number", Type: arrow.BinaryTypes.String},
}

// DeviceHWInfo contains hardware information.
type DeviceHWInfo struct {
	BiosDate         *string       `json:"bios_date,omitempty" parquet:"bios_date"`
	BiosManufacturer *string       `json:"bios_manufacturer,omitempty" parquet:"bios_manufacturer"`
	BiosVer          *string       `json:"bios_ver,omitempty" parquet:"bios_ver"`
	Chassis          *string       `json:"chassis,omitempty" parquet:"chassis"`
	CPUBits          *int          `json:"cpu_bits,omitempty" parquet:"cpu_bits"`
	CPUCores         *int          `json:"cpu_cores,omitempty" parquet:"cpu_cores"`
	CPUCount         *int          `json:"cpu_count,omitempty" parquet:"cpu_count"`
	CPUSpeed         *int          `json:"cpu_speed,omitempty" parquet:"cpu_speed"`
	CPUType          *string       `json:"cpu_type,omitempty" parquet:"cpu_type"`
	DesktopDisplay   *Display      `json:"desktop_display,omitempty" parquet:"desktop_display"`
	KeyboardInfo     *KeyboardInfo `json:"keyboard_info,omitempty" parquet:"keyboard_info"`
	RamSize          *int          `json:"ram_size,omitempty" parquet:"ram_size"`
	SerialNumber     *string       `json:"serial_number,omitempty" parquet:"serial_number"`
}
