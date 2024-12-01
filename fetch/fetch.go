package fetch

import (
	"encoding/json"
	"log"
	"net/http"
	"random-learning/models"
)

// Question represents a StackExchange question
type Question struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

const apiURL = "https://api.stackexchange.com/2.3/sites"

// FetchSites fetches available StackExchange sites from the API
func FetchSites() []models.Site {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Failed to fetch data from StackExchange API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var response models.StackExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	if len(response.Items) == 0 {
		log.Fatal("No sites found in the response")
	}

	return response.Items
}
