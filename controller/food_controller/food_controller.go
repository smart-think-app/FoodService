package food_controller

import (
	"FoodService/model/api_model/request_model"
	"FoodService/model/common_model"
	"FoodService/repository/food_repository"
	"FoodService/repository/recipe_repository"
	"FoodService/service/food_service"
	"FoodService/service/recipe_service"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	recipeRepository := recipe_repository.NewRecipeRepository(cls.Db)
	foodRepository := food_repository.NewFoodRepository(cls.Db)
	foodSv := food_service.NewFoodService(foodRepository, recipeRepository)
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

	foodIdParam := c.Param("id")
	foodId, errParse := strconv.Atoi(foodIdParam)
	if errParse != nil {
		return c.JSON(http.StatusBadRequest, common_model.SuccessResponseDto{
			Code: http.StatusBadRequest,
			Data: "Invalid Food",
		})
	}
	var request request_model.AddRecipeRequestDto
	if errBody := c.Bind(&request); errBody != nil {
		return c.JSON(http.StatusBadRequest, common_model.SuccessResponseDto{
			Code: http.StatusBadRequest,
			Data: "Invalid Body",
		})
	}
	RecipeRepository := recipe_repository.NewRecipeRepository(cls.Db)
	FoodRepository := food_repository.NewFoodRepository(cls.Db)
	RecipeService := recipe_service.NewRecipeService(RecipeRepository, FoodRepository)

	err := RecipeService.AddRecipeByFoodId(foodId, request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common_model.SuccessResponseDto{
			Code: http.StatusInternalServerError,
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, common_model.SuccessResponseDto{
		Code: http.StatusOK,
		Data: "Add Food Recipe",
	})
}
