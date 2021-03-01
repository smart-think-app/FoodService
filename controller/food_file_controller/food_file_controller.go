package food_file_controller

import (
	"FoodService/model/common_model"
	"FoodService/provider/rabbitmq_provider"
	"FoodService/repository/food_repository"
	"FoodService/service/food_file_service"
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"net/http"
)

type FoodFileControllerCls struct {
	Db       *sql.DB
	RabbitMQ *amqp.Channel
}

func (cls FoodFileControllerCls) AddFoodByFileExcelController(c echo.Context) error {
	NewFoodRepository := food_repository.NewFoodRepository(cls.Db)
	RabbitMQSupport := rabbitmq_provider.RabbitMQSupport{
		Ch: cls.RabbitMQ,
	}
	FoodFileService := food_file_service.NewFoodFileService(NewFoodRepository, RabbitMQSupport)
	err := FoodFileService.AddFoodByFileExcel(c)
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
