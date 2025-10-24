package v1

import (
	"encoding/json"
	"net/http"

	"handler/pkg/types"
	"handler/pkg/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []types.Product{
		{ID: 1, Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, Stock: 10, CreatedAt: "2024-01-01"},
		{ID: 2, Name: "Mouse", Description: "Wireless Mouse", Price: 250000, Stock: 50, CreatedAt: "2024-01-02"},
	}

	utils.SendSuccess(w, "Products retrieved successfully", products, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req types.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product := types.Product{
		ID:          3,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   "2024-01-03",
	}

	utils.SendSuccess(w, "Product created successfully", product, http.StatusCreated)
}
