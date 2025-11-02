package main

import (
	database "store-api/database/connection"
	"store-api/database/setup"
	"store-api/router/server"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		println(".env not found")
	}

	database.ConnectDb()

	setup.CreatedAdmin()

	server.StartServer().Listen(":3000")

}
