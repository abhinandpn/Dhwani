package main

import (
	"log"

	"github.com/abhinandpn/Dhwani/internal/app"
)

func main() {
	myApp, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	myApp.Start()
}
