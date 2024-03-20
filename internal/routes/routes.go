package routes

import (
	"fmt"
	"main/internal/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	fmt.Println("Initialisation du serveur...")

	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	http.HandleFunc("/search", handler.SearchHandler)
	http.HandleFunc("/api/search/artists", handler.SearchArtistsHandler)
	http.HandleFunc("/gallery", handler.GalleryHandler)
	http.HandleFunc("/artist_info", handler.ArtisteInfo)

	fmt.Println("Server started at http://localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}
