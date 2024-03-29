package request_model

type AddRecipeRequestDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keyword     []string `json:"keyword"`
	Price       float64  `json:"price"`
	Level       int      `json:"level"`
	Images      []string `json:"images"`
}
