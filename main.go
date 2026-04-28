package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	var port string = ":8080"
	if len(args) > 0 {
		port = ":" + args[0]
	}
	fmt.Println("Started go backend at", port)

	http.HandleFunc("/chat", handleRequest)
	http.ListenAndServe(port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Validate the request method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Try to read the request body
	body, err_resp := io.ReadAll(r.Body)
	if err_resp != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Call Gemini
	resp, err_gem := CallGemini(string(body))
	if err_gem != nil {
		http.Error(w, "Error in sending request to Gemini: "+err_gem.Error(), http.StatusBadGateway)
		return
	}

	// Return the response!
	var result GeminiResponse
	err_dec := json.NewDecoder(resp.Body).Decode(&result)
	if err_dec != nil {
		http.Error(w, "Could not decode response"+err_gem.Error(), http.StatusBadRequest)
		return
	}

	contents := result.Candidates[0].Content.Parts[0].Text
	w.Write([]byte(contents))
}
