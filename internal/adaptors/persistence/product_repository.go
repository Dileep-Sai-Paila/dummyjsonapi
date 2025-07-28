package persistence

import (
	"dummyjsonapi/internal/core/models"
	"dummyjsonapi/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductRepository interface {
	FindAll() (*models.ProductList, error)
	FindByCategory(category string) (*models.ProductList, error)
	FindAllCategories() ([]models.Category, error)
}

type productAPIRepository struct {
	apiURL string
}

func NewProductAPIRepository(url string) ProductRepository {
	return &productAPIRepository{apiURL: url}
}

func (r *productAPIRepository) FindAll() (*models.ProductList, error) {
	fullURL := fmt.Sprintf("%s/products", r.apiURL)

	resp, err := utils.MakeGETRequest(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making request to dummyjson api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dummyjson api returned non-200 status: %d", resp.StatusCode)
	}

	var productList models.ProductList
	if err := json.NewDecoder(resp.Body).Decode(&productList); err != nil {
		return nil, fmt.Errorf("error decoding response from dummyjson api: %w", err)
	}

	return &productList, nil
}

func (r *productAPIRepository) FindByCategory(category string) (*models.ProductList, error) {
	fullURL := fmt.Sprintf("%s/products/category/%s", r.apiURL, category)

	resp, err := utils.MakeGETRequest(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making request to dummyjson api for category %s: %w", category, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dummyjson api returned non-200 status for category %s: %d", category, resp.StatusCode)
	}

	var productList models.ProductList
	if err := json.NewDecoder(resp.Body).Decode(&productList); err != nil {
		return nil, fmt.Errorf("error decoding response from dummyjson api for category %s: %w", category, err)
	}

	return &productList, nil
}

func (r *productAPIRepository) FindAllCategories() ([]models.Category, error) {
	fullURL := fmt.Sprintf("%s/products/categories", r.apiURL)

	resp, err := utils.MakeGETRequest(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making request for categories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dummyjson api returned non-200 status for categories: %d", resp.StatusCode)
	}

	var categories []models.Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, fmt.Errorf("error decoding categories response: %w", err)
	}

	return categories, nil
}
