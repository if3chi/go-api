package kernel

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func (app *Application) Respond(response http.ResponseWriter, request *http.Request, data interface{}, status int) {
	response.WriteHeader(status)
	response.Header().Add("Content-Type", app.Config.HTTP.Content)

	if data != nil {
		if err := json.NewEncoder(response).Encode(data); err != nil {
			app.Logger.Fatal(err.Error())
			panic(err)
		}
	}

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

func (app *Application) Run() {
	if err := app.Server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err.Error())
		panic(err)
	}
}

func (app *Application) ListenForShutdown() {
	interuptChannel := make(chan os.Signal, 1)
	signal.Notify(interuptChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interuptChannel

	app.Logger.Info("Recieved exit signal, shutting down.....")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.Server.Shutdown(ctx)
	os.Exit(1)
}
