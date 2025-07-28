package usecase

import (
	"dummyjsonapi/internal/adaptors/persistence"
	"dummyjsonapi/internal/core/models"
	"log"
)

// i defined a type for the grouped response for clarity, just to avoid confusion.
type GroupedProducts map[string][]models.Product

// this interface defines the business logic operations for products.
type ProductUsecase interface {
	FetchAllProducts() (*models.ProductList, error)
	FetchProductsByCategory(category string) (*models.ProductList, error)
	FetchProductsForAllCategories() (GroupedProducts, error)
	FetchProductsByMultipleCategories(categories []string) (*models.ProductList, error)
}

type productUsecase struct {
	repo persistence.ProductRepository
}

func NewProductUsecase(repo persistence.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (u *productUsecase) FetchAllProducts() (*models.ProductList, error) {
	return u.repo.FindAll()
}

func (u *productUsecase) FetchProductsByCategory(category string) (*models.ProductList, error) {
	return u.repo.FindByCategory(category)
}

func (u *productUsecase) FetchProductsForAllCategories() (GroupedProducts, error) {
	categories, err := u.repo.FindAllCategories()
	if err != nil {
		return nil, err
	}
	productsByCategory := make(GroupedProducts)
	for _, category := range categories {
		productList, err := u.repo.FindByCategory(category.Slug)
		if err != nil {
			log.Printf("WARN: could not fetch products for category '%s': %v", category.Slug, err)
			continue
		}
		if len(productList.Products) > 0 {
			productsByCategory[category.Slug] = productList.Products
		}
	}
	return productsByCategory, nil
}

// takes in a slice of category strings, fetches products for each, and merges them into a single list.
func (u *productUsecase) FetchProductsByMultipleCategories(categories []string) (*models.ProductList, error) {
	var allProducts []models.Product

	// just a loop through the list of category names provided by the user.
	for _, category := range categories {
		productList, err := u.repo.FindByCategory(category)
		if err != nil {
			log.Printf("WARN: could not fetch products for category '%s': %v", category, err) // Log the error for the failed category but continue with the others.
			continue
		}
		// Append the found products to our master list.
		allProducts = append(allProducts, productList.Products...)
	}

	// Assemble the final ProductList to return.
	finalList := &models.ProductList{
		Products: allProducts,
		Total:    len(allProducts),
		Skip:     0,
		Limit:    len(allProducts),
	}

	return finalList, nil
}
