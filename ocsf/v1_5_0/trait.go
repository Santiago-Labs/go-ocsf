// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type Trait struct {

	// Category: The high-level grouping or classification this trait belongs to.
	Category *string `json:"category,omitempty" parquet:"category,optional"`

	// Name: The name of the trait.
	Name *string `json:"name,omitempty" parquet:"name,optional"`

	// Type: The type of the trait. For example, this can be used to indicate if the trait acts as a contributing factor (increases risk/severity) or a mitigating factor (decreases risk/severity), in the context of the related finding.
	Type *string `json:"type,omitempty" parquet:"type,optional"`

	// Unique ID: The unique identifier of the trait.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`

	// Values: The values of the trait.
	Values []string `json:"values,omitempty" parquet:"values,optional,list"`
}

var TraitFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "values", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
}

var TraitStruct = arrow.StructOf(TraitFields...)

var TraitSchema = arrow.NewSchema(TraitFields, nil)
