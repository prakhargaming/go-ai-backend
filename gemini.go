package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

func CallGemini(query string) (*http.Response, error) {
	url := os.Getenv("GEMINI_URL")
	api_key := os.Getenv("GOOGLE_API_KEY")
	if url == "" || api_key == "" {
		return nil, errors.New("GEMINI_URL or GOOGLE_API_KEY not found.")
	}
	reqBody := BuildRequest(query)
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", api_key)

	// Send it
	client := &http.Client{}
	return client.Do(req)
}
