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
	AllowNull   bool         `json:"allow_null" db:"allow_null"`
	Description nulls.String `json:"description" db:"description"`
	Index       int          `json:"index" db:"index"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Version     Version      `belongs_to:"version"`
}

func (e Field) GetBreadcumbs() []Breadcrumb {
	breadcrumbs := []Breadcrumb{}
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/", "Home"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services", "Services"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.Version.Event.ServiceID.String(), e.Version.Event.ServiceID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Events"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.Version.Event.ServiceID.String() + "/events/" + e.Version.EventID.String(), e.Version.EventID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Versions"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.Version.Event.ServiceID.String() + "/events/" + e.Version.EventID.String() + "/versions/" + e.VersionID.String(), e.Version.ID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Fields"})
	return breadcrumbs

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
