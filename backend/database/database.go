package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, dbPassword)
	// var err error

	DB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfull connection to the DB")
}
