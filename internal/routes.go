package internal

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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
	// Serveur de fichiers statiques pour les assets
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Route pour la page d'accueil")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/register.html")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// Récupération des paramètres de recherche à partir de l'URL
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		year := queryParams.Get("year")
		members := queryParams.Get("members")
		city := queryParams.Get("city")

		// Fetch data from the API using the function from api.go
		artists, err := GetArtists()
		if err != nil {
			http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
			return
		}

		// Filtrer les artistes en fonction des paramètres de recherche
		filteredArtists := filterArtists(artists, name, year, members, city)
		var years []int
		for year := 1960; year <= 2024; year++ {
			years = append(years, year)
		}

		data := GalleryData{
			Artists: filteredArtists,
			Years:   years,
		}

		tmpl, err := template.ParseFiles("web/templates/search.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/api/search/artists", func(writer http.ResponseWriter, request *http.Request) {
		request.URL.Query().Get("query")

		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(Body io.ReadCloser) {
			err1 := Body.Close()
			if err1 != nil {

			}
		}(resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if request.Method == http.MethodOptions {
			writer.Header().Set("Access-Control-Allow-Origin", "")
			writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			writer.WriteHeader(http.StatusNoContent)
			return
		}

		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Content-Type", "application/json")

		_, err = writer.Write(body)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from the API using the function from api.go
		artists, err := GetArtists()
		if err != nil {
			http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
			return
		}

		data := GalleryData{
			Artists: artists,
		}

		tmpl, err := template.ParseFiles("web/templates/gallery.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

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

func filterArtists(artists []Artist, name, year, members, city string) []Artist {
	var filtered []Artist
	for _, artist := range artists {
		if name != "" && !strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name)) {
			continue
		}
		if year != "" && artist.FirstAlbum[:4] != year {
			continue
		}
		if members != "" {
			membersInt, _ := strconv.Atoi(members) // Convertit le nombre de membres en entier
			if len(artist.Members) != membersInt {
				continue
			}
		}
		if city != "" {
			found := false
			city = strings.ToLower(strings.TrimSpace(city)) // Nettoyez l'entrée de recherche pour la ville
			for _, concertCity := range artist.ConcertCities {
				if city == strings.ToLower(strings.TrimSpace(concertCity)) {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}
		filtered = append(filtered, artist)
	}
	return filtered
}
