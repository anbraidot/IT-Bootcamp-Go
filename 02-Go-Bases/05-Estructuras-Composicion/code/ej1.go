package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{}

func (p Product) Save() (err string) {
	Products = append(Products, p)
	return
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) (p Product, err string) {
	for _, product := range Products {
		if product.ID == id {
			return product, ""
		}
	}
	return Product{}, "Product not found"
}

func main() {

	p1 := Product{
		ID:          1,
		Name:        "Product 1",
		Price:       100.00,
		Description: "Product 1 Description",
		Category:    "Category 1",
	}

	p1.Save()

	p2 := Product{
		ID:          2,
		Name:        "Product 2",
		Price:       200.00,
		Description: "Product 2 Description",
		Category:    "Category 2",
	}

	p2.Save()

	p1.GetAll()

	p3, err := getById(1)
	fmt.Println("p3: ", p3, "err: ", err)

	p4, err := getById(4)
	fmt.Println("p4: ", p4, "err: ", err)

}
