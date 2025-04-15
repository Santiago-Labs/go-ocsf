package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
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
	Hash          *Fingerprint     `json:"hash,omitempty" parquet:"hash"`
	Image         *Image           `json:"image,omitempty" parquet:"image"`
	Labels        []string         `json:"labels,omitempty" parquet:"labels"`
	Name          *string          `json:"name,omitempty" parquet:"name"`
	NetworkDriver *string          `json:"network_driver,omitempty" parquet:"network_driver"`
	Orchestrator  *string          `json:"orchestrator,omitempty" parquet:"orchestrator"`
	PodUUID       *string          `json:"pod_uuid,omitempty" parquet:"pod_uuid"`
	Runtime       *string          `json:"runtime,omitempty" parquet:"runtime"`
	Size          *int64           `json:"size,omitempty" parquet:"size"`
	Tag           *string          `json:"tag,omitempty" parquet:"tag"`
	Tags          []KeyValueObject `json:"tags,omitempty" parquet:"tags"`
	UID           *string          `json:"uid,omitempty" parquet:"uid"`
}
