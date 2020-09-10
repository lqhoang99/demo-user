package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (

	// UserBSON ...
	UserBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		Name             string             `bson:"name"`
		TotalTransaction int64                `bson:"totalTransaction"`
		TotalCommission  float64            `bson:"totalCommission"`
		CreatedAt        time.Time          `bson:"createdAt"`
		UpdatedAt        time.Time          `bson:"updatedAt"`
	}

	// UserDetail ...
	UserDetail struct {
		ID               primitive.ObjectID `json:"_id"`
		Name             string             `json:"name"`
		TotalTransaction int64                `json:"totalTransaction"`
		TotalCommission  float64            `json:"totalCommission"`
		CreatedAt        time.Time          `json:"createdAt"`
		UpdatedAt        time.Time          `json:"updatedAt"`
	}
)
