package main

import (
	"FoodService/module_cmd"
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Print("Miss argrument")
	} else {
		switch argsWithoutProg[0] {
		case "api":
			err := module_cmd.RunAPI()
			if err != nil {
				fmt.Print(err.Error())
			}
		case "consumer":
			err := module_cmd.RunConsumer()
			if err != nil {
				fmt.Print(err.Error())
			}
		}
	}
}
