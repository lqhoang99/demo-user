package routes

import (
	"github.com/labstack/echo/v4"

	"demo-user/controllers"
	"demo-user/validations"
)

// User ...
func User(e *echo.Echo) {
	routes := e.Group("/users")

	routes.GET("", controllers.UserList)
	routes.POST("", controllers.UserCreate, validations.UserCreate)
}
