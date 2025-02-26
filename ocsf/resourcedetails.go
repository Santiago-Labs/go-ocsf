package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// ResourceDetailsFields defines the Arrow fields for ResourceDetails.
var ResourceDetailsFields = []arrow.Field{
	{Name: "criticality", Type: arrow.BinaryTypes.String},
	{Name: "data", Type: arrow.BinaryTypes.String},
	{Name: "group", Type: arrow.StructOf(GroupFields...)},
	{Name: "labels", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "uid", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "namespace", Type: arrow.BinaryTypes.String},
	{Name: "owner", Type: arrow.StructOf(UserFields...)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "version", Type: arrow.BinaryTypes.String},
}

// ResourceDetailsSchema is the Arrow schema for ResourceDetails.
var ResourceDetailsSchema = arrow.NewSchema(ResourceDetailsFields, nil)

// ResourceDetails represents a single resource.
type ResourceDetails struct {
	Criticality *string  `json:"criticality"`
	Data        *string  `json:"data"` // JSON blob
	Group       *Group   `json:"group"`
	Labels      []string `json:"labels"`
	UID         *string  `json:"uid"`
	Name        *string  `json:"name"`
	Namespace   *string  `json:"namespace"`
	Owner       *User    `json:"owner"`
	Type        *string  `json:"type"`
	Version     *string  `json:"version"`
}

// WriteToParquet writes the ResourceDetails fields to the provided Arrow StructBuilder.
func (r *ResourceDetails) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: criticality.
	critB := sb.FieldBuilder(0).(*array.StringBuilder)
	if r.Criticality != nil {
		critB.Append(*r.Criticality)
	} else {
		critB.AppendNull()
	}

	// Field 1: data.
	dataB := sb.FieldBuilder(1).(*array.StringBuilder)
	if r.Data != nil {
		dataB.Append(*r.Data)
	} else {
		dataB.AppendNull()
	}

	// Field 2: group (nested struct).
	groupB := sb.FieldBuilder(2).(*array.StructBuilder)
	if r.Group != nil {
		groupB.Append(true)
		r.Group.WriteToParquet(groupB)
	} else {
		groupB.AppendNull()
	}

	// Field 3: labels (list of strings).
	labelsB := sb.FieldBuilder(3).(*array.ListBuilder)
	if len(r.Labels) > 0 {
		labelsB.Append(true)
		labelsValB := labelsB.ValueBuilder().(*array.StringBuilder)
		for _, label := range r.Labels {
			labelsValB.Append(label)
		}
	} else {
		labelsB.AppendNull()
	}

	// Field 4: uid.
	uidB := sb.FieldBuilder(4).(*array.StringBuilder)
	if r.UID != nil {
		uidB.Append(*r.UID)
	} else {
		uidB.AppendNull()
	}

	// Field 5: name.
	nameB := sb.FieldBuilder(5).(*array.StringBuilder)
	if r.Name != nil {
		nameB.Append(*r.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 6: namespace.
	namespaceB := sb.FieldBuilder(6).(*array.StringBuilder)
	if r.Namespace != nil {
		namespaceB.Append(*r.Namespace)
	} else {
		namespaceB.AppendNull()
	}

	// Field 7: owner (nested struct).
	ownerB := sb.FieldBuilder(7).(*array.StructBuilder)
	if r.Owner != nil {
		ownerB.Append(true)
		r.Owner.WriteToParquet(ownerB)
	} else {
		ownerB.AppendNull()
	}

	// Field 8: type.
	typeB := sb.FieldBuilder(8).(*array.StringBuilder)
	if r.Type != nil {
		typeB.Append(*r.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 9: version.
	versionB := sb.FieldBuilder(9).(*array.StringBuilder)
	if r.Version != nil {
		versionB.Append(*r.Version)
	} else {
		versionB.AppendNull()
	}
}
