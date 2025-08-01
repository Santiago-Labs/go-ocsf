// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_4_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type SoftwareBillofMaterials struct {

	// Created Time: The time when the SBOM was created.
	CreatedTime int64 `json:"created_time,omitempty" parquet:"created_time,timestamp_millis,timestamp(millisecond),optional"`

	// Software Package: The device software that is being discovered by an inventory process.
	Package SoftwarePackage `json:"package" parquet:"package"`

	// Product: The product that generated the SBOM e.g. cdxgen or Syft.
	Product *Product `json:"product,omitempty" parquet:"product,optional"`

	// Software Components: The list of software components used in the software package.
	SoftwareComponents []SoftwareComponent `json:"software_components" parquet:"software_components,list"`
}

var SoftwareBillofMaterialsFields = []arrow.Field{
	{Name: "created_time", Type: arrow.FixedWidthTypes.Timestamp_ms, Nullable: true},
	{Name: "package", Type: SoftwarePackageStruct, Nullable: false},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "software_components", Type: arrow.ListOf(SoftwareComponentStruct), Nullable: false},
}

var SoftwareBillofMaterialsStruct = arrow.StructOf(SoftwareBillofMaterialsFields...)

var SoftwareBillofMaterialsSchema = arrow.NewSchema(SoftwareBillofMaterialsFields, nil)
var SoftwareBillofMaterialsClassname = "sbom"
