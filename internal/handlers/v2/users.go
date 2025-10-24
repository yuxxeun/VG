package v2

import (
	"net/http"

	"github.com/yuxxeun/VG/internal/types"
	"github.com/yuxxeun/VG/internal/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// V2 dengan format berbeda atau fitur tambahan
	users := []types.User{
		{ID: 1, Name: "John Doe V2", Email: "john@example.com", Role: "super_admin", CreatedAt: "2024-01-01"},
	}

	response := map[string]interface{}{
		"version":  "2.0",
		"users":    users,
		"features": []string{"advanced_filtering", "bulk_operations"},
	}

	utils.SendSuccess(w, "Users retrieved successfully (v2)", response, http.StatusOK)
}
