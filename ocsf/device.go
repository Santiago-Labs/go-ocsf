package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// DeviceFields defines the Arrow fields for Device.
var DeviceFields = []arrow.Field{
	{Name: "region", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "interface_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "interface_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "os", Type: OSStruct, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hypervisor", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "location", Type: GeoLocationStruct, Nullable: true},
	{Name: "instance_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "first_seen_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "mac", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "org", Type: OrganizationStruct, Nullable: true},
	{Name: "risk_level", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "image", Type: ImageStruct, Nullable: true},
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "subnet_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "zone", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "groups", Type: arrow.ListOf(GroupStruct), Nullable: true},
	{Name: "risk_score", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "is_personal", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hw_info", Type: DeviceHWInfoStruct, Nullable: true},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "last_seen_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "is_managed", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "risk_level_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "is_trusted", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
	{Name: "network_interfaces", Type: arrow.ListOf(NetworkInterfaceStruct), Nullable: true},
	{Name: "autoscale_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vpc_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "subnet", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "imei", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ip", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hostname", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vlan_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "is_compliant", Type: arrow.FixedWidthTypes.Boolean, Nullable: true},
}

var DeviceStruct = arrow.StructOf(DeviceFields...)
var DeviceClassname = "device"

type Device struct {
	Region            *string             `json:"region,omitempty" parquet:"region,optional"`
	InterfaceName     *string             `json:"interface_name,omitempty" parquet:"interface_name,optional"`
	UID               *string             `json:"uid,omitempty" parquet:"uid,optional"`
	InterfaceUID      *string             `json:"interface_uid,omitempty" parquet:"interface_uid,optional"`
	ModifiedTime      *int64              `json:"modified_time,omitempty" parquet:"modified_time,optional"`
	OS                *OS                 `json:"os,omitempty" parquet:"os,optional"`
	Desc              *string             `json:"desc,omitempty" parquet:"desc,optional"`
	Hypervisor        *string             `json:"hypervisor,omitempty" parquet:"hypervisor,optional"`
	Type              *string             `json:"type,omitempty" parquet:"type,optional"`
	Location          *GeoLocation        `json:"location,omitempty" parquet:"location,optional"`
	InstanceUID       *string             `json:"instance_uid,omitempty" parquet:"instance_uid,optional"`
	FirstSeenTime     *int64              `json:"first_seen_time,omitempty" parquet:"first_seen_time,optional"`
	MAC               *string             `json:"mac,omitempty" parquet:"mac,optional"`
	Org               *Organization       `json:"org,omitempty" parquet:"org,optional"`
	RiskLevel         *string             `json:"risk_level,omitempty" parquet:"risk_level,optional"`
	Image             *Image              `json:"image,omitempty" parquet:"image,optional"`
	CreatedTime       *int64              `json:"created_time,omitempty" parquet:"created_time,optional"`
	SubnetUID         *string             `json:"subnet_uid,omitempty" parquet:"subnet_uid,optional"`
	Zone              *string             `json:"zone,omitempty" parquet:"zone,optional"`
	Groups            []*Group            `json:"groups,omitempty" parquet:"groups,list,optional"`
	RiskScore         *int                `json:"risk_score,omitempty" parquet:"risk_score,optional"`
	IsPersonal        *bool               `json:"is_personal,omitempty" parquet:"is_personal,optional"`
	Name              *string             `json:"name,omitempty" parquet:"name,optional"`
	HWInfo            *DeviceHWInfo       `json:"hw_info,omitempty" parquet:"hw_info,optional"`
	UIDAlt            *string             `json:"uid_alt,omitempty" parquet:"uid_alt,optional"`
	TypeID            int                 `json:"type_id" parquet:"type_id"`
	LastSeenTime      *int64              `json:"last_seen_time,omitempty" parquet:"last_seen_time,optional"`
	IsManaged         *bool               `json:"is_managed,omitempty" parquet:"is_managed,optional"`
	RiskLevelID       *int                `json:"risk_level_id,omitempty" parquet:"risk_level_id,optional"`
	IsTrusted         *bool               `json:"is_trusted,omitempty" parquet:"is_trusted,optional"`
	NetworkInterfaces []*NetworkInterface `json:"network_interfaces,omitempty" parquet:"network_interfaces,list,optional"`
	AutoscaleUID      *string             `json:"autoscale_uid,omitempty" parquet:"autoscale_uid,optional"`
	VpcUID            *string             `json:"vpc_uid,omitempty" parquet:"vpc_uid,optional"`
	Subnet            *string             `json:"subnet,omitempty" parquet:"subnet,optional"`
	Domain            *string             `json:"domain,omitempty" parquet:"domain,optional"`
	IMEI              *string             `json:"imei,omitempty" parquet:"imei,optional"`
	IP                *string             `json:"ip,omitempty" parquet:"ip,optional"`
	Hostname          *string             `json:"hostname,omitempty" parquet:"hostname,optional"`
	VlanUID           *string             `json:"vlan_uid,omitempty" parquet:"vlan_uid,optional"`
	IsCompliant       *bool               `json:"is_compliant,omitempty" parquet:"is_compliant,optional"`
}
