package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Field struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	VersionID   uuid.UUID    `json:"version_id" db:"version_id"`
	Name        string       `json:"name" db:"name"`
	Type        string       `json:"type" db:"type"`
	AllowNull   string       `json:"allow_null" db:"allow_null"`
	Description nulls.String `json:"description" db:"description"`
	Index       int          `json:"index" db:"index"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (f Field) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Fields is not required by pop and may be deleted
type Fields []Field

// String is not required by pop and may be deleted
func (f Fields) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Field) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Name, Name: "Name"},
		&validators.StringIsPresent{Field: f.Type, Name: "Type"},
		&validators.StringIsPresent{Field: f.AllowNull, Name: "AllowNull"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Field) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Field) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
