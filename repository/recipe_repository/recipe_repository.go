package recipe_repository

import (
	"FoodService/model/schema_model"
	"database/sql"
	"fmt"
	"strings"
)

type recipeRepository struct {
	Db *sql.DB
}

func NewRecipeRepository(Db *sql.DB) *recipeRepository {
	return &recipeRepository{
		Db: Db,
	}
}

func (r *recipeRepository) Add(recipe schema_model.RecipeSchemaModel) error {
	rows := r.Db.QueryRow("INSERT INTO \"Recipe\"(food_id, name, description, keyword, price, level, images) \nVALUES ($1,$2,$3,$4,$5,$6,$7)",
		recipe.FoodId,
		recipe.Name,
		recipe.Description,
		fmt.Sprintf("{%s}", strings.Join(recipe.Keyword, ",")),
		recipe.Price,
		recipe.Level,
		recipe.Images)

	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

func (r *recipeRepository) AddMany(model []schema_model.RecipeSchemaModel) error {
	return nil
}
