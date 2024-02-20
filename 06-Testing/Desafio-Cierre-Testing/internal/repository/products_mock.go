package repository

import (
	"06-Testing/Desafio-Cierre-Testing/internal"

	"github.com/stretchr/testify/mock"
)

// ProductRepositoryMock is an struct that contains the mock of the repository.
type ProductRepositoryMock struct {
	// Mock is the mock of the repository.
	mock.Mock
	// SearchProducts is the mock of the SearchProducts function.
	FuncSearchProducts func(query internal.ProductQuery) (map[int]internal.Product, error)
}

func NewProductRepositoryMock() *ProductRepositoryMock {
	return &ProductRepositoryMock{}
}

func (m *ProductRepositoryMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
