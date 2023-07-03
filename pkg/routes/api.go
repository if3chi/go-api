package routes

import (
	"net/http"

	control "github.com/if3chi/go-api/pkg/controllers"
	externalControl "github.com/if3chi/go-api/pkg/controllers/externalapi"
	"github.com/if3chi/go-api/pkg/kernel"
)

func Load(app *kernel.Application) {
	router := app.Router.Methods(http.MethodGet).Subrouter()

	router.HandleFunc("/", control.HandleApiRoot(app)).Name("api:root")
	router.HandleFunc("/clients", control.HandleClientsList(app)).Name("clients:index")
	router.HandleFunc("/photos", externalControl.HandleGetPhotos(app)).Name("photos:list")
}
