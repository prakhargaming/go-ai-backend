package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CallGemini(query string) (*http.Response, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env vars
	url := os.Getenv("GEMINI_URL")
	api_key := os.Getenv("GEMINI_API_KEY")
	if url == "" || api_key == "" {
		return nil, errors.New("GEMINI_URL or GEMINI_API_KEY not found.")
	}

	// Build da request
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
