package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/time", timeHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/unzip", unzipHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Starting web server on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /           - Home page")
	fmt.Println("  GET  /api/time    - Current time")
	fmt.Println("  GET  /api/health  - Health check")
	fmt.Println("  GET  /api/unzip   - Unzip and list crc-linux-amd64.tar.xz contents")

	log.Fatal(http.ListenAndServe(port, nil))
}
