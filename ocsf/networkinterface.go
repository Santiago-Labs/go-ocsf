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
	Hostname     *string `json:"hostname" parquet:"hostname,optional" ch:"hostname"`
	IP           *string `json:"ip" parquet:"ip,optional" ch:"ip"`
	MAC          *string `json:"mac" parquet:"mac,optional" ch:"mac"`
	Name         *string `json:"name" parquet:"name,optional" ch:"name"`
	Namespace    *string `json:"namespace" parquet:"namespace,optional" ch:"namespace"`
	SubnetPrefix *int64  `json:"subnet_prefix" parquet:"subnet_prefix,optional" ch:"subnet_prefix"`
	Type         *string `json:"type" parquet:"type,optional" ch:"type"`
	TypeID       int     `json:"type_id" parquet:"type_id" ch:"type_id"`
	UID          *string `json:"uid" parquet:"uid,optional" ch:"uid"`
}
