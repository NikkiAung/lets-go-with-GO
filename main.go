package main

import "github.com/NikkiAung/go-fundmentals/internal/app"

func main () {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("hello our logger.")
}