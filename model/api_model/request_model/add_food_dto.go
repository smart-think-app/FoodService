package request_model

type AddFoodBodyRequestDto struct {
	Name        string                `json:"name"`
	TypeFood    int                   `json:"type_food"`
	Status      int                   `json:"status"`
	Description string                `json:"description"`
	Recipes     []AddRecipeRequestDto `json:"recipes"`
}

type AddRecipeRequestDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keyword     []string `json:"keyword"`
	Price       float64  `json:"price"`
	Level       int      `json:"level"`
	Images      []string `json:"images"`
}
