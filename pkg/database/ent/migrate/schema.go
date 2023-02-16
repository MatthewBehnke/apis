// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorizationPoliciesColumns holds the columns for the "authorization_policies" table.
	AuthorizationPoliciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ptype", Type: field.TypeString, Default: ""},
		{Name: "v0", Type: field.TypeString, Default: ""},
		{Name: "v1", Type: field.TypeString, Default: ""},
		{Name: "v2", Type: field.TypeString, Default: ""},
		{Name: "v3", Type: field.TypeString, Default: ""},
		{Name: "v4", Type: field.TypeString, Default: ""},
		{Name: "v5", Type: field.TypeString, Default: ""},
	}
	// AuthorizationPoliciesTable holds the schema information for the "authorization_policies" table.
	AuthorizationPoliciesTable = &schema.Table{
		Name:       "authorization_policies",
		Columns:    AuthorizationPoliciesColumns,
		PrimaryKey: []*schema.Column{AuthorizationPoliciesColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "data", Type: field.TypeBytes},
		{Name: "expiry", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString, Nullable: true},
		{Name: "attempt_count", Type: field.TypeInt, Nullable: true},
		{Name: "last_attempt", Type: field.TypeTime, Nullable: true},
		{Name: "locked", Type: field.TypeTime, Nullable: true},
		{Name: "role", Type: field.TypeString, Default: "user"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorizationPoliciesTable,
		SessionsTable,
		UsersTable,
	}
)

func init() {
}
