package internal

import "errors"

var (
	// ErrProductNotFound is the error for when the product is not found
	ErrProductNotFound = errors.New("product not found")
	// ErrProductAlreadyExists is the error for when the product already exists
	ErrProductAlreadyExists = errors.New("product already exists")
	// ErrFieldRequired is the error for when a field is required
	ErrFieldRequired = errors.New("field is required")
	// ErrInvalidField is the error for when a field is invalid
	ErrInvalidField = errors.New("field is invalid")
)

// ProductService is the interface for the product service
type ProductService interface {
	// Save saves the product in the service
	Save(product *Product) (err error)
	// GetAll gets all the products from the service
	GetAll() (products []Product, err error)
	// GetByID gets the product by id from the service
	GetByID(id int) (product Product, err error)
	// Update updates the product in the service
	Update(product *Product) (err error)
	// Delete deletes the product from the service
	Delete(id int) (err error)
}
