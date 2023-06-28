package main

import (
	"github.com/if3chi/go-api/pkg/kernel"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No env file found")
	}

	app := kernel.Boot()

	go func() {
		app.Run()
	}()

	app.ListenForShutdown()
}
