package db

import (
	"database/sql"
	"fmt"
)

func ConectaBanco() *sql.DB {
	conexao := "user=postgres dbname=teste password=12345 host=localhost sslmode=disable "
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		fmt.Println("Ocorreu um erro na conexão")
		panic(err.Error())
	}

	return db
}
