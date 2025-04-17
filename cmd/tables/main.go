package main

import (
	"fullsteak/internal/database"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	err := database.InitTables(db, "tables.sql")
	if err != nil {
		log.Fatal(err)
	}
}
