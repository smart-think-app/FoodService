package main

import (
	"FoodService/model/common_model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Hello World",
	})
}
