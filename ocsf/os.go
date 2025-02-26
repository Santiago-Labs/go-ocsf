package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// OSFields defines the Arrow fields for OS in the specified order.
var OSFields = []arrow.Field{
	{Name: "build", Type: arrow.BinaryTypes.String},
	{Name: "country", Type: arrow.BinaryTypes.String},
	{Name: "cpe_name", Type: arrow.BinaryTypes.String},
	{Name: "cpu_bits", Type: arrow.PrimitiveTypes.Int32},
	{Name: "edition", Type: arrow.BinaryTypes.String},
	{Name: "lang", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "sp_name", Type: arrow.BinaryTypes.String},
	{Name: "sp_ver", Type: arrow.PrimitiveTypes.Int32},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// OSSchema is the Arrow schema for OS.
var OSSchema = arrow.NewSchema(OSFields, nil)

// OS represents operating system information.
type OS struct {
	Build   *string `json:"build,omitempty"`
	Country *string `json:"country,omitempty"`
	CpeName *string `json:"cpe_name,omitempty"`
	CPUBits *int    `json:"cpu_bits,omitempty"`
	Edition *string `json:"edition,omitempty"`
	Lang    *string `json:"lang,omitempty"`
	Name    string  `json:"name"` // required field
	SPName  *string `json:"sp_name,omitempty"`
	SPVer   *int    `json:"sp_ver,omitempty"`
	Type    *string `json:"type,omitempty"`
	TypeID  int     `json:"type_id"` // required field
	Version *string `json:"version,omitempty"`
}

// WriteToParquet writes the OS fields to the provided Arrow StructBuilder.
func (os *OS) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: build.
	buildB := sb.FieldBuilder(0).(*array.StringBuilder)
	if os.Build != nil {
		buildB.Append(*os.Build)
	} else {
		buildB.AppendNull()
	}

	// Field 1: country.
	countryB := sb.FieldBuilder(1).(*array.StringBuilder)
	if os.Country != nil {
		countryB.Append(*os.Country)
	} else {
		countryB.AppendNull()
	}

	// Field 2: cpe_name.
	cpeNameB := sb.FieldBuilder(2).(*array.StringBuilder)
	if os.CpeName != nil {
		cpeNameB.Append(*os.CpeName)
	} else {
		cpeNameB.AppendNull()
	}

	// Field 3: cpu_bits.
	cpuBitsB := sb.FieldBuilder(3).(*array.Int32Builder)
	if os.CPUBits != nil {
		cpuBitsB.Append(int32(*os.CPUBits))
	} else {
		cpuBitsB.AppendNull()
	}

	// Field 4: edition.
	editionB := sb.FieldBuilder(4).(*array.StringBuilder)
	if os.Edition != nil {
		editionB.Append(*os.Edition)
	} else {
		editionB.AppendNull()
	}

	// Field 5: lang.
	langB := sb.FieldBuilder(5).(*array.StringBuilder)
	if os.Lang != nil {
		langB.Append(*os.Lang)
	} else {
		langB.AppendNull()
	}

	// Field 6: name (required).
	nameB := sb.FieldBuilder(6).(*array.StringBuilder)
	nameB.Append(os.Name)

	// Field 7: sp_name.
	spNameB := sb.FieldBuilder(7).(*array.StringBuilder)
	if os.SPName != nil {
		spNameB.Append(*os.SPName)
	} else {
		spNameB.AppendNull()
	}

	// Field 8: sp_ver.
	spVerB := sb.FieldBuilder(8).(*array.Int32Builder)
	if os.SPVer != nil {
		spVerB.Append(int32(*os.SPVer))
	} else {
		spVerB.AppendNull()
	}

	// Field 9: type.
	typeB := sb.FieldBuilder(9).(*array.StringBuilder)
	if os.Type != nil {
		typeB.Append(*os.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 10: type_id (required).
	typeIDB := sb.FieldBuilder(10).(*array.Int32Builder)
	typeIDB.Append(int32(os.TypeID))

	// Field 11: version.
	versionB := sb.FieldBuilder(11).(*array.StringBuilder)
	if os.Version != nil {
		versionB.Append(*os.Version)
	} else {
		versionB.AppendNull()
	}
}
