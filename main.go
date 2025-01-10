package main

import (
	"github.com/joho/godotenv"
	"log"
	"main.go/internal/tui/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	// Load .env file
	// maybe change or add config file
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("Error loading .env file")
		return err
	}

	return app.Start()
}
