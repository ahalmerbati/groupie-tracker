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

	// Extra variables for home page design
	MemberCount int
	LastConcert string
}
