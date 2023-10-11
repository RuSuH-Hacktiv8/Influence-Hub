package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var RapidAPIKey = os.Getenv("RAPID_API_KEY")
var RapidAPIHost = "instagram-data1.p.rapidapi.com"

func GetInstagramFollowers(username string) (int, error) {
	url := "https://instagram-data1.p.rapidapi.com/followers?username=" + username

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("X-RapidAPI-Key", RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", RapidAPIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("instagram API request failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var followersResponse struct {
		Count int `json:"count"`
	}

	if err := json.Unmarshal(body, &followersResponse); err != nil {
		return 0, err
	}

	return followersResponse.Count, nil
}
