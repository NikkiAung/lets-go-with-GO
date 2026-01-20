package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/NikkiAung/go-fundmentals/internal/app"
	"github.com/NikkiAung/go-fundmentals/internal/routes"
)

func main () {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Our app is live.")

	var port int 

	flag.IntVar(&port, "port", 8080, "Use to change server port.")
	flag.Parse()

	r := routes.SetUpRoutes(app)

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	app.Logger.Printf("Server is running on port %d\n", port)

	err = s.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}