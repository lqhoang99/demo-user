package services

import (
	"demo-user/models"
)

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
