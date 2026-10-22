package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Kubia server starting...")

	handler := func(w http.ResponseWriter, r *http.Request) {
		// Log the remote IP/address of the incoming request
		log.Printf("Received request from %s\n", r.RemoteAddr)

		// Get the pod/container hostname
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		// Write HTTP 200 status header and response string
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You've hit %s\n", hostname)
	}

	http.HandleFunc("/", handler)

	// Start listening on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
