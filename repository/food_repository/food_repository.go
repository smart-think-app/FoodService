package food_repository

import (
	"FoodService/model/schema_model"
	"database/sql"
	"errors"
)

type foodRepository struct {
	Db *sql.DB
}

func NewFoodRepository(db *sql.DB) *foodRepository {
	return &foodRepository{Db: db}
}

func (r *foodRepository) Add(foodSchema schema_model.FoodSchemaModel) (int, error) {
	foodId := 0
	err := r.Db.QueryRow(`Insert Into "Food"(name , type_food ,status ,description , updated_date) values ($1,$2,$3,$4,$5) RETURNING Id`,
		foodSchema.Name, foodSchema.TypeFood, foodSchema.Status, foodSchema.Description, foodSchema.UpdatedDate).Scan(&foodId)
	if err != nil {
		return 0, err
	}
	if foodId == 0 {
		return 0, errors.New("Insert Food Fail. ")
	}

	return foodId, nil
}

func (r *foodRepository) CheckById(foodId int) (bool, error) {

	var count int64
	err := r.Db.QueryRow("select count(id) from \"Food\" where id = $1", foodId).Scan(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *foodRepository) AddMany(foods []schema_model.FoodSchemaModel) error {
	return nil
}
