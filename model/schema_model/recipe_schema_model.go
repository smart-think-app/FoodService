package schema_model

type RecipeSchemaModel struct {
	Id          int      `db:"id"`
	FoodId      int      `db:"food_id"`
	Name        string   `db:"name"`
	Description string   `db:"description"`
	Keyword     []string `db:"keyword"`
	Price       float64  `db:"price"`
	Level       int      `db:"level"`
	Images      string   `db:"images"`
}
