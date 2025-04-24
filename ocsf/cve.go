package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var CVEFields = []arrow.Field{
	{Name: "created_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "cvss", Type: arrow.ListOf(CVSSStruct), Nullable: true},
	{Name: "cwe", Type: CWEStruct, Nullable: true},
	{Name: "cwe_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "cwe_url", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "epss", Type: EPSSStruct, Nullable: true},
	{Name: "modified_time", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "product", Type: ProductStruct, Nullable: true},
	{Name: "references", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "title", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: false},
}

var CVEStruct = arrow.StructOf(CVEFields...)
var CVEClassname = "cve"

type CVE struct {
	CreatedTime  *int64   `json:"created_time,omitempty" parquet:"created_time,optional" ch:"created_time"`
	CVSS         []*CVSS  `json:"cvss,omitempty" parquet:"cvss,list,optional" ch:"cvss"`
	CWE          *CWE     `json:"cwe,omitempty" parquet:"cwe,optional" ch:"cwe"`
	CWEUID       *string  `json:"cwe_uid,omitempty" parquet:"cwe_uid,optional" ch:"cwe_uid"`
	CWEURL       *string  `json:"cwe_url,omitempty" parquet:"cwe_url,optional" ch:"cwe_url"`
	Desc         *string  `json:"desc,omitempty" parquet:"desc,optional" ch:"desc"`
	EPSS         *EPSS    `json:"epss,omitempty" parquet:"epss,optional" ch:"epss"`
	ModifiedTime *int64   `json:"modified_time,omitempty" parquet:"modified_time,optional" ch:"modified_time"`
	Product      *Product `json:"product,omitempty" parquet:"product,optional" ch:"product"`
	References   []string `json:"references,omitempty" parquet:"references,list,optional" ch:"references"`
	Title        *string  `json:"title,omitempty" parquet:"title,optional" ch:"title"`
	Type         *string  `json:"type,omitempty" parquet:"type,optional" ch:"type"`
	UID          string   `json:"uid" parquet:"uid" ch:"uid"`
}
