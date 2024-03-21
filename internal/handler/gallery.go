package handler

import (
	"html/template"
	"main/internal/api"
	"net/http"
)

type GalleryData struct {
	Artists []api.Artist
	Years   []int
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.GetArtists()
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
