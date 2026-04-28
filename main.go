package main

import (
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Call Gemini
	resp, err := CallGemini(string(body))
	if err != nil {
		http.Error(w, "Error in sending request to Gemini: "+err.Error(), http.StatusBadGateway)
		return
	}

	// Try to read the response
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Return the response!
	w.Write([]byte(string(contents)))
}
