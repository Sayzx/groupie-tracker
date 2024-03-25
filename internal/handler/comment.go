package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Comment struct {
	DiscordName   string `json:"discordName"`
	DiscordAvatar string `json:"discordAvatar"`
	CommentText   string `json:"commentText"`
}

func SubmitCommentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %v", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		discordName := r.FormValue("discordName")
		discordAvatar := r.FormValue("discordAvatar")
		comment := r.FormValue("comment")
		idStr := r.FormValue("id")

		// Afficher les valeurs reçues pour débogage
		log.Println("Received form data:", discordName, discordAvatar, comment, idStr)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Invalid artist ID: %s, error: %v", idStr, err)
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO comments (discord_name, discord_avatar, comment, artist_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Printf("Error preparing SQL statement: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(discordName, discordAvatar, comment, id)
		if err != nil {
			log.Printf("Error executing SQL statement: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/artist_info?id="+idStr, http.StatusSeeOther)
	}
}

func GetCommentsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		artistID := r.URL.Query().Get("id")

		query := "SELECT discord_name, discord_avatar, comment FROM comments WHERE artist_id = ?"
		rows, err := db.Query(query, artistID)
		if err != nil {
			log.Printf("Erreur lors de la requête SQL: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var comments []Comment
		for rows.Next() {
			var c Comment
			if err := rows.Scan(&c.DiscordName, &c.DiscordAvatar, &c.CommentText); err != nil {
				log.Printf("Erreur lors de la récupération d'un commentaire: %v", err)
				continue
			}
			comments = append(comments, c)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Erreur rencontrée lors du parcours des lignes: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse
		if err := json.NewEncoder(w).Encode(comments); err != nil {
			log.Printf("Erreur lors de l'encodage des commentaires en JSON: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		}
	}
}
