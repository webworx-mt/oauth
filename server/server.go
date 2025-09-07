package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webworx-mt/oauth/handlers"
)

// StartServer starts the HTTP server
func StartServer(port string) {
	// Define routes
	http.HandleFunc("/health", handlers.HealthHandler)

	// Start server
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /health - Health check")

	log.Fatal(http.ListenAndServe(port, nil))
}
