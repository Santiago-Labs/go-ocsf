package snyk

import (
	"encoding/json"
	"fmt"
	"time"
)

type ProjectResponse struct {
	Data    Project    `json:"data"`
	JSONAPI JSONAPI    `json:"jsonapi"`
	Links   APILinks   `json:"links"`
	Errors  []APIError `json:"errors,omitempty"`
}

type Project struct {
	Attributes ProjectAttributes `json:"attributes"`
	ID         string            `json:"id"`
	Type       string            `json:"type"`
}

type ProjectAttributes struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	TargetFile      string `json:"target_file"`
	TargetReference string `json:"target_reference"`
	Origin          string `json:"origin"`
}

type Semver struct {
	HashesRange []string `json:"hashesRange"`
	Vulnerable  []string `json:"vulnerable"`
}
type Identifiers struct {
	Cve   []string `json:"CVE"`
	Cwe   []string `json:"CWE"`
	Osvdb []any    `json:"OSVDB"`
}
type ReportResults struct {
	Project        Project `json:"project"`
	Issue          Issue   `json:"issue"`
	IsFixed        bool    `json:"isFixed"`
	IntroducedDate string  `json:"introducedDate"`
}

type SnykTestResult struct {
	Vulnerabilities []SnykVulnerability `json:"vulnerabilities"`
	Ok              bool                `json:"ok"`
	DependencyCount int                 `json:"dependencyCount"`
	Org             string              `json:"org"`
	Policy          string              `json:"policy"`
	IsPrivate       bool                `json:"isPrivate"`
	PackageManager  string              `json:"packageManager"`
	IgnoreSettings  struct {
		AdminOnly                  bool `json:"adminOnly"`
		ReasonRequired             bool `json:"reasonRequired"`
		DisregardFilesystemIgnores bool `json:"disregardFilesystemIgnores"`
		AutoApproveIgnores         bool `json:"autoApproveIgnores"`
	} `json:"ignoreSettings"`
	Summary          string `json:"summary"`
	FilesystemPolicy bool   `json:"filesystemPolicy"`
	Filtered         struct {
		Ignore []any `json:"ignore"`
		Patch  []any `json:"patch"`
	} `json:"filtered"`
	UniqueCount        int    `json:"uniqueCount"`
	TargetFile         string `json:"targetFile"`
	ProjectName        string `json:"projectName"`
	DisplayTargetFile  string `json:"displayTargetFile"`
	HasUnknownVersions bool   `json:"hasUnknownVersions"`
	Path               string `json:"path"`
}

type SnykVulnerability struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	CVSSv3 string   `json:"CVSSv3,omitempty"`
	Credit []string `json:"credit,omitempty"`
	Semver []struct {
		Vulnerable       []string `json:"vulnerable"`
		VulnerableHashes []string `json:"vulnerableHashes"`
	} `json:"semver,omitempty"`
	Exploit  string `json:"exploit,omitempty"`
	FixedIn  []any  `json:"fixedIn,omitempty"`
	Patches  []any  `json:"patches,omitempty"`
	Insights struct {
		TriageAdvice any `json:"triageAdvice"`
	} `json:"insights,omitempty"`
	Language   string  `json:"language"`
	Severity   string  `json:"severity"`
	CvssScore  float64 `json:"cvssScore,omitempty"`
	Functions  []any   `json:"functions,omitempty"`
	Malicious  bool    `json:"malicious,omitempty"`
	IsDisputed bool    `json:"isDisputed,omitempty"`
	ModuleName string  `json:"moduleName,omitempty"`
	References []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"references,omitempty"`
	CvssDetails []struct {
		Assigner         string    `json:"assigner"`
		Severity         string    `json:"severity"`
		CvssV3Vector     string    `json:"cvssV3Vector"`
		CvssV3BaseScore  float64   `json:"cvssV3BaseScore"`
		ModificationTime time.Time `json:"modificationTime"`
	} `json:"cvssDetails,omitempty"`
	CvssSources []struct {
		Type             string    `json:"type"`
		Vector           string    `json:"vector"`
		Assigner         string    `json:"assigner"`
		Severity         string    `json:"severity"`
		BaseScore        float64   `json:"baseScore"`
		CvssVersion      string    `json:"cvssVersion"`
		ModificationTime time.Time `json:"modificationTime"`
	} `json:"cvssSources,omitempty"`
	Description string `json:"description"`
	EpssDetails struct {
		Percentile   string `json:"percentile"`
		Probability  string `json:"probability"`
		ModelVersion string `json:"modelVersion"`
	} `json:"epssDetails,omitempty"`
	Identifiers struct {
		Cve  []string `json:"CVE"`
		Cwe  []string `json:"CWE"`
		Ghsa []string `json:"GHSA"`
	} `json:"identifiers,omitempty"`
	PackageName    string    `json:"packageName"`
	Proprietary    bool      `json:"proprietary,omitempty"`
	CreationTime   time.Time `json:"creationTime"`
	FunctionsNew   []any     `json:"functions_new,omitempty"`
	AlternativeIds []any     `json:"alternativeIds,omitempty"`
	DisclosureTime time.Time `json:"disclosureTime,omitempty"`
	ExploitDetails struct {
		Sources        []any `json:"sources"`
		MaturityLevels []struct {
			Type   string `json:"type"`
			Level  string `json:"level"`
			Format string `json:"format"`
		} `json:"maturityLevels"`
	} `json:"exploitDetails,omitempty"`
	PackageManager       string    `json:"packageManager"`
	PublicationTime      time.Time `json:"publicationTime"`
	SeverityBasedOn      string    `json:"severityBasedOn,omitempty"`
	ModificationTime     time.Time `json:"modificationTime,omitempty"`
	SocialTrendAlert     bool      `json:"socialTrendAlert,omitempty"`
	SeverityWithCritical string    `json:"severityWithCritical"`
	From                 []string  `json:"from"`
	UpgradePath          []any     `json:"upgradePath"`
	IsUpgradable         bool      `json:"isUpgradable"`
	IsPatchable          bool      `json:"isPatchable"`
	Name                 string    `json:"name"`
	Version              string    `json:"version"`
	Type                 string    `json:"type,omitempty"`
	License              string    `json:"license,omitempty"`
}

type SnykTestAllProjectsResult []SnykTestResult

type OrganizationRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TargetRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type UserRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// IssuesResponse corresponds to the top-level API response
// e.g. GET /rest/orgs/<orgId>/issues or GET /rest/orgs/<orgId>/issues/<issueId>
type IssuesResponse struct {
	Data    []Issue    `json:"data"`
	JSONAPI JSONAPI    `json:"jsonapi"`
	Links   APILinks   `json:"links"`
	Errors  []APIError `json:"errors,omitempty"` // if any error objects are returned
}

type APILinks struct {
	Self  LinkObject `json:"self,omitempty"`
	Next  LinkObject `json:"next,omitempty"`
	Prev  LinkObject `json:"prev,omitempty"`
	Last  LinkObject `json:"last,omitempty"`
	First LinkObject `json:"first,omitempty"`
}

type LinkObject struct {
	StringVal *string        `json:"-"`
	ObjectVal *LinkObjectVal `json:"-"`
}

type LinkObjectVal struct {
	Href string                 `json:"href"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// MarshalJSON handles turning LinkObject into JSON.
