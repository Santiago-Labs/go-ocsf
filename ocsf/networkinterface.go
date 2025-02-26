package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
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

// NetworkInterfaceSchema is the Arrow schema for NetworkInterface.
var NetworkInterfaceSchema = arrow.NewSchema(NetworkInterfaceFields, nil)

// NetworkInterface represents a network interface.
type NetworkInterface struct {
	Hostname     *string `json:"hostname,omitempty"`
	IP           *string `json:"ip,omitempty"`
	MAC          *string `json:"mac,omitempty"`
	Name         *string `json:"name,omitempty"`
	Namespace    *string `json:"namespace,omitempty"`
	SubnetPrefix *int    `json:"subnet_prefix,omitempty"`
	Type         *string `json:"type,omitempty"`
	TypeID       int     `json:"type_id"`
	UID          *string `json:"uid,omitempty"`
}

// WriteToParquet writes the NetworkInterface fields to the provided Arrow StructBuilder.
func (ni *NetworkInterface) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Hostname.
	hostnameB := sb.FieldBuilder(0).(*array.StringBuilder)
	if ni.Hostname != nil {
		hostnameB.Append(*ni.Hostname)
	} else {
		hostnameB.AppendNull()
	}

	// Field 1: IP.
	ipB := sb.FieldBuilder(1).(*array.StringBuilder)
	if ni.IP != nil {
		ipB.Append(*ni.IP)
	} else {
		ipB.AppendNull()
	}

	// Field 2: MAC.
	macB := sb.FieldBuilder(2).(*array.StringBuilder)
	if ni.MAC != nil {
		macB.Append(*ni.MAC)
	} else {
		macB.AppendNull()
	}

	// Field 3: Name.
	nameB := sb.FieldBuilder(3).(*array.StringBuilder)
	if ni.Name != nil {
		nameB.Append(*ni.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 4: Namespace.
	namespaceB := sb.FieldBuilder(4).(*array.StringBuilder)
	if ni.Namespace != nil {
		namespaceB.Append(*ni.Namespace)
	} else {
		namespaceB.AppendNull()
	}

	// Field 5: SubnetPrefix.
	subnetPrefixB := sb.FieldBuilder(5).(*array.Int32Builder)
	if ni.SubnetPrefix != nil {
		subnetPrefixB.Append(int32(*ni.SubnetPrefix))
	} else {
		subnetPrefixB.AppendNull()
	}

	// Field 6: Type.
	typeB := sb.FieldBuilder(6).(*array.StringBuilder)
	if ni.Type != nil {
		typeB.Append(*ni.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 7: TypeID.
	typeIDB := sb.FieldBuilder(7).(*array.Int32Builder)
	typeIDB.Append(int32(ni.TypeID))

	// Field 8: UID.
	uidB := sb.FieldBuilder(8).(*array.StringBuilder)
	if ni.UID != nil {
		uidB.Append(*ni.UID)
	} else {
		uidB.AppendNull()
	}
}
