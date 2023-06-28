package routes

import (
	"net/http"

	"github.com/if3chi/go-api/pkg/kernel"
)

func Load(app *kernel.Application) {
	router := app.Router.Methods(http.MethodGet).Subrouter()

	router.HandleFunc("/", ApiRoot(app)).Name("api:root")
}

func ApiRoot(app *kernel.Application) http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}

	return func(rw http.ResponseWriter, req *http.Request) {
		app.Respond(rw, req, &response{Message: "You are viewing " + app.Config.App.Name}, http.StatusOK)
	}
}
