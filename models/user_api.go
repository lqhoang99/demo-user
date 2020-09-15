package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	// UserCreatePayload ...
	UserCreatePayload struct {
		Name string `json:"name" valid:"stringlength(3|30),type(string)"`
	}
)

// Validate UserCreatePayload
func (payload UserCreatePayload) Validate() error {
	err := validation.Errors{
		"name": validation.Validate(payload.Name, validation.Required, validation.Length(3, 20), is.Alpha),
	}.Filter()
	return err
}
