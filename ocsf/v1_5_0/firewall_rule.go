// autogenerated by scripts/model_gen.go. DO NOT EDIT
package v1_5_0

import (
	"github.com/apache/arrow-go/v18/arrow"
)

type FirewallRule struct {

	// Category: The rule category.
	Category *string `json:"category,omitempty" parquet:"category,optional"`

	// Condition: The rule trigger condition for the rule. For example: SQL_INJECTION.
	Condition *string `json:"condition,omitempty" parquet:"condition,optional"`

	// Description: The description of the rule that generated the event.
	Desc *string `json:"desc,omitempty" parquet:"desc,optional"`

	// Duration Milliseconds: The rule response time duration, usually used for challenge completion time.
	Duration *int64 `json:"duration,omitempty" parquet:"duration,optional"`

	// Match Details: The data in a request that rule matched. For example: '["10","and","1"]'.
	MatchDetails []string `json:"match_details,omitempty" parquet:"match_details,optional,list"`

	// Match Location: The location of the matched data in the source which resulted in the triggered firewall rule. For example: HEADER.
	MatchLocation *string `json:"match_location,omitempty" parquet:"match_location,optional"`

	// Name: The name of the rule that generated the event.
	Name *string `json:"name,omitempty" parquet:"name,optional"`

	// Rate Limit: The rate limit for a rate-based rule.
	RateLimit *int32 `json:"rate_limit,omitempty" parquet:"rate_limit,optional"`

	// Sensitivity: The sensitivity of the firewall rule in the matched event. For example: HIGH.
	Sensitivity *string `json:"sensitivity,omitempty" parquet:"sensitivity,optional"`

	// Type: The rule type.
	Type *string `json:"type,omitempty" parquet:"type,optional"`

	// Unique ID: The unique identifier of the rule that generated the event.
	Uid *string `json:"uid,omitempty" parquet:"uid,optional"`

	// Version: The rule version. For example: <code>1.1</code>.
	Version *string `json:"version,omitempty" parquet:"version,optional"`
}

var FirewallRuleFields = []arrow.Field{
	{Name: "category", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "condition", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "desc", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "duration", Type: arrow.PrimitiveTypes.Int64, Nullable: true},
	{Name: "match_details", Type: arrow.ListOf(arrow.BinaryTypes.String), Nullable: true},
	{Name: "match_location", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "rate_limit", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
	{Name: "sensitivity", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "type", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "version", Type: arrow.BinaryTypes.String, Nullable: true},
}

var FirewallRuleStruct = arrow.StructOf(FirewallRuleFields...)

var FirewallRuleSchema = arrow.NewSchema(FirewallRuleFields, nil)
