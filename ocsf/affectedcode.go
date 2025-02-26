package ocsf

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

// AffectedCodeFields defines the Arrow fields for AffectedCode.
var AffectedCodeFields = []arrow.Field{
	{Name: "end_line", Type: arrow.PrimitiveTypes.Int32},
	{Name: "start_line", Type: arrow.PrimitiveTypes.Int32},
	{Name: "file", Type: arrow.StructOf(FileFields...)},
	{Name: "owner", Type: arrow.StructOf(UserFields...)},
	{Name: "remediation", Type: arrow.StructOf(RemediationFields...)},
}

// AffectedCodeSchema is the Arrow schema for AffectedCode.
var AffectedCodeSchema = arrow.NewSchema(AffectedCodeFields, nil)

type AffectedCode struct {
	EndLine     int32        `json:"end_line"`
	StartLine   int32        `json:"start_line"`
	File        File         `json:"file"`
	Owner       *User        `json:"owner"`
	Remediation *Remediation `json:"remediation"`
}

func (ac *AffectedCode) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: EndLine.
	endLineB := sb.FieldBuilder(0).(*array.Int32Builder)
	endLineB.Append(ac.EndLine)

	// Field 1: StartLine.
	startLineB := sb.FieldBuilder(1).(*array.Int32Builder)
	startLineB.Append(ac.StartLine)

	// Field 2: File (assume File has its own WriteToParquet method).
	fileB := sb.FieldBuilder(2).(*array.StructBuilder)
	ac.File.WriteToParquet(fileB)

	// Field 3: Owner.
	ownerB := sb.FieldBuilder(3).(*array.StructBuilder)
	if ac.Owner != nil {
		ownerB.Append(true)
		ac.Owner.WriteToParquet(ownerB)
	} else {
		ownerB.AppendNull()
	}

	// Field 4: Remediation.
	remediationB := sb.FieldBuilder(4).(*array.StructBuilder)
	if ac.Remediation != nil {
		remediationB.Append(true)
		ac.Remediation.WriteToParquet(remediationB)
	} else {
		remediationB.AppendNull()
	}
}
