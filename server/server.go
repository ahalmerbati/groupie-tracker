package server

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func StartServer() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/band/", handlers.BandHandler)

	log.Printf("Server starting on http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
