package main

import (
	"fmt"
	"fullsteak/internal/app"
)

func main() {
	server := app.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
