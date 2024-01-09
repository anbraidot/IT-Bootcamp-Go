package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	// Read the products from the file
	Products, err := readProducts("02-GET/docs/products.json")
	if err != nil {
		fmt.Println("Error reading the products: ", err)
		return
	}

	// Create a new router
	router := chi.NewRouter()

	// Ping route
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// Products route
	router.Route("/products", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Products)
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			idParam := chi.URLParam(r, "id")

			// Parse the idParam to int64
			id, err := strconv.ParseInt(idParam, 0, 0)
			if err != nil {
				// Return Error 500 Internal Server Error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Search for the product
			for _, product := range Products {
				if int64(product.Id) == id {
					json.NewEncoder(w).Encode(product)
					return
				}
			}

			// Return Error 404 Not Found
			w.WriteHeader(http.StatusNotFound)
		})

		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			priceParam := r.URL.Query().Get("price")

			// Parse the priceParam to float64
			price, err := strconv.ParseFloat(priceParam, 64)
			if err != nil {
				// Return Error 500 Internal Server Error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Define a slice of products to store the results
			var results []Product

			// Search for the product
			for _, product := range Products {
				if product.Price > price {
					results = append(results, product)
				}
			}

			// Check if there are results
			if len(results) == 0 {
				// Return Error 404 Not Found
				w.WriteHeader(http.StatusNotFound)
			} else {
				// Return the results
				json.NewEncoder(w).Encode(results)
			}
		})
	})

	// Start the server
	http.ListenAndServe(":8080", router)
}

// readProducts reads the products from the file and returns a slice of products
func readProducts(path string) (products []Product, err error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return products, err
	}

	// Read data from the file
	data, err := io.ReadAll(file)
	if err != nil {
		return products, err
	}

	// Decode the file into a slice of products
	err = json.Unmarshal(data, &products)
	if err != nil {
		return products, err
	}

	return products, nil
}