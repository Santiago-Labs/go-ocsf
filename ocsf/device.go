package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

// DeviceFields defines the Arrow fields for Device.
var DeviceFields = []arrow.Field{
	{Name: "region", Type: arrow.BinaryTypes.String},
	{Name: "interface_name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "interface_uid", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "os", Type: OSStruct},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "hypervisor", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "location", Type: GeoLocationStruct},
	{Name: "instance_uid", Type: arrow.BinaryTypes.String},
	{Name: "first_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "mac", Type: arrow.BinaryTypes.String},
	{Name: "org", Type: OrganizationStruct},
	{Name: "risk_level", Type: arrow.BinaryTypes.String},
	{Name: "image", Type: ImageStruct},
	{Name: "created_time", Type: arrow.BinaryTypes.String},
	{Name: "subnet_uid", Type: arrow.BinaryTypes.String},
	{Name: "zone", Type: arrow.BinaryTypes.String},
	{Name: "groups", Type: arrow.ListOf(GroupStruct)},
	{Name: "risk_score", Type: arrow.PrimitiveTypes.Int32},
	{Name: "is_personal", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "hw_info", Type: DeviceHWInfoStruct},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "last_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "is_managed", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "risk_level_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "is_trusted", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "network_interfaces", Type: arrow.ListOf(NetworkInterfaceStruct)},
	{Name: "autoscale_uid", Type: arrow.BinaryTypes.String},
	{Name: "vpc_uid", Type: arrow.BinaryTypes.String},
	{Name: "subnet", Type: arrow.BinaryTypes.String},
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "imei", Type: arrow.BinaryTypes.String},
	{Name: "ip", Type: arrow.BinaryTypes.String},
	{Name: "hostname", Type: arrow.BinaryTypes.String},
	{Name: "vlan_uid", Type: arrow.BinaryTypes.String},
	{Name: "is_compliant", Type: arrow.FixedWidthTypes.Boolean},
}

var DeviceStruct = arrow.StructOf(DeviceFields...)

// DeviceSchema is the Arrow schema for Device.
var DeviceSchema = arrow.NewSchema(DeviceFields, nil)

type Device struct {
	Region            *string            `json:"region,omitempty" parquet:"region"`
	InterfaceName     *string            `json:"interface_name,omitempty" parquet:"interface_name"`
	UID               *string            `json:"uid,omitempty" parquet:"uid"`
	InterfaceUID      *string            `json:"interface_uid,omitempty" parquet:"interface_uid"`
	ModifiedTime      *time.Time         `json:"modified_time,omitempty" parquet:"modified_time"`
	OS                *OS                `json:"os,omitempty" parquet:"os"`
	Desc              *string            `json:"desc,omitempty" parquet:"desc"`
	Hypervisor        *string            `json:"hypervisor,omitempty" parquet:"hypervisor"`
	Type              *string            `json:"type,omitempty" parquet:"type"`
	Location          *GeoLocation       `json:"location,omitempty" parquet:"location"`
	InstanceUID       *string            `json:"instance_uid,omitempty" parquet:"instance_uid"`
	FirstSeenTime     *time.Time         `json:"first_seen_time,omitempty" parquet:"first_seen_time"`
	MAC               *string            `json:"mac,omitempty" parquet:"mac"`
	Org               *Organization      `json:"org,omitempty" parquet:"org"`
	RiskLevel         *string            `json:"risk_level,omitempty" parquet:"risk_level"`
	Image             *Image             `json:"image,omitempty" parquet:"image"`
	CreatedTime       *time.Time         `json:"created_time,omitempty" parquet:"created_time"`
	SubnetUID         *string            `json:"subnet_uid,omitempty" parquet:"subnet_uid"`
	Zone              *string            `json:"zone,omitempty" parquet:"zone"`
	Groups            []Group            `json:"groups,omitempty" parquet:"groups"`
	RiskScore         *int               `json:"risk_score,omitempty" parquet:"risk_score"`
	IsPersonal        *bool              `json:"is_personal,omitempty" parquet:"is_personal"`
	Name              *string            `json:"name,omitempty" parquet:"name"`
	HWInfo            *DeviceHWInfo      `json:"hw_info,omitempty" parquet:"hw_info"`
	UIDAlt            *string            `json:"uid_alt,omitempty" parquet:"uid_alt"`
	TypeID            int                `json:"type_id" parquet:"type_id"`
	LastSeenTime      *time.Time         `json:"last_seen_time,omitempty" parquet:"last_seen_time"`
	IsManaged         *bool              `json:"is_managed,omitempty" parquet:"is_managed"`
	RiskLevelID       *int               `json:"risk_level_id,omitempty" parquet:"risk_level_id"`
	IsTrusted         *bool              `json:"is_trusted,omitempty" parquet:"is_trusted"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces,omitempty" parquet:"network_interfaces"`
	AutoscaleUID      *string            `json:"autoscale_uid,omitempty" parquet:"autoscale_uid"`
	VpcUID            *string            `json:"vpc_uid,omitempty" parquet:"vpc_uid"`
	Subnet            *string            `json:"subnet,omitempty" parquet:"subnet"`
	Domain            *string            `json:"domain,omitempty" parquet:"domain"`
	IMEI              *string            `json:"imei,omitempty" parquet:"imei"`
	IP                *string            `json:"ip,omitempty" parquet:"ip"`
	Hostname          *string            `json:"hostname,omitempty" parquet:"hostname"`
	VlanUID           *string            `json:"vlan_uid,omitempty" parquet:"vlan_uid"`
	IsCompliant       *bool              `json:"is_compliant,omitempty" parquet:"is_compliant"`
}
