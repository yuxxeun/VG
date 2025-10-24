package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Handler adalah entry point untuk Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Router sederhana
	switch r.URL.Path {
	case "/", "/api":
		handleHome(w, r)
	case "/api/users":
		handleUsers(w, r)
	case "/api/health":
		handleHealth(w, r)
	default:
		handleNotFound(w, r)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Welcome to Go REST API on Vercel",
		Data: map[string]string{
			"version": "1.0.0",
			"status":  "running",
		},
	}
	sendJSON(w, response, http.StatusOK)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		}
		response := Response{
			Message: "Users retrieved successfully",
			Data:    users,
		}
		sendJSON(w, response, http.StatusOK)

	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		user.ID = 3 // Simulasi ID baru
		response := Response{
			Message: "User created successfully",
			Data:    user,
		}
		sendJSON(w, response, http.StatusCreated)

	default:
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Service is healthy",
		Data: map[string]bool{
			"healthy": true,
		},
	}
	sendJSON(w, response, http.StatusOK)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	sendError(w, "Route not found", http.StatusNotFound)
}

func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Message: message})
}
