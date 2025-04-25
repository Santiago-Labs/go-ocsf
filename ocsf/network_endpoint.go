package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// NetworkEndpointFields defines the Arrow fields for NetworkEndpoint.
var NetworkEndpointFields = []arrow.Field{
	{Name: "domain", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "hostname", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "instance_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "interface_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "interface_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "ip", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "mac", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "port", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "subnet_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "svc_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vlan_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "vpc_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "zone", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "agent_list", Type: arrow.ListOf(AgentStruct), Nullable: true},
	{Name: "autonomous_system", Type: AutonomousSystemStruct, Nullable: true},
	{Name: "hw_info", Type: DeviceHWInfoStruct, Nullable: true},
	{Name: "intermediate_ips", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "location", Type: GeoLocationStruct, Nullable: true},
	{Name: "os", Type: OSStruct, Nullable: true},
	{Name: "owner", Type: UserStruct, Nullable: true},
	{Name: "proxy_endpoint", Type: NetworkProxyEndpointStruct, Nullable: true},
}

var NetworkEndpointStruct = arrow.StructOf(NetworkEndpointFields...)
var NetworkEndpointClassname = "network_endpoint"

type NetworkEndpoint struct {
	Domain           *string               `json:"domain" parquet:"domain,optional" ch:"domain"`
	Hostname         *string               `json:"hostname" parquet:"hostname,optional" ch:"hostname"`
	InstanceUID      *string               `json:"instance_uid" parquet:"instance_uid,optional" ch:"instance_uid"`
	InterfaceName    *string               `json:"interface_name" parquet:"interface_name,optional" ch:"interface_name"`
	InterfaceUID     *string               `json:"interface_uid" parquet:"interface_uid,optional" ch:"interface_uid"`
	IP               *string               `json:"ip" parquet:"ip,optional" ch:"ip"`
	MAC              *string               `json:"mac" parquet:"mac,optional" ch:"mac"`
	Name             *string               `json:"name" parquet:"name,optional" ch:"name"`
	Port             *int64                `json:"port" parquet:"port,optional" ch:"port"`
	SubnetUID        *string               `json:"subnet_uid" parquet:"subnet_uid,optional" ch:"subnet_uid"`
	SvcName          *string               `json:"svc_name" parquet:"svc_name,optional" ch:"svc_name"`
	Type             *string               `json:"type" parquet:"type,optional" ch:"type"`
	TypeID           *int64                `json:"type_id" parquet:"type_id,optional" ch:"type_id"`
	UID              *string               `json:"uid" parquet:"uid" ch:"uid"`
	VLANUID          *string               `json:"vlan_uid" parquet:"vlan_uid" ch:"vlan_uid"`
	VPCUID           *string               `json:"vpc_uid" parquet:"vpc_uid,optional" ch:"vpc_uid"`
	Zone             *string               `json:"zone" parquet:"zone,optional" ch:"zone"`
	AgentList        []*Agent              `json:"agent_list" parquet:"agent_list,list,optional" ch:"agent_list"`
	AutonomousSystem *AutonomousSystem     `json:"autonomous_system" parquet:"autonomous_system,optional" ch:"autonomous_system"`
	HWInfo           *DeviceHWInfo         `json:"hw_info" parquet:"hw_info,optional" ch:"hw_info"`
	IntermediateIPs  []string              `json:"intermediate_ips" parquet:"intermediate_ips,list,optional" ch:"intermediate_ips"`
	Location         *GeoLocation          `json:"location" parquet:"location,optional" ch:"location"`
	OS               *OS                   `json:"os" parquet:"os,optional" ch:"os"`
	Owner            *User                 `json:"owner" parquet:"owner,optional" ch:"owner"`
	ProxyEndpoint    *NetworkProxyEndpoint `json:"proxy_endpoint" parquet:"proxy_endpoint,optional" ch:"proxy_endpoint"`
}
