package router

import (
	"FoodService/controller/food_controller"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func FoodRouter(db *sql.DB, e *echo.Echo) {
	foodCls := food_controller.FoodControllerClass{Db: db}
	foodRouter := e.Group("/food")
	{
		foodRouter.GET("/", foodCls.SearchFoodController)
		foodRouter.GET("/:id", foodCls.GetDetailFoodController)
		foodRouter.POST("/", foodCls.AddFoodController)
		foodRouter.POST("/recipe", foodCls.AddFoodRecipeController)
	}
}
