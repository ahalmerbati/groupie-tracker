package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/processor"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func BandHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/band/") {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[2]
	if len(parts) != 3 || idStr == "" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	artists, err := processor.ProcessData()
	if err != nil {
		log.Printf("Data processing error in band handler: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	var selectedArtist models.Artist
	found := false

	for _, artist := range artists {
		if artist.ID == id {
			selectedArtist = artist
			found = true
			break
		}
	}

	if !found {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	renderTemplate(w, r, "band.html", selectedArtist)
}
