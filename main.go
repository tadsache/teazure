package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main.go/internal/tui"
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

	// Load THeme
	err = tui.LoadTheme()
	if err != nil {
		return err
	}
	tui.LoadStyles() // error Hanlding?

	fmt.Println(tui.GlobalTheme.SelectionBackground)
	return app.Start()
}
