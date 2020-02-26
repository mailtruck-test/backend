package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Registration is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Registration struct {
    ID uuid.UUID `json:"id" db:"id"`
    First string `json:"first" db:"first"`
    Last string `json:"last" db:"last"`
    Address string `json:"address" db:"address"`
    Address2 nulls.String `json:"address2" db:"address2"`
    City string `json:"city" db:"city"`
    State string `json:"state" db:"state"`
    Zip string `json:"zip" db:"zip"`
    Country string `json:"country" db:"country"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Registration) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Registrations is not required by pop and may be deleted
type Registrations []Registration

// String is not required by pop and may be deleted
func (r Registrations) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Registration) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.First, Name: "First"},
		&validators.StringIsPresent{Field: r.Last, Name: "Last"},
		&validators.StringIsPresent{Field: r.Address, Name: "Address"},
		&validators.StringIsPresent{Field: r.City, Name: "City"},
		&validators.StringIsPresent{Field: r.State, Name: "State"},
		&validators.StringIsPresent{Field: r.Zip, Name: "Zip"},
		&validators.StringIsPresent{Field: r.Country, Name: "Country"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Registration) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Registration) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
