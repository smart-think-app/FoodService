package router

import (
	"FoodService/controller/drink_controller"
	"github.com/labstack/echo/v4"
)

func DrinkRouter(e *echo.Echo) {
	drinkRouter := e.Group("/drink")
	{
		drinkRouter.GET("/", drink_controller.SearchDrinkController)
		drinkRouter.GET("/:id", drink_controller.GetDetailDrinkController)
		drinkRouter.POST("/", drink_controller.AddDrinkController)
		drinkRouter.POST("/recipe", drink_controller.AddDrinkRecipeController)
	}
}
