package main

import (
	"FoodService/module_cmd"
	"fmt"
)

func main() {
	err := module_cmd.RunAPI()
	if err != nil {
		fmt.Print(err.Error())
	}
}
