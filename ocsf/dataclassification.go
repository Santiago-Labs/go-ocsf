package ocsf

type DataClassification struct {
	Category          *string            `json:"category,omitempty" parquet:"category,optional"`
	CategoryID        *int32             `json:"category_id,omitempty" parquet:"category_id,optional"`
	ClassifierDetails *ClassifierDetails `json:"classifier_details,omitempty" parquet:"classifier_details,optional"`
	Confidentiality   *string            `json:"confidentiality,omitempty" parquet:"confidentiality,optional"`
	ConfidentialityID *int32             `json:"confidentiality_id,omitempty" parquet:"confidentiality_id,optional"`
	DiscoveryDetails  *DiscoveryDetails  `json:"discovery_details,omitempty" parquet:"discovery_details,optional"`
	Policy            *Policy            `json:"policy,omitempty" parquet:"policy,optional"`
	Size              *int64             `json:"size,omitempty" parquet:"size,optional"`
	SrcURL            *string            `json:"src_url,omitempty" parquet:"src_url,optional"`
	Status            *string            `json:"status,omitempty" parquet:"status,optional"`
	StatusDetails     *string            `json:"status_details,omitempty" parquet:"status_details,optional"`
	StatusID          *int32             `json:"status_id,omitempty" parquet:"status_id,optional"`
	Total             *int32             `json:"total,omitempty" parquet:"total,optional"`
	UID               *string            `json:"uid,omitempty" parquet:"uid,optional"`
}
