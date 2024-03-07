package internal

import (
	"html/template"
	"net/http"
)

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
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
}
