package repository

import (
	"03-POST/internal"
)

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
	if err = p.ValidateIdExistance((*product).Id); err != nil {
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
func (p *ProductMap) GetAll() (products []internal.Product, err error) {
	// get all the products from the map
	for _, p := range (*p).db {
		products = append(products, p)
	}
	return
}

// GetByID gets the product by id from the repository
func (p *ProductMap) GetByID(id int) (product internal.Product, err error) {
	// get the product from the map
	product, ok := (*p).db[id]
	if !ok {
		err = internal.ErrProductIdNotFound
		return
	}
	return
}

// Update updates the product in the repository
func (p *ProductMap) Update(product *internal.Product) (err error) {
	// update the product in the map
	data, ok := (*p).db[(*product).Id]
	if !ok {
		err = internal.ErrProductIdNotFound
		return
	}
	(*p).db[data.Id] = *product
	return
}

// Delete deletes the product from the repository
func (p *ProductMap) Delete(id int) (err error) {
	// validate if the id exists
	if _, ok := (*p).db[id]; !ok {
		err = internal.ErrProductIdNotFound
		return
	} else {
		// delete the product from the map
		delete((*p).db, id)
	}
	return
}

// ValidateIdExistance validates if the id exists
func (p *ProductMap) ValidateIdExistance(id int) (err error) {
	// validate the id
	_, ok := (*p).db[id]
	if !ok {
		err = internal.ErrProductIdNotFound
		return
	}
	return
}