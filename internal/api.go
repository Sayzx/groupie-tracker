package internal

import (
	"encoding/json"
	"fmt"
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
	Dates        string   `json:"dates"`
}

// GetArtists function to fetch artists from the API
func GetArtists() ([]Artist, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	fmt.Println(response.Body)

	// Decode the JSON response
	var artists []Artist
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func GetArtistByID(id string) (Artist, error) {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists/" + id
	response, err := http.Get(apiURL)
	if err != nil {
		return Artist{}, err
	}
	defer response.Body.Close()

	// Decode the JSON response
	var artist Artist
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&artist); err != nil {
		return Artist{}, err
	}
	return artist, nil
}
