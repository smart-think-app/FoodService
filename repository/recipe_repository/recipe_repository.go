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

func (r *recipeRepository) AddMany(recipe []schema_model.RecipeSchemaModel) error {
	sql_arr := make([]string, 0)
	values := make([]interface{}, 0)
	for index, item := range recipe {
		loopInc := 7 * index
		sql_arr = append(sql_arr, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			1+loopInc, 2+loopInc, 3+loopInc, 4+loopInc, 5+loopInc, 6+loopInc, 7+loopInc))
		values = append(values,
			item.FoodId,
			item.Name,
			item.Description,
			fmt.Sprintf("{%s}", strings.Join(item.Keyword, ",")),
			item.Price,
			item.Level,
			item.Images)
	}
	sqlQuery := fmt.Sprintf("INSERT INTO \"Recipe\"(food_id, name, description, keyword, price, level, images) \nVALUES %s",
		strings.Join(sql_arr, ","))
	rows := r.Db.QueryRow(sqlQuery, values...)

	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}
