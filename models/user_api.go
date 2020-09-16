package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (

	// UserCreatePayload ...
	UserCreatePayload struct {
		Name string `json:"name"`
	}
)

// Validate UserCreatePayload
func (payload UserCreatePayload) Validate() error {
	err := validation.Errors{
		"name": validation.Validate(payload.Name, validation.Required, validation.Length(3, 20), is.Alpha),
	}.Filter()
	return err
}

// ConvertToBSON ....
func (payload UserCreatePayload) ConvertToBSON() UserBSON{
	result := UserBSON{
		ID : primitive.NewObjectID(),
		Name :payload.Name,
		CreatedAt:time.Now(),
	}
	return result
}