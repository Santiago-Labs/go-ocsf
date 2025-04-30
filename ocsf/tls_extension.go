package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var TLSExtensionFields = []arrow.Field{
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type_id", Type: arrow.PrimitiveTypes.Int32, Nullable: false},
}

var TLSExtensionStruct = arrow.StructOf(TLSExtensionFields...)
var TLSExtensionClassname = "tls_extension"

type TLSExtension struct {
	Type   *string `json:"type,omitempty" parquet:"type,optional"`
	TypeID int32   `json:"type_id" parquet:"type_id"`
}
