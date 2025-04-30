package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

var EvidenceArtifactsFields = []arrow.Field{
	{Name: "actor", Type: ActorStruct, Nullable: true},
	{Name: "api", Type: APIStruct, Nullable: true},
	{Name: "connection_info", Type: NetworkConnectionInfoStruct, Nullable: true},
	{Name: "container", Type: ContainerStruct, Nullable: true},
	{Name: "data", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "database", Type: DatabaseStruct, Nullable: true},
	{Name: "databucket", Type: DatabucketStruct, Nullable: true},
	{Name: "device", Type: DeviceStruct, Nullable: true},
	{Name: "dst_endpoint", Type: NetworkEndpointStruct, Nullable: true},
	{Name: "email", Type: EmailStruct, Nullable: true},
	{Name: "file", Type: FileStruct, Nullable: true},
	{Name: "http_request", Type: HTTPRequestStruct, Nullable: true},
	{Name: "http_response", Type: HTTPResponseStruct, Nullable: true},
	{Name: "ja4_fingerprint_list", Type: arrow.ListOf(JA4FingerprintStruct), Nullable: true},
	{Name: "job", Type: JobStruct, Nullable: true},
	{Name: "process", Type: ProcessStruct, Nullable: true},
	{Name: "query", Type: DNSQueryStruct, Nullable: true},
	{Name: "reg_key", Type: WindowsRegistryKeyStruct, Nullable: true},
	{Name: "reg_value", Type: WindowsRegistryValueStruct, Nullable: true},
	{Name: "resources", Type: arrow.ListOf(ResourceDetailsStruct), Nullable: true},
	{Name: "script", Type: ScriptStruct, Nullable: true},
	{Name: "src_endpoint", Type: NetworkEndpointStruct, Nullable: true},
	{Name: "tls", Type: TLSStruct, Nullable: true},
	{Name: "url", Type: URLStruct, Nullable: true},
	{Name: "user", Type: UserStruct, Nullable: true},
	{Name: "win_service", Type: WindowsServiceStruct, Nullable: true},
}

var EvidenceArtifactsStruct = arrow.StructOf(EvidenceArtifactsFields...)
var EvidenceArtifactsClassname = "evidences"

type EvidenceArtifacts struct {
	Actor              *Actor                 `json:"actor,omitempty" parquet:"actor,optional"`
	API                *API                   `json:"api,omitempty" parquet:"api,optional"`
	ConnectionInfo     *NetworkConnectionInfo `json:"connection_info,omitempty" parquet:"connection_info,optional"`
	Container          *Container             `json:"container,omitempty" parquet:"container,optional"`
	Data               *string                `json:"data,omitempty" parquet:"data,optional"`
	Database           *Database              `json:"database,omitempty" parquet:"database,optional"`
	Databucket         *Databucket            `json:"databucket,omitempty" parquet:"databucket,optional"`
	Device             *Device                `json:"device,omitempty" parquet:"device,optional"`
	DstEndpoint        *NetworkEndpoint       `json:"dst_endpoint,omitempty" parquet:"dst_endpoint,optional"`
	Email              *Email                 `json:"email,omitempty" parquet:"email,optional"`
	File               *File                  `json:"file,omitempty" parquet:"file,optional"`
	HTTPRequest        *HTTPRequest           `json:"http_request,omitempty" parquet:"http_request,optional"`
	HTTPResponse       *HTTPResponse          `json:"http_response,omitempty" parquet:"http_response,optional"`
	JA4FingerprintList []*JA4Fingerprint      `json:"ja4_fingerprint_list,omitempty" parquet:"ja4_fingerprint_list,list,optional"`
	Job                *Job                   `json:"job,omitempty" parquet:"job,optional"`
	Process            *Process               `json:"process,omitempty" parquet:"process,optional"`
	Query              *DNSQuery              `json:"query,omitempty" parquet:"query,optional"`
	Resources          []*ResourceDetails     `json:"resources,omitempty" parquet:"resources,list,optional"`
	Script             *Script                `json:"script,omitempty" parquet:"script,optional"`
	SrcEndpoint        *NetworkEndpoint       `json:"src_endpoint,omitempty" parquet:"src_endpoint,optional"`
	TLS                *TLS                   `json:"tls,omitempty" parquet:"tls,optional"`
	UID                *string                `json:"uid,omitempty" parquet:"uid,optional"`
	URL                *URL                   `json:"url,omitempty" parquet:"url,optional"`
	User               *User                  `json:"user,omitempty" parquet:"user,optional"`
	WinService         *WindowsService        `json:"win_service,omitempty" parquet:"win_service,optional"`
}
