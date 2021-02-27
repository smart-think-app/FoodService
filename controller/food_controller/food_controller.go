package food_controller

import (
	"FoodService/model/api_model/request_model"
	"FoodService/model/common_model"
	"FoodService/repository/food_repository"
	"FoodService/service/food_service"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FoodControllerClass struct {
	Db *sql.DB
}

func (cls FoodControllerClass) SearchFoodController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Search Food",
	})
}

func (cls FoodControllerClass) GetDetailFoodController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Get Detail Food",
	})
}

func (cls FoodControllerClass) AddFoodController(c echo.Context) error {
	var request request_model.AddFoodBodyRequestDto

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, common_model.ErrorResponseDto{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
	}
	foodRepository := food_repository.NewFoodRepository(cls.Db)
	foodSv := food_service.NewFoodService(foodRepository)
	if err := foodSv.AddFoodService(request); err != nil {
		return c.JSON(http.StatusInternalServerError, common_model.ErrorResponseDto{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food",
	})
}

func (cls FoodControllerClass) AddFoodRecipeController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food Recipe",
	})
}
