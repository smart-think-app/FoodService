package food_service

import (
	"FoodService/interface/repository"
	"FoodService/model/api_model/request_model"
	"FoodService/model/schema_model"
	"fmt"
	"strings"
	"time"
)

type FoodServiceModel struct {
	IFoodRepository   repository.IFoodRepository
	IRecipeRepository repository.IRecipeRepository
}

func NewFoodService(iFoodRepository repository.IFoodRepository, iRecipeRepository repository.IRecipeRepository) *FoodServiceModel {
	return &FoodServiceModel{
		IFoodRepository:   iFoodRepository,
		IRecipeRepository: iRecipeRepository,
	}
}

func (s *FoodServiceModel) AddFoodService(request request_model.AddFoodBodyRequestDto) error {
	foodId, err := s.IFoodRepository.Add(schema_model.FoodSchemaModel{
		Name:        request.Name,
		TypeFood:    request.TypeFood,
		Status:      request.Status,
		Description: request.Description,
		UpdatedDate: time.Now(),
	})

	if err != nil {
		return err
	}
	if len(request.Recipes) > 0 {
		fmt.Print(foodId)
		for _, item := range request.Recipes {
			errRecipe := s.IRecipeRepository.Add(schema_model.RecipeSchemaModel{
				FoodId:      foodId,
				Name:        item.Name,
				Description: item.Description,
				Keyword:     item.Keyword,
				Price:       item.Price,
				Level:       item.Level,
				Images:      strings.Join(item.Images, ","),
			})
			//TODO Push On Queue
			if errRecipe != nil {
				fmt.Print(errRecipe.Error())
			}
		}
	}
	return nil
}

func (s *FoodServiceModel) GetFoodService() error {
	return nil
}
