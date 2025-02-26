package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
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
	BiosDate         *string       `json:"bios_date,omitempty"`
	BiosManufacturer *string       `json:"bios_manufacturer,omitempty"`
	BiosVer          *string       `json:"bios_ver,omitempty"`
	Chassis          *string       `json:"chassis,omitempty"`
	CPUBits          *int          `json:"cpu_bits,omitempty"`
	CPUCores         *int          `json:"cpu_cores,omitempty"`
	CPUCount         *int          `json:"cpu_count,omitempty"`
	CPUSpeed         *int          `json:"cpu_speed,omitempty"`
	CPUType          *string       `json:"cpu_type,omitempty"`
	DesktopDisplay   *Display      `json:"desktop_display,omitempty"`
	KeyboardInfo     *KeyboardInfo `json:"keyboard_info,omitempty"`
	RamSize          *int          `json:"ram_size,omitempty"`
	SerialNumber     *string       `json:"serial_number,omitempty"`
}

// WriteToParquet writes the DeviceHWInfo fields to the provided Arrow StructBuilder.
func (d *DeviceHWInfo) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: BiosDate.
	biosDateB := sb.FieldBuilder(0).(*array.StringBuilder)
	if d.BiosDate != nil {
		biosDateB.Append(*d.BiosDate)
	} else {
		biosDateB.AppendNull()
	}

	// Field 1: BiosManufacturer.
	biosManB := sb.FieldBuilder(1).(*array.StringBuilder)
	if d.BiosManufacturer != nil {
		biosManB.Append(*d.BiosManufacturer)
	} else {
		biosManB.AppendNull()
	}

	// Field 2: BiosVer.
	biosVerB := sb.FieldBuilder(2).(*array.StringBuilder)
	if d.BiosVer != nil {
		biosVerB.Append(*d.BiosVer)
	} else {
		biosVerB.AppendNull()
	}

	// Field 3: Chassis.
	chassisB := sb.FieldBuilder(3).(*array.StringBuilder)
	if d.Chassis != nil {
		chassisB.Append(*d.Chassis)
	} else {
		chassisB.AppendNull()
	}

	// Field 4: CPUBits.
	cpuBitsB := sb.FieldBuilder(4).(*array.Int32Builder)
	if d.CPUBits != nil {
		cpuBitsB.Append(int32(*d.CPUBits))
	} else {
		cpuBitsB.AppendNull()
	}

	// Field 5: CPUCores.
	cpuCoresB := sb.FieldBuilder(5).(*array.Int32Builder)
	if d.CPUCores != nil {
		cpuCoresB.Append(int32(*d.CPUCores))
	} else {
		cpuCoresB.AppendNull()
	}

	// Field 6: CPUCount.
	cpuCountB := sb.FieldBuilder(6).(*array.Int32Builder)
	if d.CPUCount != nil {
		cpuCountB.Append(int32(*d.CPUCount))
	} else {
		cpuCountB.AppendNull()
	}

	// Field 7: CPUSpeed.
	cpuSpeedB := sb.FieldBuilder(7).(*array.Int32Builder)
	if d.CPUSpeed != nil {
		cpuSpeedB.Append(int32(*d.CPUSpeed))
	} else {
		cpuSpeedB.AppendNull()
	}

	// Field 8: CPUType.
	cpuTypeB := sb.FieldBuilder(8).(*array.StringBuilder)
	if d.CPUType != nil {
		cpuTypeB.Append(*d.CPUType)
	} else {
		cpuTypeB.AppendNull()
	}

	// Field 9: DesktopDisplay (nested struct).
	ddB := sb.FieldBuilder(9).(*array.StructBuilder)
	if d.DesktopDisplay != nil {
		ddB.Append(true)
		d.DesktopDisplay.WriteToParquet(ddB)
	} else {
		ddB.AppendNull()
	}

	// Field 10: KeyboardInfo (nested struct).
	kiB := sb.FieldBuilder(10).(*array.StructBuilder)
	if d.KeyboardInfo != nil {
		kiB.Append(true)
		d.KeyboardInfo.WriteToParquet(kiB)
	} else {
		kiB.AppendNull()
	}

	// Field 11: RamSize.
	ramSizeB := sb.FieldBuilder(11).(*array.Int32Builder)
	if d.RamSize != nil {
		ramSizeB.Append(int32(*d.RamSize))
	} else {
		ramSizeB.AppendNull()
	}

	// Field 12: SerialNumber.
	serialB := sb.FieldBuilder(12).(*array.StringBuilder)
	if d.SerialNumber != nil {
		serialB.Append(*d.SerialNumber)
	} else {
		serialB.AppendNull()
	}
}
