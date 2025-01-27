package _utils

// Config holds all application configuration
type Config struct {
	OllamaURL string
}

var AppConfig Config

// InitConfig initializes application configuration
func InitConfig() {
	AppConfig = Config{
		OllamaURL: "http://localhost:11434/api/chat", // default value
	}
} 