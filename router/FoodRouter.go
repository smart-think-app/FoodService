package router

import (
	"FoodService/controller/food_controller"
	"github.com/labstack/echo/v4"
)

func FoodRouter(e *echo.Echo) {
	foodRouter := e.Group("/food")
	{
		foodRouter.GET("/", food_controller.SearchFoodController)
		foodRouter.GET("/:id", food_controller.GetDetailFoodController)
	}
}
