package handler

import (
	"io"
	"net/http"
)

func Proxy() {
	http.HandleFunc("/api/locations", ProxyHandler)
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'URL de l'API externe
	apiURL := "https://groupietrackers.herokuapp.com" + r.URL.Path

	// Faire une requête vers l'API externe
	response, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error proxying request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Copier les en-têtes de la réponse de l'API externe vers la réponse du serveur
	for name, values := range response.Header {
		w.Header().Set(name, values[0])
	}

	// Envoyer le statut HTTP de la réponse de l'API externe
	w.WriteHeader(response.StatusCode)

	// Copier le corps de la réponse de l'API externe vers la réponse du serveur
	io.Copy(w, response.Body)
}
