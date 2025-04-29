package ocsf

type DiscoveryDetails struct {
	Count             *int32             `json:"count,omitempty" parquet:"count,optional"`
	OccurrenceDetails *OccurrenceDetails `json:"occurrence_details,omitempty" parquet:"occurrence_details,optional"`
	Type              *string            `json:"type,omitempty" parquet:"type,optional"`
	Value             *string            `json:"value,omitempty" parquet:"value,optional"`
}
