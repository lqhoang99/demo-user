package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-user/models"
	"demo-user/services"
	"demo-user/utils"
)

// UserList ...
func UserList(c echo.Context) error {
	// Process data
	rawData, err := services.UserList()

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, rawData, "")
}

// UserCreate ...
func UserCreate(c echo.Context) error {
	var (
		payload = c.Get("payload").(models.UserCreatePayload)
	)

	// Process data
	rawData, err := services.UserCreate(payload)

	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")

}

// TransactionFindByUserID ...
func TransactionFindByUserID(c echo.Context) error {
	var (
		userID          = c.Param("id")
	)

	// Process data
	rawData, err := services.TransactionFindByUserID(userID)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, rawData, "")
}