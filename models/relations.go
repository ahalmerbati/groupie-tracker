package models

// The API LINK: https://groupietrackers.herokuapp.com/api/relations

type Relation struct {
	// Converting json --> go variables
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
