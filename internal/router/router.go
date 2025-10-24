package router

import (
	"net/http"
	"strings"

	v1 "github.com/yuxxeun/VG/internal/handlers/v1"
	v2 "github.com/yuxxeun/VG/internal/handlers/v2"
	"github.com/yuxxeun/VG/internal/middleware"
	"github.com/yuxxeun/VG/internal/utils"
)

func Route(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	// Apply middleware
	handler := middleware.CORS(middleware.Logger(func(w http.ResponseWriter, r *http.Request) {
		// Root
		if path == "/" || path == "/api" {
			handleRoot(w, r)
			return
		}

		// API V1 Routes
		if strings.HasPrefix(path, "/api/v1") {
			handleV1Routes(w, r, path, method)
			return
		}

		// API V2 Routes
		if strings.HasPrefix(path, "/api/v2") {
			handleV2Routes(w, r, path, method)
			return
		}

		// Health check
		if path == "/health" {
			handleHealth(w, r)
			return
		}

		// 404
		utils.SendError(w, "Route not found", http.StatusNotFound)
	}))

	handler(w, r)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"service": "Go REST API",
		"version": "1.0.0",
		"endpoints": map[string]interface{}{
			"v1":     "/api/v1",
			"v2":     "/api/v2",
			"health": "/health",
		},
	}
	utils.SendSuccess(w, "Welcome to API", data, http.StatusOK)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status": "healthy",
		"uptime": "running",
	}
	utils.SendSuccess(w, "Service is healthy", data, http.StatusOK)
}

func handleV1Routes(w http.ResponseWriter, r *http.Request, path, method string) {
	switch {
	case path == "/api/v1/users" && method == "GET":
		v1.GetUsers(w, r)
	case path == "/api/v1/users" && method == "POST":
		v1.CreateUser(w, r)
	case strings.HasPrefix(path, "/api/v1/users/") && method == "GET":
		v1.GetUserByID(w, r)
	case path == "/api/v1/products" && method == "GET":
		v1.GetProducts(w, r)
	case path == "/api/v1/products" && method == "POST":
		v1.CreateProduct(w, r)
	default:
		utils.SendError(w, "Route not found", http.StatusNotFound)
	}
}

func handleV2Routes(w http.ResponseWriter, r *http.Request, path, method string) {
	switch {
	case path == "/api/v2/users" && method == "GET":
		v2.GetUsers(w, r)
	default:
		utils.SendError(w, "Route not found", http.StatusNotFound)
	}
}
