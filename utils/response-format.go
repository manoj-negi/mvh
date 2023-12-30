package util

import (
	"encoding/json"
	"net/http"
)

type ErrorFormat struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	errResponse := ErrorFormat{
		Status:  statusCode,
		Message: message,
	}

	json.NewEncoder(w).Encode(errResponse)

}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
