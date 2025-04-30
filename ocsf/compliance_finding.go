package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var ComplianceFindingFields = []arrow.Field{
	{Name: "action", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "action_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "activity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "activity_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "category_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "category_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "class_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "class_uid", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "comment", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "compliance", Type: ComplianceStruct, Nullable: true},
	{Name: "confidence", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "confidence_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "confidence_score", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "count", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "disposition", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "disposition_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "duration", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "end_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "enrichments", Type: arrow.ListOf(EnrichmentStruct), Nullable: true},
	{Name: "evidences", Type: arrow.ListOf(EvidenceArtifactsStruct), Nullable: true},
	{Name: "finding_info", Type: FindingInfoStruct, Nullable: false},
	{Name: "message", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "metadata", Type: MetadataStruct, Nullable: false},
	{Name: "policy", Type: PolicyStruct, Nullable: true},
	{Name: "resources", Type: arrow.ListOf(ResourceDetailsStruct), Nullable: true},
	{Name: "risk_details", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "risk_level", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "risk_level_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "severity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "severity_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "observables", Type: arrow.ListOf(ObservableStruct), Nullable: true},
	{Name: "raw_data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "start_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "status", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_code", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_detail", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "status_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "time", Type: arrow.FixedWidthTypes.Timestamp_s, Nullable: false},
	{Name: "timezone_offset", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "type_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_uid", Type: arrow.PrimitiveTypes.Int64, Nullable: false},
	{Name: "vendor_attributes", Type: VendorAttributesStruct, Nullable: true},
}

var ComplianceFindingStruct = arrow.StructOf(ComplianceFindingFields...)
var ComplianceFindingSchema = arrow.NewSchema(ComplianceFindingFields, nil)

var ComplianceFindingClassname = "compliance_finding"

type ComplianceFinding struct {
	Action           *string                 `json:"action,omitempty" parquet:"action,optional"`
	ActionID         *int32                  `json:"action_id,omitempty" parquet:"action_id,optional"`
	ActivityID       int32                   `json:"activity_id" parquet:"activity_id"`
	ActivityName     *string                 `json:"activity_name,omitempty" parquet:"activity_name,optional"`
	CategoryName     *string                 `json:"category_name,omitempty" parquet:"category_name,optional"`
	CategoryUID      int32                   `json:"category_uid" parquet:"category_uid"`
	ClassName        *string                 `json:"class_name,omitempty" parquet:"class_name,optional"`
	ClassUID         int32                   `json:"class_uid" parquet:"class_uid"`
	Comment          *string                 `json:"comment,omitempty" parquet:"comment,optional"`
	Compliance       Compliance              `json:"compliance" parquet:"compliance"`
	Confidence       *string                 `json:"confidence,omitempty" parquet:"confidence,optional"`
	ConfidenceID     *int32                  `json:"confidence_id,omitempty" parquet:"confidence_id,optional"`
	ConfidenceScore  *int32                  `json:"confidence_score,omitempty" parquet:"confidence_score,optional"`
	Count            *int32                  `json:"count,omitempty" parquet:"count,optional"`
	Disposition      *string                 `json:"disposition,omitempty" parquet:"disposition,optional"`
	DispositionID    *int32                  `json:"disposition_id,omitempty" parquet:"disposition_id,optional"`
	Duration         *int64                  `json:"duration,omitempty" parquet:"duration,optional"`
	EndTime          *int64                  `json:"end_time,omitempty" parquet:"end_time,optional"`
	Enrichments      []*Enrichment           `json:"enrichments,omitempty" parquet:"enrichments,optional"`
	Evidences        []*EvidenceArtifacts    `json:"evidences,omitempty" parquet:"evidences,optional"`
	FindingInfo      FindingInfo             `json:"finding_info" parquet:"finding_info"`
	Message          *string                 `json:"message,omitempty" parquet:"message,optional"`
	Metadata         Metadata                `json:"metadata" parquet:"metadata"`
	Resources        []*ResourceDetails      `json:"resources,omitempty" parquet:"resources,optional"`
	Severity         *string                 `json:"severity,omitempty" parquet:"severity,optional"`
	SeverityID       int32                   `json:"severity_id" parquet:"severity_id"`
	Observables      []*Observable           `json:"observables,omitempty" parquet:"observables,optional"`
	RawData          *string                 `json:"raw_data,omitempty" parquet:"raw_data,optional"`
	StartTime        *int64                  `json:"start_time,omitempty" parquet:"start_time,optional"`
	Status           *string                 `json:"status,omitempty" parquet:"status,optional"`
	StatusCode       *string                 `json:"status_code,omitempty" parquet:"status_code,optional"`
	StatusDetail     *string                 `json:"status_detail,omitempty" parquet:"status_detail,optional"`
	StatusID         *int32                  `json:"status_id,omitempty" parquet:"status_id,optional"`
	TimezoneOffset   *int32                  `json:"timezone_offset,omitempty" parquet:"timezone_offset,optional"`
	TypeUID          int64                   `json:"type_uid" parquet:"type_uid"`
	TypeName         *string                 `json:"type_name,omitempty" parquet:"type_name,optional"`
	Unmapped         *string                 `json:"unmapped,omitempty" parquet:"unmapped,optional"`
	Vulnerabilities  []*VulnerabilityDetails `json:"vulnerabilities,omitempty" parquet:"vulnerabilities,optional"`
	VendorAttributes *VendorAttributes       `json:"vendor_attributes,omitempty" parquet:"vendor_attributes,optional"`
}
