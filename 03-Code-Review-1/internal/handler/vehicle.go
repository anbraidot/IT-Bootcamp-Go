package handler

import (
	"03-Code-Review-1/internal"
	"strconv"
	"strings"

	"errors"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetById is a method that returns a handler for the route GET /vehicles/{id}
func (h *VehicleDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the vehicle id is bad formatted",
			})
			return
		}

		// process
		// - get vehicle by id
		v, err := h.sv.FindById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, map[string]string{
					"message": "vehicle not found",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, nil)
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, v)
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get body
		var body VehicleJSON
		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the vehicle data is bad formatted or incomplete",
			})
			return
		}

		// process
		// - validate data
		if err := body.validate(); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}
		// - deserialize data
		vehicle := body.deserialize()
		// - create vehicle
		if err := h.sv.Create(&vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.JSON(w, http.StatusConflict, map[string]string{
					"message": "vehicle id already exists",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "vehicle was created successfully",
		})
	}
}

// CreateInBatch is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateInBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body []VehicleJSON
		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the body request is not valid",
			})
			return
		}
		// process
		// - validate data
		for _, vehicle := range body {
			if err := vehicle.validate(); err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
				return
			}
		}
		// - deserialize data
		vehicles := make([]internal.Vehicle, len(body))
		for i, vehicle := range body {
			vehicles[i] = vehicle.deserialize()
		}
		// - create vehicles in batch
		if err := h.sv.CreateInBatch(&vehicles); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleAlreadyExists):
				response.JSON(w, http.StatusConflict, map[string]string{
					"message": "vehicle id already exists",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "vehicles were created successfully",
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		color := chi.URLParam(r, "color")
		yearString := chi.URLParam(r, "year")

		// process
		// - validate data
		if yearString == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the year is required",
			})
			return
		}
		if color == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the color is required",
			})
			return
		}
		// - convert year to int
		year, err := strconv.Atoi(yearString)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the year is invalid",
			})
			return
		}
		// - get vehicles by color and year
		result, err := h.sv.FindByColorAndYear(color, int(year))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			data := make(map[int]VehicleJSON)
			for key, value := range result {
				data[key] = serialize(&value)
			}
			response.JSON(w, http.StatusOK, data)
		}
	}
}

// UpdateSpeed is a method that returns a handler for the route PUT /vehicles/{id}/update_speed
func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the id is invalid",
			})
			return
		}
		// - get body
		var body VehicleJSON
		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the vehicle data is bad formatted or incomplete",
			})
			return
		}

		// process
		// - validate data
		if err := body.validate(); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}
		if id != body.ID {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the ids are different",
			})
			return
		}
		// - deserialize data
		vehicle := body.deserialize()
		// - update vehicle
		err = h.sv.Update(&vehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, map[string]string{
					"message": "vehicle not found",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "vehicle was updated successfully",
		})
	}
}

// UpdateFuelType is a method that returns a handler for the route PUT /vehicles/{id}/update_fuel
func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the id is invalid",
			})
			return
		}
		// - get body
		var body VehicleJSON
		if err := request.JSON(r, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the vehicle data is bad formatted or incomplete",
			})
			return
		}

		// process
		// - validate data
		if err := body.validate(); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}
		if id != body.ID {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the ids are different",
			})
			return
		}
		// - deserialize data
		vehicle := body.deserialize()
		// - update vehicle
		err = h.sv.Update(&vehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, map[string]string{
					"message": "vehicle not found",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "vehicle was updated successfully",
		})
	}
}

// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
func (v *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get query params
		height := strings.Split(r.URL.Query()["height"][0], "-")
		width := strings.Split(r.URL.Query()["width"][0], "-")

		// process
		// - convert params to float64
		minHeight, err := strconv.ParseFloat(height[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_height is invalid",
			})
			return
		}
		maxHeight, err := strconv.ParseFloat(height[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the max_height is invalid",
			})
			return
		}
		minWidth, err := strconv.ParseFloat(width[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_width is invalid",
			})
			return
		}
		maxWidth, err := strconv.ParseFloat(width[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the max_width is invalid",
			})
			return
		}
		// - validate data
		if minHeight < 0 || maxHeight < 0 || minWidth < 0 || maxWidth < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the dimensions cannot be negative",
			})
			return
		}
		if minHeight > maxHeight {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_height is greater than the max_height",
			})
			return
		}
		if minWidth > maxWidth {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_width is greater than the max_width",
			})
			return
		}
		// - get vehicles by dimensions
		result, err := v.sv.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			data := make(map[int]VehicleJSON)
			for key, value := range result {
				data[key] = serialize(&value)
			}
			response.JSON(w, http.StatusOK, data)
			return
		}
	}
}

// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel_type/{type}
func (v *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		fuelType := chi.URLParam(r, "type")

		// process
		// - validate data
		if fuelType == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the fuel_type is invalid",
			})
			return
		}
		// - get vehicles by fuel type
		result, err := v.sv.FindByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			response.JSON(w, http.StatusOK, result)
			return
		}
	}
}

// GetByTransmission is a method that returns a handler for the route GET /vehicles/transmission/{type}
func (v *VehicleDefault) GetByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		transmission := chi.URLParam(r, "type")

		// process
		// - validate data
		if transmission == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the transmission is invalid",
			})
			return
		}
		// - get vehicles by transmission
		result, err := v.sv.FindByTransmission(transmission)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			response.JSON(w, http.StatusOK, result)
			return
		}
	}
}

