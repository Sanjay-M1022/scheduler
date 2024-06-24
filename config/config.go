package config

import (
	"fmt"
	"log"
	"os"
)

// ApiServerConfigurations holds the configuration for the API server
type ApiServerConfigurations struct {
	Port string
	Host string
}

// WsServerConfiguration holds the configuration for the WebSocket server
type WsServerConfiguration struct {
	Port string
	Host string
}

// Initialize default configurations for both API and WebSocket servers
var ApiServerConfig = ApiServerConfigurations{
	Host: ":",
	Port: "8000",
}

var WsServerConfig = WsServerConfiguration{
	Host: ":",
	Port: "8080",
}

// getEnv fetches the environment variable if set, otherwise it returns the default value.
// It logs a fatal error if the environment variable is not found.
func getEnv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Printf("env varible not set for, %s using default value %s ", key, defaultValue)
		return defaultValue
	}
	return value
}

// LoadConfig loads the configuration from environment variables.
// It overrides the default configurations if the environment variables are set.
func LoadConfig() {
	// API server config
	ApiServerConfig.Host = getEnv("API_SERVER_HOST", ApiServerConfig.Host)
	ApiServerConfig.Port = getEnv("API_SERVER_PORT", ApiServerConfig.Port)
	// Websocket config
	WsServerConfig.Host = getEnv("WS_SERVER_HOST", WsServerConfig.Host)
	WsServerConfig.Port = getEnv("WS_SERVER_PORT", WsServerConfig.Port)
}

func JoinHostAndPort(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
