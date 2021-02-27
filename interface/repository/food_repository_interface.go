package repository

import "FoodService/model/schema_model"

type IFoodRepository interface {
	Add(foodSchema schema_model.FoodSchemaModel) (int, error)
}
