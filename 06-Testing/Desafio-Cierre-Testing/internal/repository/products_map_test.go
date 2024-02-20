package repository_test

import (
	"06-Testing/Desafio-Cierre-Testing/internal"
	"06-Testing/Desafio-Cierre-Testing/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestProductsMap_SearchProducts tests the SearchProducts function.
func TestProductsMap_SearchProducts(t *testing.T) {
	t.Run("success - empty query in empty db", func(t *testing.T) {
		// arrange
		// - create products map
		r := repository.NewProductsMap(nil)
		// - create query
		query := internal.ProductQuery{}

		// act
		// - search products
		result, err := r.SearchProducts(query)

		// assert
		expectedResult := make(map[int]internal.Product)

		// - check results
		require.Equal(t, expectedResult, result)
		require.NoError(t, err)
	})

	t.Run("success - query by id in empty db", func(t *testing.T) {
		// arrange
		// - create products map
		r := repository.NewProductsMap(nil)
		// - create query
		query := internal.ProductQuery{
			Id: 1,
		}

		// act
		// - search products
		result, err := r.SearchProducts(query)

		// assert
		expectedResult := make(map[int]internal.Product)

		// - check results
		require.Equal(t, expectedResult, result)
		require.NoError(t, err)
	})

	t.Run("success - query by id", func(t *testing.T) {
		// arrange
		// - create products map
		db := make(map[int]internal.Product)
		// - insert products
		db[1] = internal.Product{
			Id: 1,
			ProductAttributes: internal.ProductAttributes{
				Description: "description 1",
				Price:       100,
				SellerId:    1,
			},
		}
		db[2] = internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "description 2",
				Price:       200,
				SellerId:    2,
			},
		}
		// - create repository
		r := repository.NewProductsMap(db)
		// - create query
		query := internal.ProductQuery{
			Id: 2,
		}

		// act
		// - search products
		result, err := r.SearchProducts(query)

		// assert
		expectedResult := make(map[int]internal.Product)
		expectedResult[2] = internal.Product{
			Id: 2,
			ProductAttributes: internal.ProductAttributes{
				Description: "description 2",
				Price:       200,
				SellerId:    2,
			},
		}

		// - check results
		require.Equal(t, expectedResult, result)
		require.Len(t, result, 1)
		require.NoError(t, err)
	})
}
