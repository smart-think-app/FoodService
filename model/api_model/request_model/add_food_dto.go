package request_model

type AddFoodBodyRequestDto struct {
	Name        string                `json:"name"`
	TypeFood    int                   `json:"type_food"`
	Status      int                   `json:"status"`
	Description string                `json:"description"`
	Recipes     []AddRecipeRequestDto `json:"recipes"`
}
