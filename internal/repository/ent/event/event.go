// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// Table holds the table name of the event in the database.
	Table = "events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldName,
	FieldDescription,
	FieldLocation,
	FieldDate,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
