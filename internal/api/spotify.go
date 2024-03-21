package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpotifyArtist struct {
	ID   string
	Name string
}

func GetSpotifyToken(clientID, clientSecret string) (string, error) {
	url := "https://accounts.spotify.com/api/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	err = json.NewDecoder(res.Body).Decode(&tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}

func SearchArtist(artistName, accessToken string) (SpotifyArtist, error) {
	var artist SpotifyArtist
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist&limit=1", url.QueryEscape(artistName))

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()

	var searchResult struct {
		Artists struct {
			Items []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"items"`
		} `json:"artists"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return artist, err
	}

	if len(searchResult.Artists.Items) > 0 {
		artist.ID = searchResult.Artists.Items[0].ID
		artist.Name = searchResult.Artists.Items[0].Name
		return artist, nil
	}

	return artist, fmt.Errorf("artist not found")
}
