package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// AffectedSoftwarePackageFields defines the Arrow fields for AffectedSoftwarePackage.
var AffectedSoftwarePackageFields = []arrow.Field{
	{Name: "architecture", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "epoch", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "fixed_in_version", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "license", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "package_manager", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "path", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "purl", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "release", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "remediation", Type: RemediationStruct, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: false},
}

var AffectedSoftwarePackageStruct = arrow.StructOf(AffectedSoftwarePackageFields...)
var AffectedSoftwarePackageClassname = "affected_software_package"

type AffectedSoftwarePackage struct {
	Architecture   *string      `json:"architecture,omitempty" parquet:"architecture,optional" ch:"architecture"`
	Epoch          *int32       `json:"epoch,omitempty" parquet:"epoch,optional" ch:"epoch"`
	FixedInVersion *string      `json:"fixed_in_version,omitempty" parquet:"fixed_in_version,optional" ch:"fixed_in_version"`
	License        *string      `json:"license,omitempty" parquet:"license,optional" ch:"license"`
	Name           string       `json:"name" parquet:"name" ch:"name"`
	PackageManager *string      `json:"package_manager,omitempty" parquet:"package_manager,optional" ch:"package_manager"`
	Path           *string      `json:"path,omitempty" parquet:"path,optional" ch:"path"`
	Purl           *string      `json:"purl,omitempty" parquet:"purl,optional" ch:"purl"`
	Release        *string      `json:"release,omitempty" parquet:"release,optional" ch:"release"`
	Remediation    *Remediation `json:"remediation,omitempty" parquet:"remediation,optional" ch:"remediation"`
	Version        string       `json:"version" parquet:"version" ch:"version"`
}
