package controllers

import (
	"net/http"

	"github.com/if3chi/go-api/pkg/database"
	"github.com/if3chi/go-api/pkg/kernel"
	"github.com/if3chi/go-api/pkg/models"
)

func HandleClientsList(app *kernel.Application) http.HandlerFunc {
	db := database.Connect(app)

	return func(rw http.ResponseWriter, req *http.Request) {
		var clients []models.Client

		db.Find(&clients)

		app.Respond(rw, req, clients, http.StatusOK)
	}
}
