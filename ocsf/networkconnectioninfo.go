package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var NetworkConnectionInfoFields = []arrow.Field{
	{Name: "boundary", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "boundary_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "community_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "direction", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "direction_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
	{Name: "flag_history", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_num", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "protocol_ver", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "protocol_ver_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "session", Type: SessionStruct, Nullable: true},
	{Name: "tcp_flags", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var NetworkConnectionInfoStruct = arrow.StructOf(NetworkConnectionInfoFields...)
var NetworkConnectionInfoClassname = "network_connection_info"

type NetworkConnectionInfo struct {
	Boundary      *string  `json:"boundary,omitempty" parquet:"boundary,optional"`
	BoundaryID    *int32   `json:"boundary_id,omitempty" parquet:"boundary_id,optional"`
	CommunityUID  *string  `json:"community_uid,omitempty" parquet:"community_uid,optional"`
	Direction     *string  `json:"direction,omitempty" parquet:"direction,optional"`
	DirectionID   *int32   `json:"direction_id,omitempty" parquet:"direction_id,optional"`
	FlagHistory   *string  `json:"flag_history,omitempty" parquet:"flag_history,optional"`
	ProtocolName  *string  `json:"protocol_name,omitempty" parquet:"protocol_name,optional"`
	ProtocolNum   *int32   `json:"protocol_num,omitempty" parquet:"protocol_num,optional"`
	ProtocolVer   *string  `json:"protocol_ver,omitempty" parquet:"protocol_ver,optional"`
	ProtocolVerID *int32   `json:"protocol_ver_id,omitempty" parquet:"protocol_ver_id,optional"`
	Session       *Session `json:"session,omitempty" parquet:"session,optional"`
	TCPFlags      *int32   `json:"tcp_flags,omitempty" parquet:"tcp_flags,optional"`
	UID           *string  `json:"uid,omitempty" parquet:"uid,optional"`
}
