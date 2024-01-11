package repository

import "03-POST/internal"

// NewProductMap creates a new product map
func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {
	// default config
	if db == nil {
		db = make(map[int]internal.Product)
	}
	if lastId == 0 {
		lastId = 1
	}

	// return the product map
	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

// ProductMap is the struct for the product map
type ProductMap struct {
	//db is the map of products
	//- key is the id of the product
	//- value is the product
	db map[int]internal.Product
	//lastId is the last id of the product
	lastId int
}

func (p *ProductMap) Save(product *internal.Product) (err error) {
	// validate the product
	// code_value is unique
	if err = p.validate(product); err != nil {
		return
	}

	// increment the last id
	(*p).lastId++
	// set the id of the product
	(*product).Id = (*p).lastId

	// store the product in the map
	(*p).db[(*product).Id] = *product

	return
}

// GetAll gets all the products from the repository
func (p *ProductMap) GetAll() (products []*internal.Product, err error) {
	// get all the products from the map
	for _, p := range (*p).db {
		products = append(products, &p)
	}
	return
}

// GetByID gets the product by id from the repository
func (p *ProductMap) GetByID(id int) (product *internal.Product, err error) {
	// get the product from the map
	data, ok := (*p).db[id]
	product = &data
	if !ok {
		err = internal.ErrProductNotFound
		return 
	}
	return
}

// validate validates the product fields with business rules
func (p *ProductMap) validate(product *internal.Product) (err error) {
	// validate the product
	// code_value is unique
	for _, p := range (*p).db {
		if (*product).CodeValue == p.CodeValue {
			err = internal.ErrProductCodeAlreadyExists
			return
		}
	}
	return
}
