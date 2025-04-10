package main

import (
	"fullsteak/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	database.DropTables(db)
}
