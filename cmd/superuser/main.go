package main

import (
	"fmt"
	"fullsteak/internal/database"
	"fullsteak/internal/domain"
	"fullsteak/internal/repository"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := database.Connect()
	var err error
	newUser := domain.User{IsSuperuser: true}

	fmt.Println("Email: ")
	fmt.Scanln(&newUser.Email)
	fmt.Println("Password: ")
	fmt.Scanln(&newUser.Password)

	err = newUser.Validate()
	if err != nil {
		log.Fatal(err)
	}

	newUser.CreatePasswordHash()

	repo := repository.New(db)

	err = repo.User.CreateOne(&newUser)

	if err != nil {
		log.Fatal(err)
	}
}
