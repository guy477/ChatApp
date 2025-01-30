package _utils

// Config holds all application configuration
type Config struct {
	OllamaURL string
}

var AppConfig Config

// InitConfig initializes application configuration
func InitConfig() {
	AppConfig = Config{
		OllamaURL: "http://packapunch-b550m-ds3h-ac.tail83fc6.ts.net:11434/api/chat", // default value
	}
} 