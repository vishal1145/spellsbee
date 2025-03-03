package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DatamuseWord struct {
	Word string `json:"word"`
}

func ValidateWord(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	word := parts[len(parts)-1]  // Get last segment
	word = strings.ToLower(word) // Convert to lowercase

	if word == "" {
		http.Error(w, `{"error": "Word parameter is required"}`, http.StatusBadRequest)
		return
	}

	// Make request to Datamuse API
	apiURL := fmt.Sprintf("https://api.datamuse.com/words?sp=%s", word)
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch data from Datamuse API"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, `{"error": "Failed to read API response"}`, http.StatusInternalServerError)
		return
	}

	var words []DatamuseWord
	if err := json.Unmarshal(body, &words); err != nil {
		http.Error(w, `{"error": "Failed to parse API response"}`, http.StatusInternalServerError)
		return
	}

	isValid := false
	for _, item := range words {
		if strings.ToLower(item.Word) == strings.ToLower(word) {
			isValid = true
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"word":    word,
		"isValid": isValid,
	})
}
