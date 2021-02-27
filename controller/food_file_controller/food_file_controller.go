package food_file_controller

import (
	"FoodService/model/common_model"
	"FoodService/repository/food_repository"
	"FoodService/service/food_file_service"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FoodFileControllerCls struct {
	Db *sql.DB
}

func (cls FoodFileControllerCls) AddFoodByFileExcelController(c echo.Context) error {
	NewFoodRepository := food_repository.NewFoodRepository(cls.Db)
	FoodFileService := food_file_service.NewFoodFileService(NewFoodRepository)
	err := FoodFileService.AddFoodByFileExcel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common_model.ErrorResponseDto{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food By File",
	})
}
