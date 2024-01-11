package application

import (
	"03-POST/internal/handler"
	"03-POST/internal/repository"
	"03-POST/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// NewDefaultHTTP creates a new default HTTP
func NewDefaultHTTP(address string) *DefaultHTTP {
	// return the default HTTP
	return &DefaultHTTP{
		address: address,
	}
}

// DefaultHTTP is the struct for the default HTTP
type DefaultHTTP struct {
	// address is the address of the HTTP
	address string
}

// Run runs the HTTP server
func (d *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// create the product repository
	rp := repository.NewProductMap(nil, 0)
	// create the product service
	sv := service.NewProductDefault(rp)
	// create the product handler
	hd := handler.NewProductDefault(sv)
	// create the chi router
	rt := chi.NewRouter()

	// register the routes
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetByID())

		// r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		// 	priceParam := r.URL.Query().Get("price")

		// 	// Parse the priceParam to float64
		// 	price, err := strconv.ParseFloat(priceParam, 64)
		// 	if err != nil {
		// 		// Return Error 500 Internal Server Error
		// 		w.WriteHeader(http.StatusInternalServerError)
		// 		return
		// 	}

		// 	// Define a slice of products to store the results
		// 	var results []Product

		// 	// Search for the product
		// 	for _, product := range Products {
		// 		if product.Price > price {
		// 			results = append(results, product)
		// 		}
		// 	}

		// 	// Check if there are results
		// 	if len(results) == 0 {
		// 		// Return Error 404 Not Found
		// 		w.WriteHeader(http.StatusNotFound)
		// 	} else {
		// 		// Return the results
		// 		json.NewEncoder(w).Encode(results)
		// 	}
		// })

		//r.Get("/search", hd.GetByPriceGreaterThan())

		r.Post("/", hd.Create())
		r.Put("/{id}", hd.Update())
		r.Patch("/{id}", hd.UpdatePartial())
		r.Delete("/{id}", hd.Delete())
	})

	// run the HTTP server
	err = http.ListenAndServe(d.address, rt)
	return
}
