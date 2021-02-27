package recipe_service

import (
	"FoodService/interface/repository"
	"FoodService/model/api_model/request_model"
	"FoodService/model/schema_model"
	"errors"
	"strings"
)

type iRecipeService struct {
	IRecipeRepository repository.IRecipeRepository
	IFoodRepository   repository.IFoodRepository
}

func NewRecipeService(recipeRepository repository.IRecipeRepository, foodRepository repository.IFoodRepository) *iRecipeService {
	return &iRecipeService{
		IRecipeRepository: recipeRepository,
		IFoodRepository:   foodRepository,
	}
}

func (sv *iRecipeService) AddRecipeByFoodId(foodId int, addRecipeRequest request_model.AddRecipeRequestDto) error {
	isCheckFood, errCheckFood := sv.IFoodRepository.CheckById(foodId)
	if errCheckFood != nil {
		return errCheckFood
	}
	if isCheckFood == false {
		return errors.New("Not Found Food. ")
	}

	err := sv.IRecipeRepository.Add(schema_model.RecipeSchemaModel{
		FoodId:      foodId,
		Name:        addRecipeRequest.Name,
		Description: addRecipeRequest.Description,
		Keyword:     addRecipeRequest.Keyword,
		Price:       addRecipeRequest.Price,
		Level:       addRecipeRequest.Level,
		Images:      strings.Join(addRecipeRequest.Images, ","),
	})
	if err != nil {
		return err
	}
	return nil
}
