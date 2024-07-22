package main

import (
	"lms/database"
	"lms/routes"
)

func main() {
	database.Init()
	r := routes.SetupRouter()
	r.Run(":8008")
}
