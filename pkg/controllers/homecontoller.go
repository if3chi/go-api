package controllers

import (
	"net/http"

	"github.com/if3chi/go-api/pkg/kernel"
)

func HandleApiRoot(app *kernel.Application) http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(rw http.ResponseWriter, req *http.Request) {
		app.Respond(rw, req, &response{Message: "You are viewing " + app.Config.App.Name}, http.StatusOK)
	}
}
