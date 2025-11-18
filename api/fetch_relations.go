package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

const relationsAPI = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelations() ([]models.Relation, error) {
	resp, err := http.Get(relationsAPI)
	if err != nil {
		return nil, fmt.Errorf("error fetching relations API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var relationsIndex struct {
		Index []models.Relation `json:"index"`
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&relationsIndex)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return relationsIndex.Index, nil
}
