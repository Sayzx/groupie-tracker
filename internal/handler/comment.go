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
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %v\n", err)
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Extract form values
		comment := r.FormValue("comment")
		discordName := r.FormValue("discordName")
		discordAvatar := r.FormValue("discordAvatar")

		// Prepare SQL statement
		stmt, err := db.Prepare("INSERT INTO comments (discord_name, discord_avatar, comment) VALUES (?, ?, ?)")
		if err != nil {
			log.Fatalf("Error preparing SQL statement: %v\n", err)
		}
		defer stmt.Close()

		// Execute SQL statement
		_, err = stmt.Exec(discordName, discordAvatar, comment)
		if err != nil {
			log.Fatalf("Error executing SQL statement: %v\n", err)
		}

		// Redirect or inform the user after successful comment submission
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
