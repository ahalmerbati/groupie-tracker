package handlers

import (
	"groupie-tracker/processor"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	artists, err := processor.ProcessData()
	if err != nil {
		log.Printf("Data processing error: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	renderTemplate(w, r, "home.html", artists)
}
