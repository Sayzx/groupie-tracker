package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
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

type Relation map[string][]string

// Structure pour stocker les dates et lieux temporaires
type DateLocation struct {
	Date     time.Time
	Location string
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

// func GetArtistRelations(id string) (Relation, error) {
// 	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		return Relation{}, err
// 	}
// 	defer response.Body.Close()

// 	// Decode the JSON response
// 	var relation Relation
// 	decoder := json.NewDecoder(response.Body)
// 	if err := decoder.Decode(&relation); err != nil {
// 		return Relation{}, err
// 	}

// 	fmt.Println(relation)
// 	return relation, nil
// }

func GetAndFormatArtistRelations(id string) ([]string, error) {
	apiURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var relation Relation
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&relation); err != nil {
		return nil, err
	}

	var datesLocations []DateLocation
	for location, dates := range relation {
		for _, dateStr := range dates {
			date, err := time.Parse("02-01-2006", dateStr)
			if err != nil {
				return nil, err
			}
			datesLocations = append(datesLocations, DateLocation{Date: date, Location: location})
		}
	}

	// Trier par date
	sort.Slice(datesLocations, func(i, j int) bool {
		return datesLocations[i].Date.Before(datesLocations[j].Date)
	})

	var formatted []string
	for _, dl := range datesLocations {
		formatted = append(formatted, fmt.Sprintf("Date: %s, Lieu: %s", dl.Date.Format("02-01-2006"), dl.Location))
	}
	return formatted, nil
}
