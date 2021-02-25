package food_controller

import (
	"FoodService/model/common_model"
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
