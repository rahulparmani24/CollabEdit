package main

import (
	"log"
	"net/http"
	"os"
	"CollabEdit/ws"
)

func main() {
	// Get environment variables for Redis host and server port
	redisHost := os.Getenv("REDIS_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default port
	}

	// Initialize the WebSocket hub
	hub := ws.NewHub(redisHost)

	// Set up the WebSocket handler
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleConnections(hub, w, r)
	})

	// Run the hub to manage clients
	go hub.Run()

	// Start the HTTP server
	log.Printf("Server started on :%s", serverPort)
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
