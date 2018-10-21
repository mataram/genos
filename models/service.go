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

type Service struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Prefix      string       `json:"prefix" db:"prefix"`
	Name        string       `json:"name" db:"name"`
	Description nulls.String `json:"description" db:"description"`
	Status      int          `json:"status" db:"status"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Events      Events       `has_many:"events" order_by:"name asc"`
}

func (s Service) GetBreadcumbs() []Breadcrumb {
	breadcrumbs := []Breadcrumb{}
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/", "Home"})
	breadcrumbs = append(breadcrumbs, Breadcrumb{"/services", "Services"})
	return breadcrumbs
}

// String is not required by pop and may be deleted
func (s Service) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Services is not required by pop and may be deleted
type Services []Service

// String is not required by pop and may be deleted
func (s Services) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Service) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Service) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Service) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
