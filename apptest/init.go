package apptest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"demo-user/config"
	"demo-user/modules/database"
	"demo-user/modules/zookeeper"
	"demo-user/routes"
	"demo-user/utils"
)

// InitServer ...
func InitServer() *echo.Echo {
	config.InitENV()
	zookeeper.Connect()
	database.Connect()
	utils.HelperConnect()

	// New echo
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	// Route
	routes.Boostrap(e)
	return e
}
