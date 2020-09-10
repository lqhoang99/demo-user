package validations

import (
	"github.com/labstack/echo/v4"

	"demo-user/models"
	"demo-user/util"
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
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", payload)
		return next(c)
	}
}