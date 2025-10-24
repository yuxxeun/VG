package utils

import (
	"encoding/json"
	"net/http"

	"handler/pkg/types"
)

func SendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func SendSuccess(w http.ResponseWriter, message string, data interface{}, status int) {
	response := types.Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	SendJSON(w, response, status)
}

func SendError(w http.ResponseWriter, message string, status int) {
	response := types.Response{
		Success: false,
		Message: "Error",
		Error:   message,
	}
	SendJSON(w, response, status)
}

func SendPaginated(w http.ResponseWriter, message string, data interface{}, meta types.PaginationMeta, status int) {
	response := types.PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	SendJSON(w, response, status)
}
