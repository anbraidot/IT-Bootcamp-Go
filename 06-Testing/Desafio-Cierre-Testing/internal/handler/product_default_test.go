package handler_test

import (
	"06-Testing/Desafio-Cierre-Testing/internal"
	"06-Testing/Desafio-Cierre-Testing/internal/handler"
	"06-Testing/Desafio-Cierre-Testing/internal/repository"
	"context"
	"errors"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// TestProductsDefault_Get tests the Get method
func TestProductsDefault_Get(t *testing.T) {
	t.Run("success - empty query", func(t *testing.T) {
		// arrange
		// - create products
		products := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {
				Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		// - create repository mock
		rm := repository.NewProductRepositoryMock()
		// - SearchProducts mock function
		rm.On("SearchProducts", internal.ProductQuery{}).Return(products, nil)
		// - create handler
		hd := handler.NewProductsDefault(rm)
		// - handler func
		handlerFunc := hd.Get()
		// - request
		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		handlerFunc(res, req)

		// assert
		expectedBody := `{
			"data": {
					"1":{"description":"product 1","id":1,"price":100,"seller_id":1},
					"2":{"description":"product 2","id":2,"price":200,"seller_id":2}
				},
				"message":"success"
			}`
		expectedCode := http.StatusOK
		// - check results
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		// - check mock
		rm.AssertExpectations(t)
	})

	t.Run("success - query by id", func(t *testing.T) {
		// arrange
		// - create products
		products := map[int]internal.Product{
			1: {
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "product 1",
					Price:       100,
					SellerId:    1,
				},
			},
		}
		// - create repository mock
		rm := repository.NewProductRepositoryMock()
		// - SearchProducts mock function
		rm.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(products, nil)
		// - create handler
		hd := handler.NewProductsDefault(rm)
		// - handler func
		handlerFunc := hd.Get()
		// - request
		req := httptest.NewRequest(http.MethodGet, "/product?id=1", nil)
		// -- chi context
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// act
		handlerFunc(res, req)

		// assert
		expectedBody := `{
			"data": {	
					"1":{"description":"product 1","id":1,"price":100,"seller_id":1}
				},
				"message":"success"
			}`
		expectedCode := http.StatusOK
		// - check results
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		// - check mock
		rm.AssertExpectations(t)
	})

	t.Run("error - invalid id", func(t *testing.T) {
		// arrange
		// - create repository
		rp := repository.NewProductsMap(nil)
		// - create handler
		h := handler.NewProductsDefault(rp)
		// - handler func
		handlerFunc := h.Get()
		// - request
		req := httptest.NewRequest(http.MethodGet, "/product?id=a", nil)
		// -- chi context
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// act
		handlerFunc(res, req)

		// assert
		expectedBody := `{"status":"Bad Request","message":"invalid id"}`
		expectedCode := http.StatusBadRequest
		// - check results
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("error - internal error", func(t *testing.T) {
		// arrange
		// - create repository mock
		rm := repository.NewProductRepositoryMock()
		// - create handler
		h := handler.NewProductsDefault(rm)
		// - SearchProducts mock function
		rm.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(make(map[int]internal.Product), errors.New("internal error"))
		// - handler func
		handlerFunc := h.Get()
		// - request
		req := httptest.NewRequest(http.MethodGet, "/product?id=1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// act
		// - call handler func
		handlerFunc(res, req)

		// assert
		expectedBody := `{"status":"Internal Server Error","message":"internal error"}`
		expectedCode := http.StatusInternalServerError
		// - check results
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
