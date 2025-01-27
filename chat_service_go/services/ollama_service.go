package services

import (
	"chat_service_go/models"
	"chat_service_go/_utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CallOllamaChat(model string, msgs []models.Message, w http.ResponseWriter) (string, error) {
	// Convert messages to Ollama format
	ollamaMessages := make([]models.OllamaChatMessage, 0, len(msgs))
	for _, m := range msgs {
		ollamaMessages = append(ollamaMessages, models.OllamaChatMessage{
			Role:    m.Role,
			Content: m.Content,
		})
	}

	reqData := models.OllamaChatRequest{
		Model:    model,
		Messages: ollamaMessages,
		Stream:   true,
	}
	bodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBuf).Encode(reqData); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", _utils.AppConfig.OllamaURL, bodyBuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Hour, // Long enough for big completions
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama returned status %d: %s", resp.StatusCode, string(b))
	}

	// Stream response to client and accumulate assistant's response
	var assistantResponse string
	dec := json.NewDecoder(resp.Body)

	for {
		var chunk struct {
			Model   string `json:"model"`
			Created string `json:"created_at"`
			Done    bool   `json:"done"`
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		}
		err = dec.Decode(&chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("decoding chunk: %v", err)
		}

		// Accumulate partial content
		assistantResponse += chunk.Message.Content

		// Send as SSE event
		fmt.Fprintf(w, "data: %s\n\n", chunk.Message.Content)
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}

		if chunk.Done {
			break
		}
	}

	return assistantResponse, nil
} 