package main

import (
	"github.com/if3chi/go-api/pkg/kernel"
	"github.com/if3chi/go-api/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No env file found")
	}

	app := kernel.Boot()

	routes.Load(app)

	go func() {
		app.Run()
	}()

	app.ListenForShutdown()
}
