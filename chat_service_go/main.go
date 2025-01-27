package main

import (
	"chat_service_go/controllers"
	"chat_service_go/_utils"
	"log"
	"net/http"
)

func main() {
	// Initialize logger
	logger := _utils.NewLogger()

	// Set up routes
	http.HandleFunc("/api/chat", controllers.HandleChat(logger))

	logger.Println("Go Chat Orchestrator listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
} 