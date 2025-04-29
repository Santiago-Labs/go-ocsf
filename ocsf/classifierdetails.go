package ocsf

type ClassifierDetails struct {
	Name *string `json:"name,omitempty" parquet:"name,optional"`
	Type string  `json:"type" parquet:"type"`
	UID  *string `json:"uid,omitempty" parquet:"uid,optional"`
}
