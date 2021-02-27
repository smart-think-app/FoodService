package food_status_enum

type FoodStatusEnum struct {
	Value   int
	Name    string
	Display string
}

func Available() FoodStatusEnum {
	return FoodStatusEnum{
		Value:   1,
		Name:    "Available",
		Display: "Available",
	}
}

func Closed() FoodStatusEnum {
	return FoodStatusEnum{
		Value:   2,
		Name:    "Closed",
		Display: "Closed",
	}
}
