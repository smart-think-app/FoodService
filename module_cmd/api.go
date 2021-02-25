package module_cmd

import (
	"FoodService/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunAPI() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	router.FoodRouter(e)
	router.DrinkRouter(e)
	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
