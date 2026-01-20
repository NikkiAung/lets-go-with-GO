package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NikkiAung/go-fundmentals/internal/api"
)

type Application struct {
	Logger *log.Logger
	PostHandler *api.PostHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	postHandler := api.NewPostHandler()

	app := &Application{
		Logger: logger,
		PostHandler: postHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is ok\n");
}