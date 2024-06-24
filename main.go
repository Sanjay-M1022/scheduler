package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/scheduler/config"
	"github.com/scheduler/services"
)

func main() {
	config.LoadConfig()
	InitializeWebSocketRouter()
	router := InitializeApiRouter()

	log.Println("Starting Job runner")
	go services.JobRunner()

	log.Printf("API server is running on %s", config.WsServerConfig.Port)
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.WsServerConfig.Port), nil))
	}()

	log.Printf("Listening for websocket connections on %s", config.ApiServerConfig.Port)
	go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ApiServerConfig.Port), router))
}
