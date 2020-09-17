package helper

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func WriteJSON(writer http.ResponseWriter, status int, v interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	j, _ := json.Marshal(v)

	writer.Write(j)
}

func WriteJSONError(writer http.ResponseWriter, status int, message string) {
	WriteJSON(writer, status, StatusResponse{
		Status:  "error",
		Message: message,
	})
}
