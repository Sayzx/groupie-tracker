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
	// Extraction de l'ID de l'artiste depuis les paramètres de requête
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}

	// Récupération des informations de l'artiste
	artist, err := GetArtistByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	// Nouveau: Récupération des relations de l'artiste
	relations, err := GetRelationByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist relations", http.StatusInternalServerError)
		return
	}

	// Préparation des données pour le template, incluant les relations
	data := ArtistInfoData{
		Artist:    artist,
		Relations: relations,
	}

	// Chargement et exécution du template avec les données
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
