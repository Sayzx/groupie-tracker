package internal

import (
	"encoding/json"
	"net/http"
)

// Artist structure
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	FirstRelease string   `json:"firstRelease"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// GetArtists function to fetch artists from the API
func GetArtists() ([]Artist, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the JSON response
	var artists []Artist
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}
