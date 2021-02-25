package food_controller

import (
	"FoodService/model/api_model/request_model"
	"FoodService/model/common_model"
	"FoodService/service/food_service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SearchFoodController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Search Food",
	})
}

func GetDetailFoodController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Get Detail Food",
	})
}

func AddFoodController(c echo.Context) error {
	var request request_model.AddFoodBodyRequestDto

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, common_model.ErrorResponseDto{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
		})
	}

	foodSv := food_service.NewFoodService()
	if err := foodSv.AddFoodService(request); err != nil {
		return c.JSON(http.StatusInternalServerError, common_model.ErrorResponseDto{
			Code:    http.StatusInternalServerError,
			Message: "Internal Error",
		})
	}
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food",
	})
}

func AddFoodRecipeController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food Recipe",
	})
}
