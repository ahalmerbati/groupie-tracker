package models

// The API LINK: https://groupietrackers.herokuapp.com/api/artists

type Artist struct {
	// Converting json --> go variables
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	// Converting json --> go url links (API)
	LocationsURL   string `json:"locations"`
	ConcertDateURL string `json:"concertDates"`
	RelationsURL   string `json:"relations"`

	// Struct to store the fetched data from the API URLs
	Locations Location
	Dates     Date
	Relations Relation

	// Extra variables for home page design
	MemberCount int
	LastConcert string
}
