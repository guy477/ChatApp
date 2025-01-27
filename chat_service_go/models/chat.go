package models

type ChatRequest struct {
    UserID  string `json:"user_id"`
    ChatID  string `json:"chat_id"`
    Message string `json:"message"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OllamaChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OllamaChatRequest struct {
    Model    string               `json:"model"`
    Messages []OllamaChatMessage  `json:"messages"`
    Stream   bool                 `json:"stream"`
} 