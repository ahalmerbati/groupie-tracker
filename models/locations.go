package models

// The API LINK: https://groupietrackers.herokuapp.com/api/locations

type Location struct {
	// Converting json --> go variables
	ID        int      `json:"id"`
	Locations []string `json:"locations"`

	// Converting json --> go url links (API)
	DatesURL string `json:"dates"`

	// Struct to store the fetched data from the API URLs
	Dates Date
}
