package main

import "net/http"

// Function handler that returns an OK string to indicate that the server is running.
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}