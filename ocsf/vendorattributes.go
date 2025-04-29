package ocsf

type VendorAttributes struct {
	Severity   *string `json:"severity" parquet:"severity,optional"`
	SeverityID *int32  `json:"severity_id" parquet:"severity_id,optional"`
}
