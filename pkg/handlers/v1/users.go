package v1

import (
	"encoding/json"
	"net/http"

	"github.com/yuxxeun/VG/pkg/types"
	"github.com/yuxxeun/VG/pkg/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Mock data - ganti dengan database query
	users := []types.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Role: "admin", CreatedAt: "2024-01-01"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Role: "user", CreatedAt: "2024-01-02"},
		{ID: 3, Name: "Bob Wilson", Email: "bob@example.com", Role: "user", CreatedAt: "2024-01-03"},
	}

	meta := types.PaginationMeta{
		Page:       1,
		PerPage:    10,
		TotalPages: 1,
		TotalItems: len(users),
	}

	utils.SendPaginated(w, "Users retrieved successfully", users, meta, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req types.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi
	if req.Name == "" || req.Email == "" {
		utils.SendError(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	// Mock create - ganti dengan database insert
	user := types.User{
		ID:        4,
		Name:      req.Name,
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: "2024-01-04",
	}

	utils.SendSuccess(w, "User created successfully", user, http.StatusCreated)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Mock data - ganti dengan database query by ID
	user := types.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Role:      "admin",
		CreatedAt: "2024-01-01",
	}

	utils.SendSuccess(w, "User retrieved successfully", user, http.StatusOK)
}
