package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/oauth2"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"` // Add avatar field
}

var (
	discordOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/gallery",
		ClientID:     "1220356595108544552",
		ClientSecret: "J5WndGCx2J4O0Clj_v8__f4h0TcytI82",
		Scopes:       []string{"identify", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}

	connectedUsers = make(map[string]User)
	mu             sync.Mutex
)

func DiscordLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Début de DiscordLoginHandler")
	url := discordOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	fmt.Printf("URL d'autorisation OAuth2 générée : %s\n", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	fmt.Println("Redirection effectuée")
}

// GetUserDetails récupère les détails de l'utilisateur Discord à partir du token d'accès
func GetUserDetails(accessToken string) (*User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
