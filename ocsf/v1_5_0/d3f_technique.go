// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type MITRED3FENDTechnique struct {

	// Name: The name of the defensive technique. For example: <code>IO Port Restriction</code>.
	Name *string `json:"name,omitempty" parquet:"name,optional"`

	// Source URL: The versioned permalink of the defensive technique. For example: <code>https://d3fend.mitre.org/technique/d3f:IOPortRestriction/</code>.
	SrcUrl *string `json:"src_url,omitempty" parquet:"src_url,optional"`

	// Unique ID: The unique identifier of the defensive technique. For example: <code>D3-IOPR</code>.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`
}

var MITRED3FENDTechniqueFields = []arrow.Field{
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "src_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
}

var MITRED3FENDTechniqueStruct = arrow.StructOf(MITRED3FENDTechniqueFields...)

var MITRED3FENDTechniqueSchema = arrow.NewSchema(MITRED3FENDTechniqueFields, nil)
