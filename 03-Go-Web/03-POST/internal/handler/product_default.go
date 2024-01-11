package handler

import (
	"03-POST/internal"
	"03-POST/platform/web/request"
	"03-POST/platform/web/response"
	"errors"
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
				"message":  "products found",
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
				return
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"product": product,
			"message": "product found",
		})
		return
	}
}
