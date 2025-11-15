package models

// The API LINK: https://groupietrackers.herokuapp.com/api/dates

type Date struct {
	// Converting json --> go varibles
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
