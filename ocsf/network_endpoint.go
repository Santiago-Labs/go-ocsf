package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
)

// NetworkEndpointFields defines the Arrow fields for NetworkEndpoint.
var NetworkEndpointFields = []arrow.Field{
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "hostname", Type: arrow.BinaryTypes.String},
	{Name: "instance_uid", Type: arrow.BinaryTypes.String},
	{Name: "interface_name", Type: arrow.BinaryTypes.String},
	{Name: "interface_uid", Type: arrow.BinaryTypes.String},
	{Name: "ip", Type: arrow.BinaryTypes.String},
	{Name: "mac", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "port", Type: arrow.PrimitiveTypes.Int32},
	{Name: "subnet_uid", Type: arrow.BinaryTypes.String},
	{Name: "svc_name", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "vlan_uid", Type: arrow.BinaryTypes.String},
	{Name: "vpc_uid", Type: arrow.BinaryTypes.String},
	{Name: "zone", Type: arrow.BinaryTypes.String},
	{Name: "agent_list", Type: arrow.ListOf(AgentStruct)},
	{Name: "autonomous_system", Type: AutonomousSystemStruct},
	{Name: "hw_info", Type: DeviceHWInfoStruct},
	{Name: "intermediate_ips", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "location", Type: GeoLocationStruct},
	{Name: "os", Type: OSStruct},
	{Name: "owner", Type: UserStruct},
	{Name: "proxy_endpoint", Type: NetworkProxyEndpointStruct},
}

var NetworkEndpointStruct = arrow.StructOf(NetworkEndpointFields...)
var NetworkEndpointClassname = "network_endpoint"

type NetworkEndpoint struct {
	Domain           *string               `json:"domain,omitempty" parquet:"domain"`
	Hostname         *string               `json:"hostname,omitempty" parquet:"hostname"`
	InstanceUID      *string               `json:"instance_uid,omitempty" parquet:"instance_uid"`
	InterfaceName    *string               `json:"interface_name,omitempty" parquet:"interface_name"`
	InterfaceUID     *string               `json:"interface_uid,omitempty" parquet:"interface_uid"`
	IP               *string               `json:"ip,omitempty" parquet:"ip"`
	MAC              *string               `json:"mac,omitempty" parquet:"mac"`
	Name             *string               `json:"name,omitempty" parquet:"name"`
	Port             *int                  `json:"port,omitempty" parquet:"port"`
	SubnetUID        *string               `json:"subnet_uid,omitempty" parquet:"subnet_uid"`
	SvcName          *string               `json:"svc_name,omitempty" parquet:"svc_name"`
	Type             *string               `json:"type,omitempty" parquet:"type"`
	TypeID           *int                  `json:"type_id,omitempty" parquet:"type_id"`
	UID              *string               `json:"uid,omitempty" parquet:"uid"`
	VLANUID          *string               `json:"vlan_uid,omitempty" parquet:"vlan_uid"`
	VPCUID           *string               `json:"vpc_uid,omitempty" parquet:"vpc_uid"`
	Zone             *string               `json:"zone,omitempty" parquet:"zone"`
	AgentList        []*Agent              `json:"agent_list,omitempty" parquet:"agent_list"`
	AutonomousSystem *AutonomousSystem     `json:"autonomous_system,omitempty" parquet:"autonomous_system"`
	HWInfo           *DeviceHWInfo         `json:"hw_info,omitempty" parquet:"hw_info"`
	IntermediateIPs  []*string             `json:"intermediate_ips,omitempty" parquet:"intermediate_ips"`
	Location         *GeoLocation          `json:"location,omitempty" parquet:"location"`
	OS               *OS                   `json:"os,omitempty" parquet:"os"`
	Owner            *User                 `json:"owner,omitempty" parquet:"owner"`
	ProxyEndpoint    *NetworkProxyEndpoint `json:"proxy_endpoint,omitempty" parquet:"proxy_endpoint"`
}
