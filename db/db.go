package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	// https://pkg.go.dev/github.com/lib/pq#section-readme
	// Abre uma conexão com o banco de dados

	psqlInfo := "user=postgres dbname=alura_store password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
