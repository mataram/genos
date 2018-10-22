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

type Event struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Name        string       `json:"name" db:"name"`
	ServiceID   uuid.UUID    `json:"service_id" db:"service_id"`
	Description nulls.String `json:"description" db:"description"`
	Status      int          `json:"status" db:"status"`
	Service     Service      `belongs_to:"service"`
	Versions    Versions     `has_many:"versions" order_by:"name asc"`
}

func (e Event) GetBreadcumbs() []Breadcrumb {
	breadcrumbs := []Breadcrumb{}
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/", "Home"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services", "Service"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services/" + e.ServiceID.String(), e.ServiceID.String()})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"#", "Events"})

	return breadcrumbs

}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.IntIsPresent{Field: e.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
