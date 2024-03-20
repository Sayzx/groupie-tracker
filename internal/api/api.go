package api

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
	Relations    string   `json:"relations"`
}

type LocationDates struct {
	Location string
	Dates    []string
}

type Relation struct {
	ID             int
	DatesLocations []LocationDates
	Dates          []string
	Cities         []string // Add Cities field
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

func GetRelationByID(id string) (Relation, error) {
	var apiResponse struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	var relation Relation

	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
	response, err := http.Get(apiURL)
	if err != nil {
		return relation, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&apiResponse); err != nil {
		return relation, err
	}

	relation.ID = apiResponse.ID
	for city, dates := range apiResponse.DatesLocations {
		if len(dates) > 0 {
			relation.Dates = append(relation.Dates, dates[0])
			relation.Cities = append(relation.Cities, city)
		}
	}

	return relation, nil
}
