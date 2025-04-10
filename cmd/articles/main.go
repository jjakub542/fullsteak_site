package main

import (
	"fullsteak/internal/database"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	err := database.InsertExampleArticles(db)
	if err != nil {
		log.Fatal(err)
	}
}
