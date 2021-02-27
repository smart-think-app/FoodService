package repository

import "FoodService/model/schema_model"

type IFoodRepository interface {
	Add(foodSchema schema_model.FoodSchemaModel) (int, error)
	CheckById(foodId int) (bool, error)
	AddMany(foods []schema_model.FoodSchemaModel) error
}
