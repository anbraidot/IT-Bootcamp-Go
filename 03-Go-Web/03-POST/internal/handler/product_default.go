package handler

import (
	"03-POST/internal"
	"03-POST/platform/web/request"
	"03-POST/platform/web/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewProductDefault creates a new instance of ProductDefault
func NewProductDefault(sv internal.ProductService) *ProductDefault {
	// return the product default
	return &ProductDefault{
		sv: sv,
	}
}

// ProductDefault is the struct for the product handler
type ProductDefault struct {
	// sv is the product service
	sv internal.ProductService
}

// ProductJSON is the struct for the product JSON
type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// BodyRequestProductJSON is the struct for the body request product JSON
type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Create creates a new product
func (p *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}

		// proccess
		//- serialize the product
		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		//- save the product
		if err := p.sv.Save(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrProductAlreadyExists):
				response.Error(w, http.StatusConflict, "product code already exists")
				return
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrInvalidField):
				response.Error(w, http.StatusBadRequest, "invalid body request")
				return
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}

		// response
		//- serialize the product
		productJSON := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		//- send the product
		response.JSON(w, http.StatusCreated, map[string]any{
			"product": productJSON,
			"message": "product created",
		})
	}
}

// GetAll gets all products
func (p *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request

		//proccess
		productList, err := p.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		//response
		if len(productList) == 0 {
			response.JSON(w, http.StatusNoContent, map[string]any{
				"message": "no products found",
			})
			return
		} else {
			response.JSON(w, http.StatusOK, map[string]any{
				"products": productList,
			})
			return
		}
	}
}

// GetByID gets a product by id
func (p *ProductDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		// Get the id from the URL
		idParam := chi.URLParam(r, "id")
		// Parse the idParam to int64
		id, err := strconv.ParseInt(idParam, 0, 0)
		if err != nil {
			// Return Error 500 Internal Server Error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//proccess
		product, err := p.sv.GetByID(int(id))
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"product": product,
		})
	}
}

// Update updates a product
func (p *ProductDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - get the bytes from the body request
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}
		// - unmarshal the bytes to a map[string]any
		var mp map[string]any
		if err := json.Unmarshal(bytes, &mp); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}
		// - validate that body request contains all fields
		if err := ValidateKeyExistance(mp, "name", "quantity", "code_value", "is_published", "expiration", "price"); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}
		// - unmarshal the bytes to a BodyRequestProductJSON
		var body BodyRequestProductJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}

		// proccess
		// - serialize the product
		product := internal.Product{
			Id:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		// - save the product
		if err := p.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrInvalidField):
				response.Error(w, http.StatusBadRequest, "invalid body request")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize the product
		productJSON := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - send the product
		response.JSON(w, http.StatusAccepted, map[string]any{
			"product": productJSON,
			"message": "product updated",
		})
	}
}

// UpdatePartial updates partially a product
func (p *ProductDefault) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - get the product by id
		product, err := p.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// proccess
		// - serialize the product to a BodyRequestProductJSON
		reqBody := BodyRequestProductJSON{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		// - get the body request
		if err := request.JSON(r, &reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body request")
			return
		}
		// - serialize the BodyRequestProductJSON to a product
		product = internal.Product{
			Id:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}
		// - update the product in the database
		if err := p.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrInvalidField):
				response.Error(w, http.StatusBadRequest, "invalid body request")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize the product to a ProductJSON
		productJSON := ProductJSON{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - send the product
		response.JSON(w, http.StatusAccepted, map[string]any{
			"product": productJSON,
			"message": "product updated",
		})
	}
}

// Delete deletes a product
func (p *ProductDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// proccess
		// - delete the product by id
		if err := p.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusAccepted, map[string]any{
			"message": "product deleted",
		})
	}
}

// ValidateKeyExistance validates that a map contains all keys
func ValidateKeyExistance(mp map[string]any, keys ...string) error {
	for _, key := range keys {
		if _, ok := mp[key]; !ok {
			return fmt.Errorf("key %s not found", key)
		}
	}
	return nil
}
