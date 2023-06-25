// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 24},
		{Name: "permissions", Type: field.TypeJSON, Nullable: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Size: 32},
		{Name: "password", Type: field.TypeString},
		{Name: "pfp", Type: field.TypeBytes, Size: 4000000},
		{Name: "name", Type: field.TypeString, Size: 32},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_roles_role",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RolesTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = RolesTable
}