func (l LinkObject) MarshalJSON() ([]byte, error) {
	if l.StringVal != nil && l.ObjectVal == nil {
		return json.Marshal(l.StringVal)
	}
	if l.ObjectVal != nil {
		return json.Marshal(l.ObjectVal)
	}
	return []byte("null"), nil
}

func (l *LinkObject) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		l.StringVal = &s
		l.ObjectVal = nil
		return nil
	}

	var obj LinkObjectVal
	if err := json.Unmarshal(data, &obj); err == nil {
		l.StringVal = nil
		l.ObjectVal = &obj
		return nil
	}

	return fmt.Errorf("invalid link: %s", string(data))
}

// JSONAPI describes the "jsonapi" object
type JSONAPI struct {
	Version string `json:"version"`
}

// APIError if the response includes an error array instead of data
type APIError struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
	ID     string `json:"id"`
	Title  string `json:"title"`
	Meta   struct {
		Created time.Time `json:"created"`
	} `json:"meta"`
}

type Issue struct {
	ID            string          `json:"id"`
	Type          string          `json:"type"` // must be "issue"
	Attributes    IssueAttributes `json:"attributes"`
	Relationships IssueRelations  `json:"relationships"`
}

// IssueRelations references the "relationships" section of an Issue
type IssueRelations struct {
	Ignore         *Relationship `json:"ignore,omitempty"`
	Organization   Relationship  `json:"organization"`
	ScanItem       Relationship  `json:"scan_item"`
	TestExecutions struct {
		Data []RelationshipData `json:"data"`
	} `json:"test_executions"`
}

// Relationship is a general form for "ignore", "organization", "scan_item" etc.
type Relationship struct {
	Data *RelationshipData `json:"data,omitempty"`
}

// RelationshipData holds the type/id for a related resource
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type IssueAttributes struct {
	Classes                []Class         `json:"classes"`     // at least 1
	Coordinates            []Coordinate    `json:"coordinates"` // at least 1
	CreatedAt              time.Time       `json:"created_at"`
	Description            string          `json:"description,omitempty"`    // optional markdown
	EffectiveSeverityLevel string          `json:"effective_severity_level"` // info|low|medium|high|critical
	ExploitDetails         *ExploitDetails `json:"exploit_details,omitempty"`
	Ignored                bool            `json:"ignored"`
	Key                    string          `json:"key"`                  // unique key
	Problems               []Problem       `json:"problems"`             // at least 1
	Resolution             *Resolution     `json:"resolution,omitempty"` // if resolved
	Risk                   *Risk           `json:"risk,omitempty"`
	Severities             []Severity      `json:"severities,omitempty"` // multiple severities
	Status                 string          `json:"status"`               // open|resolved
	Title                  string          `json:"title"`
	Tool                   string          `json:"tool,omitempty"`
	Type                   string          `json:"type"` // e.g. package_vulnerability
	UpdatedAt              time.Time       `json:"updated_at"`
}

