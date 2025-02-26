package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
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
	{Name: "remediation", Type: arrow.StructOf(RemediationFields...)},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// AffectedSoftwarePackageSchema is the Arrow schema for AffectedSoftwarePackage.
var AffectedSoftwarePackageSchema = arrow.NewSchema(AffectedSoftwarePackageFields, nil)

type AffectedSoftwarePackage struct {
	Architecture   *string      `json:"architecture,omitempty"`
	Epoch          *int         `json:"epoch,omitempty"`
	FixedInVersion *string      `json:"fixed_in_version,omitempty"`
	License        *string      `json:"license,omitempty"`
	Name           string       `json:"name"`
	PackageManager *string      `json:"package_manager,omitempty"`
	Path           *string      `json:"path,omitempty"`
	Purl           *string      `json:"purl,omitempty"`
	Release        *string      `json:"release,omitempty"`
	Remediation    *Remediation `json:"remediation,omitempty"`
	Version        string       `json:"version"`
}

func (asp *AffectedSoftwarePackage) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: Architecture.
	archB := sb.FieldBuilder(0).(*array.StringBuilder)
	if asp.Architecture != nil {
		archB.Append(*asp.Architecture)
	} else {
		archB.AppendNull()
	}

	// Field 1: Epoch.
	epochB := sb.FieldBuilder(1).(*array.Int32Builder)
	if asp.Epoch != nil {
		epochB.Append(int32(*asp.Epoch))
	} else {
		epochB.AppendNull()
	}

	// Field 2: FixedInVersion.
	fixedInB := sb.FieldBuilder(2).(*array.StringBuilder)
	if asp.FixedInVersion != nil {
		fixedInB.Append(*asp.FixedInVersion)
	} else {
		fixedInB.AppendNull()
	}

	// Field 3: License.
	licenseB := sb.FieldBuilder(3).(*array.StringBuilder)
	if asp.License != nil {
		licenseB.Append(*asp.License)
	} else {
		licenseB.AppendNull()
	}

	// Field 4: Name.
	nameB := sb.FieldBuilder(4).(*array.StringBuilder)
	nameB.Append(asp.Name)

	// Field 5: PackageManager.
	pkgMgrB := sb.FieldBuilder(5).(*array.StringBuilder)
	if asp.PackageManager != nil {
		pkgMgrB.Append(*asp.PackageManager)
	} else {
		pkgMgrB.AppendNull()
	}

	// Field 6: Path.
	pathB := sb.FieldBuilder(6).(*array.StringBuilder)
	if asp.Path != nil {
		pathB.Append(*asp.Path)
	} else {
		pathB.AppendNull()
	}

	// Field 7: Purl.
	purlB := sb.FieldBuilder(7).(*array.StringBuilder)
	if asp.Purl != nil {
		purlB.Append(*asp.Purl)
	} else {
		purlB.AppendNull()
	}

	// Field 8: Release.
	releaseB := sb.FieldBuilder(8).(*array.StringBuilder)
	if asp.Release != nil {
		releaseB.Append(*asp.Release)
	} else {
		releaseB.AppendNull()
	}

	// Field 9: Remediation (nested).
	remediationB := sb.FieldBuilder(9).(*array.StructBuilder)
	if asp.Remediation != nil {
		remediationB.Append(true)
		asp.Remediation.WriteToParquet(remediationB)
	} else {
		remediationB.AppendNull()
	}

	// Field 10: Version.
	versionB := sb.FieldBuilder(10).(*array.StringBuilder)
	versionB.Append(asp.Version)
}
