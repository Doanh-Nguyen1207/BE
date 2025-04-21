package main

import (
	"fmt"
	"net/http"
)

// w stand for 'what do you write to?'
func (app *application) Home(w http.ResponseWriter, r *http.Request) { // ResponseWriter is where you write the final content you want to send to the client
	fmt.Fprintf(w, "Hello world from %s", app.Domain)
}
