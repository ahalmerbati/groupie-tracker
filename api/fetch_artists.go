package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

const artistsAPI = "https://groupietrackers.herokuapp.com/api/artists"

func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get(artistsAPI)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var artists []models.Artist

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return artists, nil
}
