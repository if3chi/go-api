package kernel

import (
	"net/http"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/if3chi/go-api/pkg/config"
	"go.uber.org/zap"
)

type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

func Boot() *Application {
	config := config.Load()
	router := mux.NewRouter()
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	logger, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	return &Application{
		Server: &http.Server{
			Addr:         ":" + config.App.Port,
			Handler:      corsHandler(router),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}
}
