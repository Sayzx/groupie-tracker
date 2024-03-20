package handler

import (
	"html/template"
	"io"
	"io/ioutil"
	"main/internal/api"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération des paramètres de recherche à partir de l'URL
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	year := queryParams.Get("year")
	members := queryParams.Get("members")
	city := queryParams.Get("city")

	// Fetch data from the API using the function from api.go
	artists, err := api.GetArtists()
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
}

func SearchArtistsHandler(writer http.ResponseWriter, request *http.Request) {
	request.URL.Query().Get("query")

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			http.Error(writer, err1.Error(), http.StatusInternalServerError)
			return
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
}
