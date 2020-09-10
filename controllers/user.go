package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-user/models"
	"demo-user/services"
	"demo-user/util"
)

// UserList ...
func UserList(c echo.Context) error {
	// Process data
	rawData, err := services.UserList()

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}

// UserCreate ...
func UserCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserCreatePayload)
	)

	//Process data
	rawData, err := services.UserCreate(body)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")

}

// TransactionFindByUserID ...
func TransactionFindByUserID(c echo.Context) error {
	var (
		userID          = c.Param("id")
	)

	// process data
	rawData, err := services.TransactionFindByUserID(userID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, rawData, "")
}