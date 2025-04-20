package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// NetworkInterfaceFields defines the Arrow fields for NetworkInterface.
var NetworkInterfaceFields = []arrow.Field{
	{Name: "hostname", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ip", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "mac", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "namespace", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "subnet_prefix", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var NetworkInterfaceStruct = arrow.StructOf(NetworkInterfaceFields...)
var NetworkInterfaceClassname = "network_interface"

type NetworkInterface struct {
	Hostname     *string `json:"hostname,omitempty" parquet:"hostname,optional"`
	IP           *string `json:"ip,omitempty" parquet:"ip,optional"`
	MAC          *string `json:"mac,omitempty" parquet:"mac,optional"`
	Name         *string `json:"name,omitempty" parquet:"name,optional"`
	Namespace    *string `json:"namespace,omitempty" parquet:"namespace,optional"`
	SubnetPrefix *int    `json:"subnet_prefix,omitempty" parquet:"subnet_prefix,optional"`
	Type         *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID       int     `json:"type_id" parquet:"type_id"`
	UID          *string `json:"uid,omitempty" parquet:"uid,optional"`
}
