package processor

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"groupie-tracker/utils"
)

func ProcessData() ([]models.Artist, error) {
	artists, err := api.FetchArtists()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artists: %w", err)
	}

	relations, err := api.FetchRelations()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch relations: %w", err)
	}

	// Key = ArtistID | Value = Relation Object/Struct
	relationsMap := make(map[int]models.Relation)
	for _, relation := range relations {
		relationsMap[relation.ID] = relation
	}

	for i := range artists {
		artist := &artists[i]

		relation, found := relationsMap[artist.ID]
		if found {
			artist.Relations = relation
		}

		artist.MemberCount = utils.CountMembers(artist.Members)
		artist.LastConcert = utils.FindLastConcertDate(artist.Relations.DatesLocations)
	}
	return artists, nil
}
