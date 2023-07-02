package routes

import (
	"net/http"

	control "github.com/if3chi/go-api/pkg/controllers"
	"github.com/if3chi/go-api/pkg/kernel"
)

func Load(app *kernel.Application) {
	router := app.Router.Methods(http.MethodGet).Subrouter()

	router.HandleFunc("/", control.HandleApiRoot(app)).Name("api:root")
	router.HandleFunc("/clients", control.HandleClientsList(app)).Name("clients:index")
}
