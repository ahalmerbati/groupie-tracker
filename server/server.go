package server

import (
	"groupie-tracker/handlers"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe(":8080", nil)
}
