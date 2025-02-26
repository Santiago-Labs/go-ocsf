package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// DeviceFields defines the Arrow fields for Device.
var DeviceFields = []arrow.Field{
	{Name: "region", Type: arrow.BinaryTypes.String},
	{Name: "interface_name", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "interface_uid", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "os", Type: arrow.StructOf(OSFields...)},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "hypervisor", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "location", Type: arrow.StructOf(GeoLocationFields...)},
	{Name: "instance_uid", Type: arrow.BinaryTypes.String},
	{Name: "first_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "mac", Type: arrow.BinaryTypes.String},
	{Name: "org", Type: arrow.StructOf(OrganizationFields...)},
	{Name: "risk_level", Type: arrow.BinaryTypes.String},
	{Name: "image", Type: arrow.StructOf(ImageFields...)},
	{Name: "created_time", Type: arrow.BinaryTypes.String},
	{Name: "subnet_uid", Type: arrow.BinaryTypes.String},
	{Name: "zone", Type: arrow.BinaryTypes.String},
	{Name: "groups", Type: arrow.ListOf(arrow.StructOf(GroupFields...))},
	{Name: "risk_score", Type: arrow.PrimitiveTypes.Int32},
	{Name: "is_personal", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "hw_info", Type: arrow.StructOf(DeviceHWInfoFields...)},
	{Name: "uid_alt", Type: arrow.BinaryTypes.String},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "last_seen_time", Type: arrow.BinaryTypes.String},
	{Name: "is_managed", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "risk_level_id", Type: arrow.PrimitiveTypes.Int32},
	{Name: "is_trusted", Type: arrow.FixedWidthTypes.Boolean},
	{Name: "network_interfaces", Type: arrow.ListOf(arrow.StructOf(NetworkInterfaceFields...))},
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

// DeviceSchema is the Arrow schema for Device.
var DeviceSchema = arrow.NewSchema(DeviceFields, nil)

type Device struct {
	Region            *string            `json:"region,omitempty"`
	InterfaceName     *string            `json:"interface_name,omitempty"`
	UID               *string            `json:"uid,omitempty"`
	InterfaceUID      *string            `json:"interface_uid,omitempty"`
	ModifiedTime      *time.Time         `json:"modified_time,omitempty"`
	OS                *OS                `json:"os,omitempty"`
	Desc              *string            `json:"desc,omitempty"`
	Hypervisor        *string            `json:"hypervisor,omitempty"`
	Type              *string            `json:"type,omitempty"`
	Location          *GeoLocation       `json:"location,omitempty"`
	InstanceUID       *string            `json:"instance_uid,omitempty"`
	FirstSeenTime     *time.Time         `json:"first_seen_time,omitempty"`
	MAC               *string            `json:"mac,omitempty"`
	Org               *Organization      `json:"org,omitempty"`
	RiskLevel         *string            `json:"risk_level,omitempty"`
	Image             *Image             `json:"image,omitempty"`
	CreatedTime       *time.Time         `json:"created_time,omitempty"`
	SubnetUID         *string            `json:"subnet_uid,omitempty"`
	Zone              *string            `json:"zone,omitempty"`
	Groups            []Group            `json:"groups,omitempty"`
	RiskScore         *int               `json:"risk_score,omitempty"`
	IsPersonal        *bool              `json:"is_personal,omitempty"`
	Name              *string            `json:"name,omitempty"`
	HWInfo            *DeviceHWInfo      `json:"hw_info,omitempty"`
	UIDAlt            *string            `json:"uid_alt,omitempty"`
	TypeID            int                `json:"type_id"`
	LastSeenTime      *time.Time         `json:"last_seen_time,omitempty"`
	IsManaged         *bool              `json:"is_managed,omitempty"`
	RiskLevelID       *int               `json:"risk_level_id,omitempty"`
	IsTrusted         *bool              `json:"is_trusted,omitempty"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces,omitempty"`
	AutoscaleUID      *string            `json:"autoscale_uid,omitempty"`
	VpcUID            *string            `json:"vpc_uid,omitempty"`
	Subnet            *string            `json:"subnet,omitempty"`
	Domain            *string            `json:"domain,omitempty"`
	IMEI              *string            `json:"imei,omitempty"`
	IP                *string            `json:"ip,omitempty"`
	Hostname          *string            `json:"hostname,omitempty"`
	VlanUID           *string            `json:"vlan_uid,omitempty"`
	IsCompliant       *bool              `json:"is_compliant,omitempty"`
}

