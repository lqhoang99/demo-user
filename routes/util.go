package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-user/dao"
	"demo-user/utils"
)

func userCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			userID = c.Get("userID").(primitive.ObjectID)
		)

		// Find
		user, _ := dao.UserFindByID(userID)

		// Check existed
		if user.ID.IsZero() {
			return utils.Response404(c, nil, "Not found user")
		}
		return next(c)
	}

}