// GetByWeight is a method that returns a handler for the route GET /vehicles/weight?min={min_weight}&max={max_weight}
func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get query params
		minWeight := r.URL.Query()["min"][0]
		maxWeight := r.URL.Query()["max"][0]

		// process
		// - convert params to float64
		minWeightFloat, err := strconv.ParseFloat(minWeight, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_weight is invalid",
			})
			return
		}
		maxWeightFloat, err := strconv.ParseFloat(maxWeight, 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the max_weight is invalid",
			})
			return
		}
		// - validate data
		if minWeightFloat < 0 || maxWeightFloat < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the weight cannot be negative",
			})
			return
		}
		if minWeightFloat > maxWeightFloat {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the min_weight is greater than the max_weight",
			})
			return
		}
		// - get vehicles by weight
		result, err := h.sv.FindByWeight(minWeightFloat, maxWeightFloat)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			data := make(map[int]VehicleJSON)
			for key, value := range result {
				data[key] = serialize(&value)
			}
			response.JSON(w, http.StatusOK, data)
			return
		}
	}
}

// GetByBrandAndYears is a method that returns a handler for the route GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")

		// process
		// - convert params to int
		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the start_year is invalid",
			})
			return
		}
		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the end_year is invalid",
			})
			return
		}
		// - validate data
		if brand == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the brand is empty",
			})
			return
		}
		if startYearInt < 0 || endYearInt < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the years cannot be negative",
			})
			return
		}
		if startYearInt > endYearInt {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the start_year must be greater than the end_year",
			})
			return
		}
		// - get vehicles by brand and years
		result, err := h.sv.FindByBrandAndYears(brand, startYearInt, endYearInt)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if len(result) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			data := make(map[int]VehicleJSON)
			for key, value := range result {
				data[key] = serialize(&value)
			}
			response.JSON(w, http.StatusOK, data)
			return
		}
	}
}

// GetAverageSpeedByBrand is a method that returns a handler for the route GET /vehicles/average-speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		brand := chi.URLParam(r, "brand")

		// process
		// validate data
		if brand == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the brand is empty",
			})
			return
		}
		// - get average speed by brand
		result, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if result == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			response.JSON(w, http.StatusOK, map[string]float64{
				"average_speed": result,
			})
			return
		}
	}
}

// GetAverageCapacityByBrand is a method that returns a handler for the route GET /vehicles/average_capacity/brand/{brand}
func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		brand := chi.URLParam(r, "brand")

		// process
		// validate data
		if brand == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the brand is empty",
			})
			return
		}
		// - get average capacity by brand
		result, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{
				"message": "internal server error",
			})
			return
		}

		// response
		if result == 0 {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"message": "no vehicles were found",
			})
			return
		} else {
			response.JSON(w, http.StatusOK, map[string]float64{
				"average_capacity": result,
			})
			return
		}
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/{id}
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get url params
		idString := chi.URLParam(r, "id")

		// process
		// - convert params to int
		id, err := strconv.Atoi(idString)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"message": "the id is invalid",
			})
			return
		}
		// - delete vehicle by id
		if err := h.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, map[string]string{
					"message": "vehicle not found",
				})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]string{
					"message": "internal server error",
				})
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "vehicle was deleted successfully",
		})
	}
}

// deserialize is a method that deserializes a vehicle from JSON format to Vehicle format
func (v *VehicleJSON) deserialize() (vehicle internal.Vehicle) {
	dimensions := internal.Dimensions{
		Height: v.Height,
		Length: v.Length,
		Width:  v.Width,
	}
	vehicleAttibutes := internal.VehicleAttributes{
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Dimensions:      dimensions,
	}
	vehicle = internal.Vehicle{
		Id:                v.ID,
		VehicleAttributes: vehicleAttibutes,
	}
	return vehicle
}

// serialize is a method that serializes a vehicle from Vehicle format to JSON format
func serialize(vehicle *internal.Vehicle) (vehicleJSON VehicleJSON) {
	vehicleJSON = VehicleJSON{
		ID:              vehicle.Id,
		Brand:           vehicle.Brand,
		Model:           vehicle.Model,
		Registration:    vehicle.Registration,
		Color:           vehicle.Color,
		FabricationYear: vehicle.FabricationYear,
		Capacity:        vehicle.Capacity,
		MaxSpeed:        vehicle.MaxSpeed,
		FuelType:        vehicle.FuelType,
		Transmission:    vehicle.Transmission,
		Weight:          vehicle.Weight,
		Height:          vehicle.Dimensions.Height,
		Length:          vehicle.Dimensions.Length,
		Width:           vehicle.Dimensions.Width,
	}
	return vehicleJSON
}

// validate is a method that validates the data of a vehicle
func (v *VehicleJSON) validate() (err error) {
	if v.ID <= 0 {
		err = errors.New("error: the id must be greater than 0")
		return
	}
	if v.Brand == "" {
		err = errors.New("error: the brand is required")
		return
	}
	if v.Model == "" {
		err = errors.New("error: the model is required")
		return
	}
	if v.Registration == "" {
		err = errors.New("error: the registration is required")
		return
	}
	if v.Color == "" {
		err = errors.New("error: the color is required")
		return
	}
	if v.FabricationYear < 1886 {
		err = errors.New("error: the fabrication year must be greater than 0")
		return
	}
	if v.Capacity <= 0 {
		err = errors.New("error: the capacity must be greater than 0")
		return
	}
	if v.MaxSpeed <= 0 {
		err = errors.New("error: the max speed must be greater than 0")
		return
	}
	if v.FuelType == "" {
		err = errors.New("error: the fuel type is required")
		return
	}
	return
}
