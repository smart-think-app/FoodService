package module_cmd

import (
	"FoodService/provider/postgres_provider"
	"FoodService/router"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db *sql.DB
)

func RunAPI() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db = postgres_provider.ConnectPostgres()
	if db == nil {
		fmt.Print("Start Fail")
	} else {
		// Routes
		router.FoodRouter(db, e)
		router.DrinkRouter(e)
		// Start server
		e.Logger.Fatal(e.Start(":3001"))

	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print("End!")
	}()
}
