package handler_test

import (
	"03-POST/internal"
	"03-POST/internal/handler"
	"03-POST/internal/repository"
	"03-POST/internal/service"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// NewRequest creates a new http.Request with the given parameters
func NewRequest(method, url string, body io.Reader, header http.Header, urlParams, queryParams map[string]string) *http.Request {
	// old request
	req := httptest.NewRequest(method, url, body)

	// set the header
	req.Header = header

	// url params
	if urlParams != nil {
		chiCtx := chi.NewRouteContext()
		for key, value := range urlParams {
			chiCtx.URLParams.Add(key, value)
		}
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	}

	// query params
	if queryParams != nil {
		// set the query params
		query := req.URL.Query()
		for key, value := range queryParams {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	return req
}

func TestProductDefault_GetAll(t *testing.T) {
	// 200 ok
	t.Run("success: should return a list of all products", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.GetAll()

		//when
		//req := httptest.NewRequest("GET", "/products", nil)
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products", nil, header, nil, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusOK
		expectedBody := `[
			{"id":1,"name":"Coca-Cola","quantity":10,"code_value":"123","is_published":true,"expiration":"01/01/2025","price":10.5},
			{"id":2,"name":"Pepsi","quantity":20,"code_value":"456","is_published":true,"expiration":"01/01/2025","price":10.5}
			]`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.GetAll()

		//when
		//req := httptest.NewRequest("GET", "/products", nil)
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products", nil, header, nil, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductDefault_GetById(t *testing.T) {
	// - 200 ok
	t.Run("success: should return a product by id", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.GetByID()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products/1", nil, header, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"Coca-Cola","quantity":10,"code_value":"123","is_published":true,"expiration":"01/01/2025","price":10.5}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 400 bad request
	t.Run("error: should return a bad request when the id is not a number", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.GetByID()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products/abc", nil, header, map[string]string{"id": "abc"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status": "Bad Request", "message": "invalid id"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 404 not found
	t.Run("error: should return a not found when the id does not exist", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.GetByID()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products/0", nil, header, map[string]string{"id": "0"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusNotFound
		expectedBody := `{"status": "Not Found", "message": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.GetByID()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("GET", "/products/1", nil, header, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductDefault_Create(t *testing.T) {
	// - 201 created
	t.Run("success: should create a product", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := make(map[int]internal.Product)
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 0)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Create()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		productJson := []byte(`{"name":"Coca-Cola","quantity":10,"code_value":"123456","is_published":true,"expiration":"01/01/2025","price":10.5}`)

		req := NewRequest("POST", "/products", bytes.NewReader(productJson), header, nil, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusCreated
		expectedBody := `{"id":1,"name":"Coca-Cola","quantity":10,"code_value":"123456","is_published":true,"expiration":"01/01/2025","price":10.5}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.Create()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		productJson := []byte(`{"name":"Coca-Cola","quantity":10,"code_value":"123456","is_published":true,"expiration":"01/01/2025","price":10.5}`)

		req := NewRequest("POST", "/products", bytes.NewReader(productJson), header, nil, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductDefault_Update(t *testing.T) {
	// - 400 bad request
	t.Run("error: should return a bad request when the id is not a number", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Update()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("PUT", "/products/abc", nil, header, map[string]string{"id": "abc"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status": "Bad Request", "message": "invalid id"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 404 not found
	t.Run("error: should return a not found when the id does not exist", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Update()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		productJson := []byte(`{"id":1,"name":"Coca-Cola","quantity":10,"code_value":"123456","is_published":true,"expiration":"01/01/2025","price":10.5}`)

		req := NewRequest("PUT", "/products/0", bytes.NewReader(productJson), header, map[string]string{"id": "0"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusNotFound
		expectedBody := `{"status": "Not Found", "message": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.Update()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("PUT", "/products/1", nil, header, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductDefault_UpdatePartial(t *testing.T) {
	// - 400 bad request
	t.Run("error: should return a bad request when the id is not a number", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.UpdatePartial()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("PATCH", "/products/abc", nil, header, map[string]string{"id": "abc"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status": "Bad Request", "message": "invalid id"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 404 not found
	t.Run("error: should return a not found when the id does not exist", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.UpdatePartial()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("PATCH", "/products/0", nil, header, map[string]string{"id": "0"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusNotFound
		expectedBody := `{"status": "Not Found", "message": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.UpdatePartial()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("PATCH", "/products/1", nil, header, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductDefault_Delete(t *testing.T) {
	// - 200 ok
	t.Run("success: should delete a product by id", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Delete()

		//when
		header := http.Header{
			"Auth": []string{"a1b2c3d4"},
		}
		req := NewRequest("DELETE", "/products/1", nil, header, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusNoContent
		expectedBody := ""
		expectedHeader := http.Header{}

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 400 bad request
	t.Run("error: should return a bad request when the id is not a number", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Delete()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("DELETE", "/products/abc", nil, header, map[string]string{"id": "abc"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status": "Bad Request", "message": "invalid id"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 404 not found
	t.Run("error: should return a not found when the id does not exist", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		db := map[int]internal.Product{
			1: {Id: 1, Name: "Coca-Cola", Quantity: 10, CodeValue: "123", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
			2: {Id: 2, Name: "Pepsi", Quantity: 20, CodeValue: "456", IsPublished: true, Expiration: "01/01/2025", Price: 10.5},
		}
		// - create a product repository, service and handler
		rp := repository.NewProductMap(db, 2)
		sv := service.NewProductDefault(rp)
		hd := handler.NewProductDefault(sv)
		hdFunc := hd.Delete()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("DELETE", "/products/0", nil, header, map[string]string{"id": "0"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusNotFound
		expectedBody := `{"status": "Not Found", "message": "product not found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// - 401 unauthorized
	t.Run("error: should return unauthorized when the auth token is invalid", func(t *testing.T) {
		//given
		// - set environment variables
		setEnvironmentVariable()
		// - create a db with some products
		hd := handler.NewProductDefault(nil)
		hdFunc := hd.Delete()

		//when
		header := http.Header{
			"Content-Type": []string{"application/json"},
			//"Auth":         []string{"a1b2c3d4"},
		}
		req := NewRequest("DELETE", "/products/0", nil, header, map[string]string{"id": "0"}, nil)
		res := httptest.NewRecorder()
		hdFunc(res, req)

		//then
		expectedCode := http.StatusUnauthorized
		expectedBody := `{"status": "Unauthorized", "message": "unauthorized"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}

func setEnvironmentVariable() {
	_ = os.Setenv("AUTH_TOKEN", "a1b2c3d4")
}
