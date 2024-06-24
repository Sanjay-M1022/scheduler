package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scheduler/models"
	"github.com/scheduler/services"
)

type RequestBody struct {
	Name         string `json:"name"`
	RequiredTime int    `json:"requiredTime"`
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resJson, _ := json.Marshal(services.GetJobsList())
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
	// w.Write([]byte(resJson))
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var requestBody RequestBody

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newJob := models.Job{
		Name:          requestBody.Name,
		Status:        "pending",
		RemainingTime: requestBody.RequiredTime,
		RequiredTime:  requestBody.RequiredTime,
	}
	services.AddJob(newJob)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Request received successfully"})

}
