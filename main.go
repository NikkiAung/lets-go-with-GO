package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NikkiAung/go-fundmentals/internal/app"
)

func main () {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Our app is live.")

	s := &http.Server{
		Addr: ":8080",
		IdleTimeout: time.Minute,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// curl localhost:8080/health
	http.HandleFunc("/health", HealthCheck)

	err = s.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is ok\n");
}