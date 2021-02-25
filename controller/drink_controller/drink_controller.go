package drink_controller

import (
	"FoodService/model/common_model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SearchDrinkController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Search Drink",
	})
}

func GetDetailDrinkController(c echo.Context) error {
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Get Detail Drink",
	})
}
