package tenable

import (
	"time"
)

// Asset represents a Tenable asset
type Asset struct {
	ID                       string   `json:"id"`
	UUID                     string   `json:"uuid"`
	HostName                 string   `json:"hostname"`
	IPV4                     string   `json:"ipv4"`
	LastSeen                 string   `json:"last_seen"`
	FirstSeen                string   `json:"first_seen"`
	OperatingSystem          []string `json:"operating_system"`
	MacAddress               string   `json:"mac_address"`
	AgentUUID                string   `json:"agent_uuid,omitempty"`
	BiosUUID                 string   `json:"bios_uuid,omitempty"`
	DeviceType               string   `json:"device_type,omitempty"`
	FQDN                     string   `json:"fqdn,omitempty"`
	LastAuthenticatedResults string   `json:"last_authenticated_results,omitempty"`
	NetworkID                string   `json:"network_id,omitempty"`
	Tracked                  bool     `json:"tracked,omitempty"`
}

// Vulnerability represents a Tenable vulnerability
type Vulnerability struct {
	PluginID           int       `json:"plugin_id"`
	PluginName         string    `json:"plugin_name"`
	PluginFamily       string    `json:"plugin_family"`
	Severity           int       `json:"severity"`
	Description        string    `json:"description"`
	Solution           string    `json:"solution"`
	FirstSeen          time.Time `json:"first_seen"`
	LastSeen           time.Time `json:"last_seen"`
	VulnerabilityState string    `json:"vulnerability_state"`
	CVEs               []string  `json:"cve"`
	CVSS3BaseScore     float64   `json:"cvss3_base_score"`
	CVSS3Vector        string    `json:"cvss3_vector"`
	RiskFactor         string    `json:"risk_factor"`
	Synopsis           string    `json:"synopsis"`
	SeeAlso            []string  `json:"see_also"`
	PluginOutput       string    `json:"plugin_output"`
	ExploitAvailable   bool      `json:"exploit_available"`
	ExploitFrameworks  []string  `json:"exploit_frameworks"`
	VPRScore           float64   `json:"vpr_score"`
}

// Port represents port information in a finding
type Port struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type Scan struct {
	UUID           string `json:"uuid"`
	StartedAt      string `json:"started_at"`
	ScheduleUUID   string `json:"schedule_uuid"`
	LastScanTarget string `json:"last_scan_target"`
}

type Plugin struct {
	Bid                        []int     `json:"bid"`
	ChecksForDefaultAccount    bool      `json:"checks_for_default_account"`
	ChecksForMalware           bool      `json:"checks_for_malware"`
	Cpe                        []any     `json:"cpe"`
	Description                string    `json:"description"`
	ExploitAvailable           bool      `json:"exploit_available"`
	ExploitFrameworkCanvas     bool      `json:"exploit_framework_canvas"`
	ExploitFrameworkCore       bool      `json:"exploit_framework_core"`
	ExploitFrameworkD2Elliot   bool      `json:"exploit_framework_d2_elliot"`
	ExploitFrameworkExploithub bool      `json:"exploit_framework_exploithub"`
	ExploitFrameworkMetasploit bool      `json:"exploit_framework_metasploit"`
	ExploitedByMalware         bool      `json:"exploited_by_malware"`
	ExploitedByNessus          bool      `json:"exploited_by_nessus"`
	Family                     string    `json:"family"`
	FamilyID                   int       `json:"family_id"`
	HasPatch                   bool      `json:"has_patch"`
	ID                         int       `json:"id"`
	InTheNews                  bool      `json:"in_the_news"`
	Name                       string    `json:"name"`
	ModificationDate           time.Time `json:"modification_date"`
	PublicationDate            time.Time `json:"publication_date"`
	RiskFactor                 string    `json:"risk_factor"`
	SeeAlso                    []string  `json:"see_also"`
	Solution                   string    `json:"solution"`
	Synopsis                   string    `json:"synopsis"`
	UnsupportedByVendor        bool      `json:"unsupported_by_vendor"`
	Version                    string    `json:"version"`
	Xrefs                      []any     `json:"xrefs"`
	Type                       string    `json:"type"`
}

// Finding represents a vulnerability finding from Tenable
type Finding struct {
	Asset                    Asset  `json:"asset"`
	Output                   string `json:"output"`
	Plugin                   Plugin `json:"plugin"`
	Port                     Port   `json:"port"`
	Scan                     Scan   `json:"scan"`
	Severity                 string `json:"severity"`
	SeverityID               int    `json:"severity_id"`
	SeverityDefaultID        int    `json:"severity_default_id"`
	SeverityModificationType string `json:"severity_modification_type"`
	FirstFound               string `json:"first_found"`
	LastFound                string `json:"last_found"`
	State                    string `json:"state"`
	Indexed                  string `json:"indexed"`
	Source                   string `json:"source"`
	FindingID                string `json:"finding_id"`
}

// ExportVulnsResponse represents the response from the export vulnerabilities API
type ExportVulnsResponse struct {
	ExportUUID string `json:"export_uuid"`
}

// ExportStatusResponse represents the status of an export job
type ExportStatusResponse struct {
	Status          string `json:"status"`
	ChunksAvailable []int  `json:"chunks_available"`
	ChunksProcessed []int  `json:"chunks_processed"`
	FinishedChunks  int    `json:"finished_chunks"`
	TotalChunks     int    `json:"total_chunks"`
}

// ExportChunkResponse represents a chunk of exported data
type ExportChunkResponse struct {
	Findings []Finding `json:"findings"`
}

// ErrorResponse represents an error response from the Tenable API
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// WorkbenchVulnerabilitiesResponse represents the response from the workbenches/vulnerabilities endpoint
type WorkbenchVulnerabilitiesResponse struct {
	Vulnerabilities         []Vulnerability `json:"vulnerabilities"`
	TotalVulnerabilityCount int             `json:"total_vulnerability_count"`
	TotalAssetCount         int             `json:"total_asset_count"`
}
