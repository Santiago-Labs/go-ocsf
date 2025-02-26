package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var GroupSchema = arrow.NewSchema(GroupFields, nil)

// Group represents a user or device group.
type Group struct {
	Desc       *string  `json:"desc,omitempty"`
	Domain     *string  `json:"domain,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	Type       *string  `json:"type,omitempty"`
	UID        *string  `json:"uid,omitempty"`
}

// GroupFields defines the fields for the Group Arrow schema.
var GroupFields = []arrow.Field{
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "domain", Type: arrow.BinaryTypes.String},
	{Name: "name", Type: arrow.BinaryTypes.String},
	{Name: "privileges", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// WriteToParquet writes the Group fields to the provided Arrow StructBuilder.
func (g *Group) WriteToParquet(sb *array.StructBuilder) {

	// Field 0: Desc.
	descB := sb.FieldBuilder(0).(*array.StringBuilder)
	if g.Desc != nil {
		descB.Append(*g.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 1: Domain.
	domainB := sb.FieldBuilder(1).(*array.StringBuilder)
	if g.Domain != nil {
		domainB.Append(*g.Domain)
	} else {
		domainB.AppendNull()
	}

	// Field 2: Name.
	nameB := sb.FieldBuilder(2).(*array.StringBuilder)
	if g.Name != nil {
		nameB.Append(*g.Name)
	} else {
		nameB.AppendNull()
	}

	// Field 3: Privileges (list of strings).
	privilegesB := sb.FieldBuilder(3).(*array.ListBuilder)
	if len(g.Privileges) > 0 {
		privilegesB.Append(true)
		privilegesValB := privilegesB.ValueBuilder().(*array.StringBuilder)
		for _, priv := range g.Privileges {
			privilegesValB.Append(priv)
		}
	} else {
		privilegesB.AppendNull()
	}

	// Field 4: Type.
	typeB := sb.FieldBuilder(4).(*array.StringBuilder)
	if g.Type != nil {
		typeB.Append(*g.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 5: UID.
	uidB := sb.FieldBuilder(5).(*array.StringBuilder)
	if g.UID != nil {
		uidB.Append(*g.UID)
	} else {
		uidB.AppendNull()
	}
}
