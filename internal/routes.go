package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type GalleryData struct {
	Artists []Artist
	Years   []int
}

// ArtistInfoData structure to pass data to the artist_info template
type ArtistInfoData struct {
	Artist    Artist
	Relations Relation
	SpotifyID string
}

func Run() {
	fmt.Println("Initialisation du serveur...")

	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/api/search/artists", SearchArtistsHandler)
	http.HandleFunc("/gallery", GalleryHandler)
	http.HandleFunc("/artist_info", ArtisteInfo)

	fmt.Println("Server started at http://localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}

func ArtisteInfo(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}

	artist, err := GetArtistByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	relations, err := GetRelationByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist relations", http.StatusInternalServerError)
		return
	}

	token, err := getSpotifyToken("ea4f316cdc894f59aed435cc6f7f0e6e", "0844fe38663c4010a9fa8f193d4aa95b") // Remplacez par vos identifiants
	if err != nil {
		http.Error(w, "Failed to get Spotify token", http.StatusInternalServerError)
		return
	}

	// Utilisez le nom de l'artiste pour obtenir l'ID Spotify
	spotifyArtist, err := searchArtist(artist.Name, token)
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

func filterArtists(artists []Artist, name, year, members, creationYear string) []Artist {
	var filtered []Artist
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
