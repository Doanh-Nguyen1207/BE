package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// app *application is any time we have a variable of type application, we have access to this function
func (app *application) routes() http.Handler { //returns a http.Handler
	//create a router mux
	mux := chi.NewRouter()

	//define any middlewares here

	mux.Use(middleware.Recoverer) //'Use' will apply to every request that comes into our application
	// middleware.Recoverer does is when our application panics, it will log it along with a backtrace and show us where the error
	// took place. It'll send back the nessesary header, which is HTTP 500
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/movies", app.AllMovies)

	return mux
}
