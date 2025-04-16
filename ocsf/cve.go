package ocsf

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
)

var CVEFields = []arrow.Field{
	{Name: "created_time", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cvss", Type: arrow.ListOf(CVSSStruct), Nullable: true},
	{Name: "cwe", Type: CWEStruct, Nullable: true},
	{Name: "cwe_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cwe_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "epss", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "modified_time", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "references", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var CVEStruct = arrow.StructOf(CVEFields...)
var CVEClassname = "cve"

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
