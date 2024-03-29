package postgres_provider

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectPostgres() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost/FoodService?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("Connect Postgres Success")
	return db
}
