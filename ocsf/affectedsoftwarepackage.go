package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// AffectedSoftwarePackageFields defines the Arrow fields for AffectedSoftwarePackage.
var AffectedSoftwarePackageFields = []arrow.Field{
	{Name: "architecture", Type: arrow.BinaryTypes.String},
	{Name: "epoch", Type: arrow.PrimitiveTypes.Int32},
	{Name: "fixed_in_version", Type: arrow.BinaryTypes.String},
	{Name: "license", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "package_manager", Type: arrow.BinaryTypes.String},
	{Name: "path", Type: arrow.BinaryTypes.String},
	{Name: "purl", Type: arrow.BinaryTypes.String},
	{Name: "release", Type: arrow.BinaryTypes.String},
	{Name: "remediation", Type: RemediationStruct},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

var AffectedSoftwarePackageStruct = arrow.StructOf(AffectedSoftwarePackageFields...)
var AffectedSoftwarePackageClassname = "affected_software_package"

type AffectedSoftwarePackage struct {
	Architecture   *string      `json:"architecture,omitempty" parquet:"architecture"`
	Epoch          *int32       `json:"epoch,omitempty" parquet:"epoch"`
	FixedInVersion *string      `json:"fixed_in_version,omitempty" parquet:"fixed_in_version"`
	License        *string      `json:"license,omitempty" parquet:"license"`
	Name           string       `json:"name" parquet:"name"`
	PackageManager *string      `json:"package_manager,omitempty" parquet:"package_manager"`
	Path           *string      `json:"path,omitempty" parquet:"path"`
	Purl           *string      `json:"purl,omitempty" parquet:"purl"`
	Release        *string      `json:"release,omitempty" parquet:"release"`
	Remediation    *Remediation `json:"remediation,omitempty" parquet:"remediation"`
	Version        string       `json:"version" parquet:"version"`
}
