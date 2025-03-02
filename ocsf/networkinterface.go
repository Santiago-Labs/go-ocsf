package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// NetworkInterfaceFields defines the Arrow fields for NetworkInterface.
var NetworkInterfaceFields = []arrow.Field{
	{Name: "hostname", Type: arrow.BinaryTypes.String},
	{Name: "ip", Type: arrow.BinaryTypes.String},
	{Name: "mac", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "namespace", Type: arrow.BinaryTypes.String},
	{Name: "subnet_prefix", Type: arrow.PrimitiveTypes.Int32},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

type NetworkInterface struct {
	Hostname     *string `json:"hostname,omitempty" parquet:"hostname"`
	IP           *string `json:"ip,omitempty" parquet:"ip"`
	MAC          *string `json:"mac,omitempty" parquet:"mac"`
	Name         *string `json:"name,omitempty" parquet:"name"`
	Namespace    *string `json:"namespace,omitempty" parquet:"namespace"`
	SubnetPrefix *int    `json:"subnet_prefix,omitempty" parquet:"subnet_prefix"`
	Type         *string `json:"type,omitempty" parquet:"type"`
	TypeID       int     `json:"type_id" parquet:"type_id"`
	UID          *string `json:"uid,omitempty" parquet:"uid"`
}
