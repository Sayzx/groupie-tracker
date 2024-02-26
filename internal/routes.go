package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

type GalleryData struct {
	Artists []Artist
}

// ArtistInfoData structure to pass data to the artist_info template
type ArtistInfoData struct {
	Artist    Artist
	Relations []string
}

func Run() {
	fmt.Println("Initialisation du serveur...")
	// Serveur de fichiers statiques pour les assets
	fs := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Route pour la page d'accueil")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/register.html")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from the API using the function from api.go
		artists, err := GetArtists()
		if err != nil {
			http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
			return
		}

		data := GalleryData{
			Artists: artists,
		}

		tmpl, err := template.ParseFiles("./web/templates/search.html")
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

		tmpl, err := template.ParseFiles("./web/templates/gallery.html")
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
	// Parse artist ID from query parameters
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}

	// Fetch data for the specific artist
	artist, err := GetArtistByID(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist data", http.StatusInternalServerError)
		return
	}

	// Fetch and format artist relations
	relations, err := GetAndFormatArtistRelations(artistID)
	if err != nil {
		http.Error(w, "Error fetching artist relations", http.StatusInternalServerError)
		return
	}

	data := ArtistInfoData{
		Artist:    artist,
		Relations: relations,
	}

	// Parse the template
	tmpl, err := template.ParseFiles("./web/templates/artist_info.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Execute the template with data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
