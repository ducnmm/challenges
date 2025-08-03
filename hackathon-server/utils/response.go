package utils

import (
	"encoding/json"
	"hackathon-server/models"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}