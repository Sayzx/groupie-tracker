package handler

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"main/internal/api"
	"net/http"
)

type GalleryData struct {
	Artists []api.Artist
	Years   []int
}

func generateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func UnifiedGalleryHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "" {
		token, err := discordOauthConfig.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, "Cannot exchange authorization code for access token", http.StatusBadRequest)
			return
		}

		user, err := GetUserDetails(token.AccessToken)
		if err != nil {
			http.Error(w, "Cannot retrieve user details", http.StatusInternalServerError)
			return
		}

		sessionID := generateSessionID()
		mu.Lock()
		connectedUsers[sessionID] = *user
		mu.Unlock()
		http.SetCookie(w, &http.Cookie{Name: "userSessionID", Value: sessionID, HttpOnly: true})

		// Récupérer l'URL de référence depuis le cookie
		referrerCookie, err := r.Cookie("referrerURL")
		if err != nil {
			// Cookie non trouvé, utiliser un fallback
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
			return
		}
		// Rediriger vers l'URL de référence
		http.Redirect(w, r, referrerCookie.Value, http.StatusSeeOther)
		return
	}

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