func (d *Device) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: Region.
	regionB := sb.FieldBuilder(0).(*array.StringBuilder)
	if d.Region != nil {
		regionB.Append(*d.Region)
	} else {
		regionB.AppendNull()
	}

	// Field 1: InterfaceName.
	ifaceB := sb.FieldBuilder(1).(*array.StringBuilder)
	if d.InterfaceName != nil {
		ifaceB.Append(*d.InterfaceName)
	} else {
		ifaceB.AppendNull()
	}

	// Field 2: UID.
	uidB := sb.FieldBuilder(2).(*array.StringBuilder)
	if d.UID != nil {
		uidB.Append(*d.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 3: InterfaceUID.
	ifaceUIDB := sb.FieldBuilder(3).(*array.StringBuilder)
	if d.InterfaceUID != nil {
		ifaceUIDB.Append(*d.InterfaceUID)
	} else {
		ifaceUIDB.AppendNull()
	}

	// Field 4: ModifiedTime.
	modTimeB := sb.FieldBuilder(4).(*array.StringBuilder)
	if d.ModifiedTime != nil {
		modTimeB.Append(d.ModifiedTime.Format(time.RFC3339))
	} else {
		modTimeB.AppendNull()
	}

	// Field 5: OS (nested struct).
	osB := sb.FieldBuilder(5).(*array.StructBuilder)
	if d.OS != nil {
		osB.Append(true)
		d.OS.WriteToParquet(osB)
	} else {
		osB.AppendNull()
	}

	// Field 6: Desc.
	descB := sb.FieldBuilder(6).(*array.StringBuilder)
	if d.Desc != nil {
		descB.Append(*d.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 7: Hypervisor.
	hyperB := sb.FieldBuilder(7).(*array.StringBuilder)
	if d.Hypervisor != nil {
		hyperB.Append(*d.Hypervisor)
	} else {
		hyperB.AppendNull()
	}

	// Field 8: Type.
	typeB := sb.FieldBuilder(8).(*array.StringBuilder)
	if d.Type != nil {
		typeB.Append(*d.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 9: Location (nested struct).
	locB := sb.FieldBuilder(9).(*array.StructBuilder)
	if d.Location != nil {
		locB.Append(true)
		d.Location.WriteToParquet(locB)
	} else {
		locB.AppendNull()
	}

	// Field 10: InstanceUID.
	instanceB := sb.FieldBuilder(10).(*array.StringBuilder)
	if d.InstanceUID != nil {
		instanceB.Append(*d.InstanceUID)
	} else {
		instanceB.AppendNull()
	}

	// Field 11: FirstSeenTime.
	fstB := sb.FieldBuilder(11).(*array.StringBuilder)
	if d.FirstSeenTime != nil {
		fstB.Append(d.FirstSeenTime.Format(time.RFC3339))
	} else {
		fstB.AppendNull()
	}

	// Field 12: MAC.
	macB := sb.FieldBuilder(12).(*array.StringBuilder)
	if d.MAC != nil {
		macB.Append(*d.MAC)
	} else {
		macB.AppendNull()
	}

	// Field 13: Org (nested struct).
	orgB := sb.FieldBuilder(13).(*array.StructBuilder)
	if d.Org != nil {
		orgB.Append(true)
		d.Org.WriteToParquet(orgB)
	} else {
		orgB.AppendNull()
	}

	// Field 14: RiskLevel.
	riskLevelB := sb.FieldBuilder(14).(*array.StringBuilder)
	if d.RiskLevel != nil {
		riskLevelB.Append(*d.RiskLevel)
	} else {
		riskLevelB.AppendNull()
	}

	// Field 15: Image (nested struct).
	imageB := sb.FieldBuilder(15).(*array.StructBuilder)
	if d.Image != nil {
		imageB.Append(true)
		d.Image.WriteToParquet(imageB)
	} else {
		imageB.AppendNull()
	}

	// Field 16: CreatedTime.
	createdTimeB := sb.FieldBuilder(16).(*array.StringBuilder)
	if d.CreatedTime != nil {
		createdTimeB.Append(d.CreatedTime.Format(time.RFC3339))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 17: SubnetUID.
	subnetUIDB := sb.FieldBuilder(17).(*array.StringBuilder)
	if d.SubnetUID != nil {
		subnetUIDB.Append(*d.SubnetUID)
	} else {
		subnetUIDB.AppendNull()
	}

	// Field 18: Zone.
	zoneB := sb.FieldBuilder(18).(*array.StringBuilder)
	if d.Zone != nil {
		zoneB.Append(*d.Zone)
	} else {
		zoneB.AppendNull()
	}

	// Field 19: Groups (list of nested Group structs).
	groupsB := sb.FieldBuilder(19).(*array.ListBuilder)
	groupsValB := groupsB.ValueBuilder().(*array.StructBuilder)
	if len(d.Groups) > 0 {
		groupsB.Append(true)
		for _, g := range d.Groups {
			groupsValB.Append(true)
			g.WriteToParquet(groupsValB)
		}
	} else {
		groupsB.AppendNull()
	}

	// Field 20: RiskScore.
	riskScoreB := sb.FieldBuilder(20).(*array.Int32Builder)
	if d.RiskScore != nil {
		riskScoreB.Append(int32(*d.RiskScore))
	} else {
		riskScoreB.AppendNull()
	}

	// Field 21: IsPersonal.
	isPersonalB := sb.FieldBuilder(21).(*array.BooleanBuilder)
	if d.IsPersonal != nil {
		isPersonalB.Append(*d.IsPersonal)
	} else {
		isPersonalB.AppendNull()
	}

	// Field 22: Name.
	nameB := sb.FieldBuilder(22).(*array.StringBuilder)
	if d.Name != nil {
		nameB.Append(*d.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 23: HWInfo (nested struct).
	hwInfoB := sb.FieldBuilder(23).(*array.StructBuilder)
	if d.HWInfo != nil {
		hwInfoB.Append(true)
		d.HWInfo.WriteToParquet(hwInfoB)
	} else {
		hwInfoB.AppendNull()
	}

	// Field 24: UIDAlt.
	uidAltB := sb.FieldBuilder(24).(*array.StringBuilder)
	if d.UIDAlt != nil {
		uidAltB.Append(*d.UIDAlt)
	} else {
		uidAltB.AppendNull()
	}

	// Field 25: TypeID.
	typeIDB := sb.FieldBuilder(25).(*array.Int32Builder)
	typeIDB.Append(int32(d.TypeID))

	// Field 26: LastSeenTime.
	lastSeenB := sb.FieldBuilder(26).(*array.StringBuilder)
	if d.LastSeenTime != nil {
		lastSeenB.Append(d.LastSeenTime.Format(time.RFC3339))
	} else {
		lastSeenB.AppendNull()
	}

	// Field 27: IsManaged.
	isManagedB := sb.FieldBuilder(27).(*array.BooleanBuilder)
	if d.IsManaged != nil {
		isManagedB.Append(*d.IsManaged)
	} else {
		isManagedB.AppendNull()
	}

	// Field 28: RiskLevelID.
	riskLevelIDB := sb.FieldBuilder(28).(*array.Int32Builder)
	if d.RiskLevelID != nil {
		riskLevelIDB.Append(int32(*d.RiskLevelID))
	} else {
		riskLevelIDB.AppendNull()
	}

	// Field 29: IsTrusted.
	isTrustedB := sb.FieldBuilder(29).(*array.BooleanBuilder)
	if d.IsTrusted != nil {
		isTrustedB.Append(*d.IsTrusted)
	} else {
		isTrustedB.AppendNull()
	}

	// Field 30: NetworkInterfaces (list of nested NetworkInterface structs).
	netIfB := sb.FieldBuilder(30).(*array.ListBuilder)
	netIfValB := netIfB.ValueBuilder().(*array.StructBuilder)
	if len(d.NetworkInterfaces) > 0 {
		netIfB.Append(true)
		for _, ni := range d.NetworkInterfaces {
			netIfValB.Append(true)
			ni.WriteToParquet(netIfValB)
		}
	} else {
		netIfB.AppendNull()
	}

	// Field 31: AutoscaleUID.
	autoscaleB := sb.FieldBuilder(31).(*array.StringBuilder)
	if d.AutoscaleUID != nil {
		autoscaleB.Append(*d.AutoscaleUID)
	} else {
		autoscaleB.AppendNull()
	}

	// Field 32: VpcUID.
	vpcB := sb.FieldBuilder(32).(*array.StringBuilder)
	if d.VpcUID != nil {
		vpcB.Append(*d.VpcUID)
	} else {
		vpcB.AppendNull()
	}

	// Field 33: Subnet.
	subnetB := sb.FieldBuilder(33).(*array.StringBuilder)
	if d.Subnet != nil {
		subnetB.Append(*d.Subnet)
	} else {
		subnetB.AppendNull()
	}

	// Field 34: Domain.
	domainB := sb.FieldBuilder(34).(*array.StringBuilder)
	if d.Domain != nil {
		domainB.Append(*d.Domain)
	} else {
		domainB.AppendNull()
	}

	// Field 35: IMEI.
	imeiB := sb.FieldBuilder(35).(*array.StringBuilder)
	if d.IMEI != nil {
		imeiB.Append(*d.IMEI)
	} else {
		imeiB.AppendNull()
	}

	// Field 36: IP.
	ipB := sb.FieldBuilder(36).(*array.StringBuilder)
	if d.IP != nil {
		ipB.Append(*d.IP)
	} else {
		ipB.AppendNull()
	}

	// Field 37: Hostname.
	hostnameB := sb.FieldBuilder(37).(*array.StringBuilder)
	if d.Hostname != nil {
		hostnameB.Append(*d.Hostname)
	} else {
		hostnameB.AppendNull()
	}

	// Field 38: VlanUID.
	vlanB := sb.FieldBuilder(38).(*array.StringBuilder)
	if d.VlanUID != nil {
		vlanB.Append(*d.VlanUID)
	} else {
		vlanB.AppendNull()
	}

	// Field 39: IsCompliant.
	isCompliantB := sb.FieldBuilder(39).(*array.BooleanBuilder)
	if d.IsCompliant != nil {
		isCompliantB.Append(*d.IsCompliant)
	} else {
		isCompliantB.AppendNull()
	}
}
