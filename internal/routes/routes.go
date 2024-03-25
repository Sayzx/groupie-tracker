package routes

import (
	"fmt"
	"main/internal/db"
	"main/internal/handler"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {
	fmt.Println("Initialisation du serveur...")
	dbConn := db.InitDB()
	defer dbConn.Close()
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	http.HandleFunc("/search", handler.SearchHandler)
	http.HandleFunc("/api/search/artists", handler.SearchArtistsHandler)
	http.HandleFunc("/gallery", handler.UnifiedGalleryHandler)
	http.HandleFunc("/artist_info", handler.ArtisteInfo)
	http.HandleFunc("/comment", handler.SubmitCommentHandler(dbConn))
	http.HandleFunc("/discord", handler.DiscordLoginHandler)
	http.HandleFunc("/api/comments", handler.GetCommentsHandler(dbConn))

	handler.Proxy()

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}
}