type Class struct {
	ID     string  `json:"id"`
	Source string  `json:"source"`
	Type   string  `json:"type"` // rule-category|compliance|weakness
	URL    *string `json:"url,omitempty"`
}

type Coordinate struct {
	IsFixableManually bool             `json:"is_fixable_manually"`
	IsFixableSnyk     bool             `json:"is_fixable_snyk"`
	IsFixableUpstream bool             `json:"is_fixable_upstream"`
	IsPatchable       bool             `json:"is_patchable"`
	IsPinnable        bool             `json:"is_pinnable"`
	IsUpgradeable     bool             `json:"is_upgradeable"`
	Reachability      string           `json:"reachability,omitempty"` // function|package|no-info|not-applicable
	Remedies          []Remedy         `json:"remedies,omitempty"`     // min 1..5
	Representations   []Representation `json:"representations,omitempty"`
}

// Remedy is an object describing how to fix or mitigate
type Remedy struct {
	CorrelationID string      `json:"correlation_id,omitempty"`
	Description   string      `json:"description,omitempty"`
	Meta          *RemedyMeta `json:"meta,omitempty"`
	Type          string      `json:"type"` // indeterminate|manual|automated|rule_result_message|terraform|cloudformation|cli|kubernetes|arm
}

// RemedyMeta includes data and a schema_version
type RemedyMeta struct {
	Data          map[string]interface{} `json:"data"` // any key/val, up to 100KB
	SchemaVersion string                 `json:"schema_version"`
}

type Representation struct {
	ResourcePath *string `json:"resourcePath,omitempty"`

	Dependency *struct {
		PackageName    string `json:"package_name"`
		PackageVersion string `json:"package_version"`
	} `json:"dependency,omitempty"`

	CloudResource *struct {
		Environment struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			NativeID string `json:"native_id,omitempty"`
			Type     string `json:"type"` // e.g. aws|azure|google|scm|cli|tfc
		} `json:"environment"`
		Resource struct {
			IacMappingsCount int               `json:"iac_mappings_count,omitempty"`
			ID               string            `json:"id,omitempty"`
			InputType        string            `json:"input_type"` // e.g. cloud_scan|tf|k8s...
			Location         string            `json:"location"`
			Name             string            `json:"name,omitempty"`
			NativeID         string            `json:"native_id,omitempty"`
			Platform         string            `json:"platform,omitempty"`
			ResourceType     string            `json:"resource_type,omitempty"`
			Tags             map[string]string `json:"tags,omitempty"`
			Type             string            `json:"type"` // e.g. cloud|iac
		} `json:"resource"`
	} `json:"cloud_resource,omitempty"`

	// 4) Source location within a file
	SourceLocation *struct {
		File   string `json:"file"`
		Region *struct {
			Start struct {
				Column int `json:"column"`
				Line   int `json:"line"`
			} `json:"start"`
			End struct {
				Column int `json:"column"`
				Line   int `json:"line"`
			} `json:"end"`
		} `json:"region,omitempty"`
	} `json:"sourceLocation,omitempty"`
}

// ExploitDetails => "maturity_levels" & "sources"
type ExploitDetails struct {
	MaturityLevels []struct {
		Format string `json:"format"`
		Level  string `json:"level"`
	} `json:"maturity_levels"`
	Sources []string `json:"sources"`
}

// Problem => source of vulnerability, policy, etc.
type Problem struct {
	DisclosedAt  *time.Time `json:"disclosed_at,omitempty"`
	DiscoveredAt *time.Time `json:"discovered_at,omitempty"`
	ID           string     `json:"id"`
	Source       string     `json:"source"`
	Type         string     `json:"type"` // rule|vulnerability
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	URL          *string    `json:"url,omitempty"`
}

type Resolution struct {
	Details    string    `json:"details,omitempty"`
	ResolvedAt time.Time `json:"resolved_at"`
	Type       string    `json:"type"` // disappeared|fixed
}

type Risk struct {
	Factors []RiskFactor `json:"factors"`
	Score   *RiskScore   `json:"score,omitempty"`
}

// RiskFactor => an item describing e.g. "has known exploit", "recently published", etc.
type RiskFactor struct {
	IncludedInScore bool       `json:"included_in_score"`
	Links           *RiskLinks `json:"links,omitempty"`
	Name            string     `json:"name"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Value           bool       `json:"value"`
}

// RiskLinks => e.g. evidence link
type RiskLinks struct {
	Evidence *EvidenceLink `json:"evidence,omitempty"`
}

// EvidenceLink can be a string or object
type EvidenceLink struct {
	Href *string                `json:"href,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// RiskScore => numeric risk scoring
type RiskScore struct {
	Model     string    `json:"model"`
	UpdatedAt time.Time `json:"updated_at"`
	Value     int       `json:"value"` // range 0..1000
}

type Severity struct {
	Level            string    `json:"level"`
	ModificationTime time.Time `json:"modification_time"`
	Score            float64   `json:"score"`
	Source           string    `json:"source"`
	Vector           string    `json:"vector"`
	Version          string    `json:"version"`
}
