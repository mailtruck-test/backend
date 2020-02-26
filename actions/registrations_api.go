package actions

import (
	"fmt"
	"regexp"
	"strings"
	"userreg/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/gobuffalo/validate"
)

// Custom API resource to handle api registration requests
func (v RegistrationsResource) POST(c buffalo.Context) error {
	// Allocate an empty Registration
	registration := &models.Registration{}

	fmt.Println(c.Request().Body)

	// Bind registration to the html form elements
	if err := c.Bind(registration); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}
	errors := validate.Validate(&ZipValidator{"Zip", registration.Zip}, &CountryValidator{"Country", registration.Country}, &StateValidator{"State", registration.State})
	if errors.HasAny() {
		return c.Render(422, r.JSON(errors))
	}

	verrs, err := tx.ValidateAndCreate(registration)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Render(422, r.JSON(verrs))
	}

	c.Flash().Add("success", T.Translate(c, "registration.created.success"))

	return c.Render(201, r.JSON(registration))
}

type ZipValidator struct {
	Field string
	Value string
}

func (z *ZipValidator) IsValid(errors *validate.Errors) {
	exp := regexp.MustCompile(`(^\d{5}$)|(^\d{9}$)|(^\d{5}-\d{4}$)`)
	match := exp.FindString(z.Value)
	if z.Value != match {
		errors.Add(strings.ToLower(z.Field), fmt.Sprintf("%s invalid format!", z.Field))
	}
}

type CountryValidator struct {
	Field string
	Value string
}

func (c *CountryValidator) IsValid(errors *validate.Errors) {
	if c.Value != "US" {
		errors.Add(strings.ToLower(c.Field), fmt.Sprintf("%s Invalid!", c.Field))
	}
}

type StateValidator struct {
	Field string
	Value string
}

func (s *StateValidator) IsValid(errors *validate.Errors) {
	_, ok := StatesMap[s.Value]
	if !ok {
		errors.Add(strings.ToLower(s.Field), fmt.Sprintf("%s Invalid!", s.Field))
	}
}
