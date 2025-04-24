package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ContainerFields defines the Arrow fields for Container.
var ContainerFields = []arrow.Field{
	{Name: "hash", Type: FingerprintStruct, Nullable: true},
	{Name: "image", Type: ImageStruct, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "network_driver", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "orchestrator", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "pod_uuid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "runtime", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "size", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "tag", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "tags", Type: arrow.ListOf(KeyValueObjectStruct), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var ContainerStruct = arrow.StructOf(ContainerFields...)
var ContainerClassname = "container"

type Container struct {
	Hash          *Fingerprint      `json:"hash,omitempty" parquet:"hash,optional" ch:"hash"`
	Image         *Image            `json:"image,omitempty" parquet:"image,optional" ch:"image"`
	Labels        []string          `json:"labels,omitempty" parquet:"labels,list,optional" ch:"labels"`
	Name          *string           `json:"name,omitempty" parquet:"name,optional" ch:"name"`
	NetworkDriver *string           `json:"network_driver,omitempty" parquet:"network_driver,optional" ch:"network_driver"`
	Orchestrator  *string           `json:"orchestrator,omitempty" parquet:"orchestrator,optional" ch:"orchestrator"`
	PodUUID       *string           `json:"pod_uuid,omitempty" parquet:"pod_uuid,optional" ch:"pod_uuid"`
	Runtime       *string           `json:"runtime,omitempty" parquet:"runtime,optional" ch:"runtime"`
	Size          *int64            `json:"size,omitempty" parquet:"size,optional" ch:"size"`
	Tag           *string           `json:"tag,omitempty" parquet:"tag,optional" ch:"tag"`
	Tags          []*KeyValueObject `json:"tags,omitempty" parquet:"tags,list,optional" ch:"tags"`
	UID           *string           `json:"uid,omitempty" parquet:"uid,optional" ch:"uid"`
}
