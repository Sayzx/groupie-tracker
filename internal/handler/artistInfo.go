package handler

import (
	"fmt"
	"html/template"
	"main/internal/api"
	"net/http"
	"strconv"
	"strings"
)

type ArtistInfoData struct {
	Artist    api.Artist
	Relations api.Relation
	SpotifyID string
}

func ArtisteInfo(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}

	artist, err := api.GetArtistByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	relations, err := api.GetRelationByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist relations", http.StatusInternalServerError)
		return
	}

	token, err := api.GetSpotifyToken("ea4f316cdc894f59aed435cc6f7f0e6e", "0844fe38663c4010a9fa8f193d4aa95b") // Remplacez par vos identifiants
	if err != nil {
		http.Error(w, "Failed to get Spotify token", http.StatusInternalServerError)
		return
	}

	// Utilisez le nom de l'artiste pour obtenir l'ID Spotify
	spotifyArtist, err := api.SearchArtist(artist.Name, token)
	if err != nil {
		http.Error(w, "Failed to fetch Spotify artist", http.StatusInternalServerError)
		return
	}

	data := ArtistInfoData{
		Artist:    artist,
		Relations: relations,
		SpotifyID: spotifyArtist.ID,
	}

	tmpl, err := template.ParseFiles("web/templates/artist_info.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}

func filterArtists(artists []api.Artist, name, year, members, creationYear string) []api.Artist {
	var filtered []api.Artist
	for _, artist := range artists {
		if name != "" && !strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name)) {
			continue
		}

		firstAlbumYear := artist.FirstAlbum[len(artist.FirstAlbum)-4:]
		if year != "" && firstAlbumYear != year {
			continue
		}

		if creationYear != "" {
			creationYearInt, err := strconv.Atoi(creationYear)
			if err != nil || artist.CreationDate != creationYearInt {
				continue
			}
		}

		if members != "" {
			membersInt, err := strconv.Atoi(members) // Convertit le nombre de membres en entier
			if err != nil || len(artist.Members) != membersInt {
				continue
			}
		}
		filtered = append(filtered, artist)
	}
	return filtered
}
