package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webworx-mt/oauth/handlers"
	"github.com/webworx-mt/oauth/router"
)

// StartServer starts the HTTP server
func StartServer(port string) {
	// Create router
	r := router.New()

	// Define routes with REST conventions
	r.GET("/health", handlers.HealthHandler)
	r.GET("/users", handlers.GetUsersHandler)
	r.GET("/user/:id", handlers.GetUserHandler)

	// Start server
	fmt.Printf("Server starting on port %s\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}
