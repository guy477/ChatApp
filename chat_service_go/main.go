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

	// Initialize configuration
	// If Ollama is hosted on a different machine, change its endpoint in config.go
	_utils.InitConfig()

	// Set up routes
	http.HandleFunc("/api/chat", controllers.HandleChat(logger))

	// Listen on all interfaces (0.0.0.0) instead of just localhost
	logger.Println("Go Chat Orchestrator listening on 0.0.0.0:3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}