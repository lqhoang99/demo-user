package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// UserCreatePayload ...
	UserCreatePayload struct {
		Name string `json:"name"`
	}
)

// Validate UserCreatePayload
func (payload UserCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("name is required"),
			validation.Length(3,30).Error("name is length: 3 -> 30"),
			is.Alpha.Error("name is alpha"),
		),
	)
}

// ConvertToBSON ....
func (payload UserCreatePayload) ConvertToBSON() UserBSON {
	result := UserBSON{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return result
}
