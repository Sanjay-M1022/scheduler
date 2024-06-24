package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scheduler/handlers"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// InitializeRouter sets up the router and routes
func InitializeApiRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	// Rest api
	router.HandleFunc("/jobs", handlers.HomeRoute).Methods("GET")
	router.HandleFunc("/jobs", handlers.CreateJob).Methods("POST")
	return router
}

func InitializeWebSocketRouter() {
	http.HandleFunc("/ws", handlers.JobUpdateBroadcaster)
	fmt.Println("here")
	// TODO: Move server start to main.
}
