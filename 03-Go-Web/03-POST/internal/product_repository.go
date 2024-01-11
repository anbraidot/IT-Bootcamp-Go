package internal

import "errors"

var (
	// ErrProductCodeAlreadyExists is the error for when the product code already exists
	ErrProductCodeAlreadyExists = errors.New("product code already exists")
)

// ProductRepository is the interface for the product repository
type ProductRepository interface {
	// Save saves the product in the repository
	Save(product *Product) (err error)
	// GetAll gets all the products from the repository
	GetAll() (products []*Product, err error)
	// GetByID gets the product by id from the repository
	GetByID(id int) (product *Product, err error)
}
