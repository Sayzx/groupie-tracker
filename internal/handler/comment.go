package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		// Vérifier que la méthode de la requête est POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parser le formulaire pour accéder aux données soumises
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %v", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Extraire les données du formulaire
		discordName := r.FormValue("discordName")
		discordAvatar := r.FormValue("discordAvatar")
		comment := r.FormValue("comment")
		idStr := r.FormValue("id")

		// Afficher les valeurs reçues pour débogage
		log.Println("Received form data:", discordName, discordAvatar, comment, idStr)

		// Convertir l'ID de l'artiste en entier
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Printf("Invalid artist ID: %s, error: %v", idStr, err)
			http.Error(w, "Invalid artist ID", http.StatusBadRequest)
			return
		}

		// Préparer l'instruction SQL pour insérer le commentaire
		stmt, err := db.Prepare("INSERT INTO comments (discord_name, discord_avatar, comment, artist_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Printf("Error preparing SQL statement: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		// Exécuter l'instruction SQL
		_, err = stmt.Exec(discordName, discordAvatar, comment, id)
		if err != nil {
			log.Printf("Error executing SQL statement: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Redirection de l'utilisateur après la soumission réussie
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func GetCommentsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupération de l'artistID depuis l'URL
		artistID := r.URL.Query().Get("id")
		fmt.Printf("Artist ID: %s\n", artistID) // Affiche l'ID pour déboguer

		// Préparation de la requête SQL
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
				continue // Tentez de lire le commentaire suivant en cas d'erreur
			}
			comments = append(comments, c)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Erreur rencontrée lors du parcours des lignes: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(comments); err != nil {
			log.Printf("Erreur lors de l'encodage des commentaires en JSON: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		}
	}
}
