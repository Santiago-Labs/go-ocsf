package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
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
	CreatedTime  *time.Time `json:"created_time,omitempty" parquet:"created_time"`
	CVSS         []CVSS     `json:"cvss,omitempty" parquet:"cvss"`
	CWE          *CWE       `json:"cwe,omitempty" parquet:"cwe"`
	CWEUID       *string    `json:"cwe_uid,omitempty" parquet:"cwe_uid"`
	CWEURL       *string    `json:"cwe_url,omitempty" parquet:"cwe_url"`
	Desc         *string    `json:"desc,omitempty" parquet:"desc"`
	EPSS         *EPSS      `json:"epss,omitempty" parquet:"epss"`
	ModifiedTime *time.Time `json:"modified_time,omitempty" parquet:"modified_time"`
	Product      *Product   `json:"product,omitempty" parquet:"product"`
	References   []string   `json:"references,omitempty" parquet:"references"`
	Title        *string    `json:"title,omitempty" parquet:"title"`
	Type         *string    `json:"type,omitempty" parquet:"type"`
	UID          string     `json:"uid" parquet:"uid"`
}
