package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NikkiAung/go-fundmentals/internal/api"
	"github.com/NikkiAung/go-fundmentals/internal/store"
	"github.com/NikkiAung/go-fundmentals/migrations"
)

type Application struct {
	Logger *log.Logger
	PostHandler *api.PostHandler
	UserHandler *api.UserHandler
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	postgresDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(postgresDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	postStore := store.NewPostgresPostStore(postgresDB)
	userStore := store.NewPostgresUserStore(postgresDB)

	postHandler := api.NewPostHandler(postStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)

	app := &Application{
		Logger: logger,
		PostHandler: postHandler,
		UserHandler: userHandler,
		DB : postgresDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is ok\n");
}