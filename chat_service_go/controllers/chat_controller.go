package controllers

import (
	"chat_service_go/models"
	"chat_service_go/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleChat(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		// Set proper headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Parse incoming request
		var cr models.ChatRequest
		err := json.NewDecoder(r.Body).Decode(&cr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 1) Ensure user exists in Python
		if err := services.PythonCreateUserIfNotExists(cr.UserID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2) If chat_id is empty or we want to allow flexible creation:
		chatID := cr.ChatID
		if chatID == "" {
			chatID, err = services.PythonCreateChat(cr.UserID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// 3) Store user message
		if err := services.PythonAddMessage(chatID, "user", cr.Message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 4) Get entire chat context
		msgs, err := services.PythonGetChatHistory(chatID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 5) Call Ollama to get a streaming response
		finalAssistantText, err := services.CallOllamaChat("phi4:latest", msgs, w)
		if err != nil {
			// Send error as SSE event
			fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
			logger.Println("ollama error:", err)
			return
		}

		// 6) Store the assistant's final text in the DB
		go func() {
			if err := services.PythonAddMessage(chatID, "assistant", finalAssistantText); err != nil {
				logger.Println("failed to store assistant message:", err)
			}
		}()
	}
} 