package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-user/models"
)

func userCreatePayloadToBSON(body models.UserCreatePayload) models.UserBSON {
	result := models.UserBSON{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		CreatedAt: time.Now(),
	}
	return result
}

func convertToUserDetail(doc models.UserBSON) models.UserDetail {
	result := models.UserDetail{
		ID:               doc.ID,
		Name:             doc.Name,
		TotalTransaction: doc.TotalTransaction,
		TotalCommission:  doc.TotalCommission,
		CreatedAt:        doc.CreatedAt,
		UpdatedAt:        doc.UpdatedAt,
	}
	return result
}
