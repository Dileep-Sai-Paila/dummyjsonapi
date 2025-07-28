package rest

import (
	"dummyjsonapi/internal/core/models"
	"dummyjsonapi/internal/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: uc}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.usecase.FetchAllProducts()
	if err != nil {
		log.Printf("ERROR: fetching all products: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetByCategory(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	if category == "" {
		http.Error(w, "Category parameter is required", http.StatusBadRequest)
		return
	}
	products, err := h.usecase.FetchProductsByCategory(category)
	if err != nil {
		log.Printf("ERROR: fetching products for category %s: %v", category, err)
		http.Error(w, "Failed to fetch products for the specified category", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetByAllCategories(w http.ResponseWriter, r *http.Request) {
	groupedProducts, err := h.usecase.FetchProductsForAllCategories()
	if err != nil {
		log.Printf("ERROR: fetching products for all categories: %v", err)
		http.Error(w, "Failed to fetch products for all categories", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groupedProducts); err != nil {
		log.Printf("ERROR: encoding grouped products to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// HTTP handler for the POST /products/by-categories endpoint.
func (h *ProductHandler) GetByMultipleCategories(w http.ResponseWriter, r *http.Request) {
	// step 1: to decode the JSON request body into my CategoriesRequest struct.
	var req models.CategoriesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// step 2: to check if the user provided any categories.
	if len(req.Categories) == 0 {
		http.Error(w, "Categories array cannot be empty", http.StatusBadRequest)
		return
	}

	// Step 3: to call the usecase with the list of categories from the request.
	products, err := h.usecase.FetchProductsByMultipleCategories(req.Categories)
	if err != nil {
		log.Printf("ERROR: fetching products for multiple categories: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	// Step 4: to send the response back to the client.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Printf("ERROR: encoding products to JSON for multiple categories: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
