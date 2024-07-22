package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Init(){
	connStr:= "user=username dbname=lms_project sslmode=disable password=password"
	// var err error

	DB,err := sql.Open("postgres",connStr)
	if err!=nil{
		log.Fatal(err)
	}

	if err := DB.Ping(); err!= nil{
		log.Fatal(err)
	}

	fmt.Println("Successfull connection to the DB")
}


