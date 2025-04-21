package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	//setup config for the application
	var app application

	//read from command line

	//connect to the database

	app.Domain = "example.com"

	log.Println("Starting server on port", port)

	//connect to the web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
