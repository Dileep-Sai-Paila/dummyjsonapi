package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"dummyjsonapi/internal/adaptors/persistence"
	"dummyjsonapi/internal/interfaces/input/api/rest"
	"dummyjsonapi/internal/usecase"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	const dummyJsonURL = "https://dummyjson.com"
	productRepo := persistence.NewProductAPIRepository(dummyJsonURL)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := rest.NewProductHandler(productUsecase)

	// API Routes
	r.Get("/products/all", productHandler.GetAll)
	r.Get("/products/category/{category}", productHandler.GetByCategory)
	r.Get("/products/categories", productHandler.GetByAllCategories)

	//  POST handler method.
	r.Post("/products/by-categories", productHandler.GetByMultipleCategories) //for multiple categories you want to ask

	// Starting the Server
	log.Println("All endpoints are live. Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
