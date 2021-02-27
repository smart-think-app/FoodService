package repository

import "FoodService/model/schema_model"

type IRecipeRepository interface {
	Add(recipe schema_model.RecipeSchemaModel) error
}
