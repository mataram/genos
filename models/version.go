package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Version struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Output    string    `json:"output" db:"output"`
	Schema    string    `json:"schema" db:"schema"`
	Status    int       `json:"status" db:"status"`
	EventID   uuid.UUID `json:"event_id" db:"event_id"`
	Event     Event     `belongs_to:"event"`
}

func (e Version) GetBreadcumbs() []Breadcrumb {
	breadcrumbs := []Breadcrumb{}
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/", "Home"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services", "Services"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.Event.ServiceID.String(), e.Event.ServiceID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Events"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.Event.ServiceID.String() + "/events/" + e.EventID.String(), e.EventID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Versions"})
	return breadcrumbs

}

// String is not required by pop and may be deleted
func (v Version) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Versions is not required by pop and may be deleted
type Versions []Version

// String is not required by pop and may be deleted
func (v Versions) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *Version) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Name, Name: "Name"},
		&validators.StringIsPresent{Field: v.Output, Name: "Output"},
		&validators.StringIsPresent{Field: v.Schema, Name: "Schema"},
		&validators.IntIsPresent{Field: v.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
