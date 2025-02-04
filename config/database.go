package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=12345 host=localhost")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")

	return db
}
