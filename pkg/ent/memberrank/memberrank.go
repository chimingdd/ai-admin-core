// Code generated by ent, DO NOT EDIT.

package memberrank

import (
	"time"
)

const (
	// Label holds the string label denoting the memberrank type in the database.
	Label = "member_rank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// Table holds the table name of the memberrank in the database.
	Table = "core_mms_rank"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "core_mms_members"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "core_mms_members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "rank_id"
)

// Columns holds all SQL columns for memberrank fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCode,
	FieldDescription,
	FieldRemark,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
