package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// WindowsServiceFields defines the Arrow fields for Windows Service.
var WindowsServiceFields = []arrow.Field{
	{Name: "cmd_line", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "load_order_group", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: false},
	{Name: "service_category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "service_category_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "service_dependencies", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "service_error_control", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "service_error_control_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "service_start_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "service_start_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "service_start_type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "service_type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "service_type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "tags", Type: arrow.ListOf(KeyValueObjectStruct), Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var WindowsServiceStruct = arrow.StructOf(WindowsServiceFields...)
var WindowsServiceClassname = "windows_service"

type WindowsService struct {
	CmdLine               *string          `json:"cmd_line,omitempty" parquet:"cmd_line,optional"`
	Labels                []string         `json:"labels,omitempty" parquet:"labels,list,optional"`
	LoadOrderGroup        *string          `json:"load_order_group,omitempty" parquet:"load_order_group,optional"`
	Name                  string           `json:"name" parquet:"name"`
	ServiceCategory       *string          `json:"service_category,omitempty" parquet:"service_category,optional"`
	ServiceCategoryID     *int32           `json:"service_category_id,omitempty" parquet:"service_category_id,optional"`
	ServiceDependencies   []string         `json:"service_dependencies,omitempty" parquet:"service_dependencies,list,optional"`
	ServiceErrorControl   *string          `json:"service_error_control,omitempty" parquet:"service_error_control,optional"`
	ServiceErrorControlID *int32           `json:"service_error_control_id,omitempty" parquet:"service_error_control_id,optional"`
	ServiceStartName      *string          `json:"service_start_name,omitempty" parquet:"service_start_name,optional"`
	ServiceStartType      *string          `json:"service_start_type,omitempty" parquet:"service_start_type,optional"`
	ServiceStartTypeID    *int32           `json:"service_start_type_id,omitempty" parquet:"service_start_type_id,optional"`
	ServiceType           *string          `json:"service_type,omitempty" parquet:"service_type,optional"`
	ServiceTypeID         *int32           `json:"service_type_id,omitempty" parquet:"service_type_id,optional"`
	Tags                  []KeyValueObject `json:"tags,omitempty" parquet:"tags,list,optional"`
	UID                   *string          `json:"uid,omitempty" parquet:"uid,optional"`
	Version               *string          `json:"version,omitempty" parquet:"version,optional"`
}
