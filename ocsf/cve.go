package ocsf

import (
	"encoding/json"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
)

var CVEFields = []arrow.Field{
	{Name: "created_time", Type: arrow.BinaryTypes.String},
	{Name: "cvss", Type: arrow.BinaryTypes.String},
	{Name: "cwe", Type: arrow.StructOf(CWEFields...)},
	{Name: "cwe_uid", Type: arrow.BinaryTypes.String},
	{Name: "cwe_url", Type: arrow.BinaryTypes.String},
	{Name: "desc", Type: arrow.BinaryTypes.String},
	{Name: "epss", Type: arrow.BinaryTypes.String},
	{Name: "modified_time", Type: arrow.BinaryTypes.String},
	{Name: "product", Type: arrow.StructOf(ProductFields...)},
	{Name: "references", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	{Name: "title", Type: arrow.BinaryTypes.String},
	{Name: "type", Type: arrow.BinaryTypes.String},
	{Name: "uid", Type: arrow.BinaryTypes.String},
}

// CVESchema is the Arrow schema for CVE.
var CVESchema = arrow.NewSchema(CVEFields, nil)

type CVE struct {
	CreatedTime  *time.Time `json:"created_time,omitempty"`
	CVSS         []CVSS     `json:"cvss,omitempty"`
	CWE          *CWE       `json:"cwe,omitempty"`
	CWEUID       *string    `json:"cwe_uid,omitempty"`
	CWEURL       *string    `json:"cwe_url,omitempty"`
	Desc         *string    `json:"desc,omitempty"`
	EPSS         *EPSS      `json:"epss,omitempty"`
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	Product      *Product   `json:"product,omitempty"`
	References   []string   `json:"references,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Type         *string    `json:"type,omitempty"`
	UID          string     `json:"uid"`
}

func (c *CVE) WriteToParquet(sb *array.StructBuilder) {
	// Field 0: CreatedTime (formatted as RFC3339 string)
	createdTimeB := sb.FieldBuilder(0).(*array.StringBuilder)
	if c.CreatedTime != nil {
		createdTimeB.Append(c.CreatedTime.Format(time.RFC3339))
	} else {
		createdTimeB.AppendNull()
	}

	// Field 1: CVSS slice; for simplicity, we marshal to JSON.
	cvssB := sb.FieldBuilder(1).(*array.StringBuilder)
	if len(c.CVSS) > 0 {
		if b, err := json.Marshal(c.CVSS); err == nil {
			cvssB.Append(string(b))
		} else {
			cvssB.Append("")
		}
	} else {
		cvssB.AppendNull()
	}

	// Field 2: CWE (nested struct)
	cweB := sb.FieldBuilder(2).(*array.StructBuilder)
	if c.CWE != nil {
		cweB.Append(true)
		c.CWE.WriteToParquet(cweB)
	} else {
		cweB.AppendNull()
	}

	// Field 3: CWEUID.
	cweUIDB := sb.FieldBuilder(3).(*array.StringBuilder)
	if c.CWEUID != nil {
		cweUIDB.Append(*c.CWEUID)
	} else {
		cweUIDB.AppendNull()
	}

	// Field 4: CWEURL.
	cweURLB := sb.FieldBuilder(4).(*array.StringBuilder)
	if c.CWEURL != nil {
		cweURLB.Append(*c.CWEURL)
	} else {
		cweURLB.AppendNull()
	}

	// Field 5: Desc.
	descB := sb.FieldBuilder(5).(*array.StringBuilder)
	if c.Desc != nil {
		descB.Append(*c.Desc)
	} else {
		descB.AppendNull()
	}

	// Field 6: EPSS; marshal to JSON.
	epssB := sb.FieldBuilder(6).(*array.StringBuilder)
	if c.EPSS != nil {
		if b, err := json.Marshal(c.EPSS); err == nil {
			epssB.Append(string(b))
		} else {
			epssB.Append("")
		}
	} else {
		epssB.AppendNull()
	}

	// Field 7: ModifiedTime.
	modTimeB := sb.FieldBuilder(7).(*array.StringBuilder)
	if c.ModifiedTime != nil {
		modTimeB.Append(c.ModifiedTime.Format(time.RFC3339))
	} else {
		modTimeB.AppendNull()
	}

	// Field 8: Product (nested struct)
	productB := sb.FieldBuilder(8).(*array.StructBuilder)
	if c.Product != nil {
		productB.Append(true)
		c.Product.WriteToParquet(productB)
	} else {
		productB.AppendNull()
	}

	// Field 9: References (list of strings).
	refB := sb.FieldBuilder(9).(*array.ListBuilder)
	refValB := refB.ValueBuilder().(*array.StringBuilder)
	if len(c.References) > 0 {
		refB.Append(true)
		for _, ref := range c.References {
			refValB.Append(ref)
		}
	} else {
		refB.AppendNull()
	}

	// Field 10: Title.
	titleB := sb.FieldBuilder(10).(*array.StringBuilder)
	if c.Title != nil {
		titleB.Append(*c.Title)
	} else {
		titleB.AppendNull()
	}

	// Field 11: Type.
	typeB := sb.FieldBuilder(11).(*array.StringBuilder)
	if c.Type != nil {
		typeB.Append(*c.Type)
	} else {
		typeB.AppendNull()
	}

	// Field 12: UID.
	uidB := sb.FieldBuilder(12).(*array.StringBuilder)
	uidB.Append(c.UID)
}
