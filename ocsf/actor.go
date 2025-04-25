package ocsf

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// ActorFields defines the Arrow fields for Actor.
var ActorFields = []arrow.Field{
	{Name: "app_name", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "app_uid", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "authorizations", Type: arrow.ListOf(AuthorizationStruct), Nullable: true},
	{Name: "idp", Type: IdentityProviderStruct, Nullable: true},
	{Name: "invoked_by", Type: arrow.BinaryTypes.String, Nullable: true},
	{Name: "process", Type: ProcessStruct, Nullable: true},
	{Name: "session", Type: SessionStruct, Nullable: true},
	{Name: "user", Type: UserStruct, Nullable: true},
}

var ActorStruct = arrow.StructOf(ActorFields...)
var ActorClassname = "actor"

type Actor struct {
	AppName        *string           `json:"app_name" parquet:"app_name,optional" ch:"app_name"`
	AppUID         *string           `json:"app_uid" parquet:"app_uid,optional" ch:"app_uid"`
	Authorizations []*Authorization  `json:"authorizations" parquet:"authorizations,list,optional" ch:"authorizations"`
	IDP            *IdentityProvider `json:"idp" parquet:"idp,optional" ch:"idp"`
	InvokedBy      *string           `json:"invoked_by" parquet:"invoked_by,optional" ch:"invoked_by"`
	Process        *Process          `json:"process" parquet:"process,optional" ch:"process"`
	Session        *Session          `json:"session" parquet:"session,optional" ch:"session"`
	User           *User             `json:"user" parquet:"user,optional" ch:"user"`
}
