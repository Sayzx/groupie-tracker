package handler

import (
	"database/sql"
	"log"
	"net/http"
)

// Assuming db is your *sql.DB connection established using initDB()
// Make sure to import your database initialization package and use it here to get the DB connection

func SubmitCommentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			log.Println("Database connection is nil, redirecting to /")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Extrait les valeurs du formulaire
		discordName := r.FormValue("discordName")
		discordAvatar := r.FormValue("discordAvatar")
		comment := r.FormValue("comment")
		id := r.FormValue("id")

		// Vérifie si les valeurs sont vides
		if discordName == "" || discordAvatar == "" {
			// Gérer le cas où les valeurs sont manquantes
			log.Println("Les valeurs de l'avatar ou du nom de l'artiste sont manquantes")
		} else {
			// Utiliser les valeurs récupérées
			log.Printf("Nom de l'artiste : %s, Avatar : %s\n", discordName, discordAvatar)
		}

		// Debug print
		log.Printf("Comment: %s\n", comment)
		log.Printf("Discord Name: %s\n", discordName)
		log.Printf("Discord Avatar: %s\n", discordAvatar)
		log.Printf("Artist ID: %s\n", id)

		// Prepare SQL statement
		log.Println("Preparing SQL statement...")
		stmt, err := db.Prepare("INSERT INTO comments (discord_name, discord_avatar, comment, artist_id) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatalf("Error preparing SQL statement: %v\n", err)
		}
		defer stmt.Close()

		// Execute SQL statement
		log.Println("Executing SQL statement...")
		_, err = stmt.Exec(discordName, discordAvatar, comment, id)
		if err != nil {
			log.Fatalf("Error executing SQL statement: %v\n", err)
		}

		// Redirect after successful comment submission
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
