package validations

import (
	"demo-user/dao"
	"github.com/labstack/echo/v4"

	"demo-user/models"
	"demo-user/utils"
)

// UserCreate ...
func UserCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload models.UserCreatePayload
		)

		// ValidateStruct
		c.Bind(&payload)
		err := payload.Validate()

		//if err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("payload", payload)
		return next(c)
	}
}

// UserValidateID ...
func UserValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id          = c.Param("id")
			userID, err = utils.ValidationObjectID(id)
		)

		// if err
		if err != nil {
			return utils.Response400(c, nil, "id khong hop le")
		}

		// Check exitsted user
		user, _ := dao.UserFindByID(userID)
		if user.ID.IsZero() {
			return utils.Response404(c, nil, "not found user")
		}

		// Success
		c.Set("userID", userID)
		return next(c)
	}
}
