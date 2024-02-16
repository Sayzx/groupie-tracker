package internal

import (
	"fmt"
	"net/http"
)

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

	http.HandleFunc("/search2", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/search.html")
	})

	fmt.Println("Server started at http://localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}

}
