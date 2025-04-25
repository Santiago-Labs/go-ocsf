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
	Region            *string             `json:"region" parquet:"region,optional" ch:"region"`
	InterfaceName     *string             `json:"interface_name" parquet:"interface_name,optional" ch:"interface_name"`
	UID               *string             `json:"uid" parquet:"uid,optional" ch:"uid"`
	InterfaceUID      *string             `json:"interface_uid" parquet:"interface_uid,optional" ch:"interface_uid"`
	ModifiedTime      *int64              `json:"modified_time" parquet:"modified_time,optional" ch:"modified_time"`
	OS                *OS                 `json:"os" parquet:"os,optional" ch:"os"`
	Desc              *string             `json:"desc" parquet:"desc,optional" ch:"desc"`
	Hypervisor        *string             `json:"hypervisor" parquet:"hypervisor,optional" ch:"hypervisor"`
	Type              *string             `json:"type" parquet:"type,optional" ch:"type"`
	Location          *GeoLocation        `json:"location" parquet:"location,optional" ch:"location"`
	InstanceUID       *string             `json:"instance_uid" parquet:"instance_uid,optional" ch:"instance_uid"`
	FirstSeenTime     *int64              `json:"first_seen_time" parquet:"first_seen_time,optional" ch:"first_seen_time"`
	MAC               *string             `json:"mac" parquet:"mac,optional" ch:"mac"`
	Org               *Organization       `json:"org" parquet:"org,optional" ch:"org"`
	RiskLevel         *string             `json:"risk_level" parquet:"risk_level,optional" ch:"risk_level"`
	Image             *Image              `json:"image" parquet:"image,optional" ch:"image"`
	CreatedTime       *int64              `json:"created_time" parquet:"created_time,optional" ch:"created_time"`
	SubnetUID         *string             `json:"subnet_uid" parquet:"subnet_uid,optional" ch:"subnet_uid"`
	Zone              *string             `json:"zone" parquet:"zone,optional" ch:"zone"`
	Groups            []*Group            `json:"groups" parquet:"groups,list,optional" ch:"groups"`
	RiskScore         *int64              `json:"risk_score" parquet:"risk_score,optional" ch:"risk_score"`
	IsPersonal        *bool               `json:"is_personal" parquet:"is_personal,optional" ch:"is_personal"`
	Name              *string             `json:"name" parquet:"name,optional" ch:"name"`
	HWInfo            *DeviceHWInfo       `json:"hw_info" parquet:"hw_info,optional" ch:"hw_info"`
	UIDAlt            *string             `json:"uid_alt" parquet:"uid_alt,optional" ch:"uid_alt"`
	TypeID            int                 `json:"type_id" parquet:"type_id" ch:"type_id"`
	LastSeenTime      *int64              `json:"last_seen_time" parquet:"last_seen_time,optional" ch:"last_seen_time"`
	IsManaged         *bool               `json:"is_managed" parquet:"is_managed,optional" ch:"is_managed"`
	RiskLevelID       *int64              `json:"risk_level_id" parquet:"risk_level_id,optional" ch:"risk_level_id"`
	IsTrusted         *bool               `json:"is_trusted" parquet:"is_trusted,optional" ch:"is_trusted"`
	NetworkInterfaces []*NetworkInterface `json:"network_interfaces" parquet:"network_interfaces,list,optional" ch:"network_interfaces"`
	AutoscaleUID      *string             `json:"autoscale_uid" parquet:"autoscale_uid,optional" ch:"autoscale_uid"`
	VpcUID            *string             `json:"vpc_uid" parquet:"vpc_uid,optional" ch:"vpc_uid"`
	Subnet            *string             `json:"subnet" parquet:"subnet,optional" ch:"subnet"`
	Domain            *string             `json:"domain" parquet:"domain,optional" ch:"domain"`
	IMEI              *string             `json:"imei" parquet:"imei,optional" ch:"imei"`
	IP                *string             `json:"ip" parquet:"ip,optional" ch:"ip"`
	Hostname          *string             `json:"hostname" parquet:"hostname,optional" ch:"hostname"`
	VlanUID           *string             `json:"vlan_uid" parquet:"vlan_uid,optional" ch:"vlan_uid"`
	IsCompliant       *bool               `json:"is_compliant" parquet:"is_compliant,optional" ch:"is_compliant"`
}
