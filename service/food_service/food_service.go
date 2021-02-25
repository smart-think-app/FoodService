package food_service

import (
	"FoodService/model/api_model/request_model"
	"fmt"
)

type FoodServiceModel struct {
}

func NewFoodService() *FoodServiceModel {
	return &FoodServiceModel{}
}

func (s *FoodServiceModel) AddFoodService(request request_model.AddFoodBodyRequestDto) error {
	fmt.Print(request)
	return nil
}

func (s *FoodServiceModel) GetFoodService() error {
	return nil
}
