package database

import (
	"WebStore/config"
	"database/sql"
	"fmt"
)

func Connect() *sql.DB {

	connectionString := config.GetEnvironmentVariables("DATABASE_CONNECTION")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connected!")
	return db
}
