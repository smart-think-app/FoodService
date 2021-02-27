package router

import (
	"FoodService/controller/food_controller"
	"FoodService/controller/food_file_controller"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func FoodRouter(db *sql.DB, e *echo.Echo) {
	foodCls := food_controller.FoodControllerClass{Db: db}
	foodFileCls := food_file_controller.FoodFileControllerCls{Db: db}

	foodRouter := e.Group("/food")
	{
		foodRouter.GET("/", foodCls.SearchFoodController)
		foodRouter.GET("/:id", foodCls.GetDetailFoodController)
		foodRouter.POST("/", foodCls.AddFoodController)
		foodRouter.POST("/:id/recipe", foodCls.AddFoodRecipeController)

	}

	foodFileRouter := e.Group("/food/file")
	{
		foodFileRouter.POST("/", foodFileCls.AddFoodByFileExcelController)
	}

}
