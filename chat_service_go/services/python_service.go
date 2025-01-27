package services

import (
	"chat_service_go/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const pythonBaseURL = "http://localhost:8000"

func PythonCreateUserIfNotExists(userID string) error {
	url := fmt.Sprintf("%s/users/%s", pythonBaseURL, userID)
	req, _ := http.NewRequest("POST", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("bad status from python: %d", resp.StatusCode)
	}
	return nil
}

func PythonCreateChat(userID string) (string, error) {
	url := fmt.Sprintf("%s/chats/?user_id=%s", pythonBaseURL, userID)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("bad status from python: %d", resp.StatusCode)
	}

	var data struct {
		ChatID string `json:"chat_id"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.ChatID, nil
}

func PythonAddMessage(chatID, role, content string) error {
	msg := models.Message{Role: role, Content: content}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(msg); err != nil {
		return err
	}

	url := fmt.Sprintf("%s/chats/%s/messages", pythonBaseURL, chatID)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad status from python: %d, body=%s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}

func PythonGetChatHistory(chatID string) ([]models.Message, error) {
	url := fmt.Sprintf("%s/chats/%s/messages", pythonBaseURL, chatID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad status from python: %d", resp.StatusCode)
	}
	var msgs []models.Message
	if err = json.NewDecoder(resp.Body).Decode(&msgs); err != nil {
		return nil, err
	}
	return msgs, nil
} 