package service

import "FoodService/model/api_model/request_model"

type IFoodServiceInterface interface {
	AddFoodService(request request_model.AddFoodBodyRequestDto) error
	GetFoodService() error
}
