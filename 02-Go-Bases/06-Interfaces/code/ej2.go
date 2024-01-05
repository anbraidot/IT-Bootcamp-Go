package main

import "fmt"

// Interface Product that represents a product
type Product interface {
	Price() float32
}

// SmallProduct is a struct that represents a small product
type SmallProduct struct {
	price float32
}

// MediumProduct is a struct that represents a medium product
type MediumProduct struct {
	price float32
}

// LargeProduct is a struct that represents a large product
type LargeProduct struct {
	price float32
}

// Price returns the price of the small product
func (s *SmallProduct) Price() float32 {
	return s.price
}

// Price returns the price of the medium product
func (m *MediumProduct) Price() float32 {
	return m.price * 1.3
}

// Price returns the price of the large product
func (l *LargeProduct) Price() float32 {
	return l.price * 1.6 + 2500.00
}

// Factory returns a product based on the type and price
func Factory(typeProduct string, price float32) Product {

	switch typeProduct {
	case "Small":
		return &SmallProduct{
			price: price,
		}
	case "Medium":
		return &MediumProduct{
			price: price,
		}
	case "Large":
		return &LargeProduct{
			price: price,
		}
	}
	return nil
}

func main() {
	result := Factory("Small", 1000.00)
	fmt.Printf("Price Small Product: $%f\n", result.Price())

	result = Factory("Medium", 1000.00)
	fmt.Printf("Price Medium Product: $%f\n", result.Price())

	result = Factory("Large", 1000.00)
	fmt.Printf("Price Large Product: $%f\n", result.Price())
}