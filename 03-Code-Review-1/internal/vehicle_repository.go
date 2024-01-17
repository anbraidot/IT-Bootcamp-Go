package internal

import "errors"

// list of errors
var (
	// ErrVehicleAlreadyExists is an error that represents that the vehicle already exists
	ErrVehicleIdIsAllreadyInUse = errors.New("the vehicle id is allready in use")
	// ErrVehicleIdNotFound is an error that represents that the vehicle id was not found
	ErrVehicleIdNotFound = errors.New("the vehicle id was not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindById is a method that returns a vehicle by id
	FindById(id int) (v Vehicle, err error)
	// FindByFuelType is a method that returns a map of vehicles filtered by fuel type
	FindByFuelType(fuelType string) (result map[int]Vehicle, err error)
	// FindByTransmission is a method that returns a map of vehicles filtered by transmission
	FindByTransmission(transmission string) (result map[int]Vehicle, err error)
	// FindByColorAndYear is a method that returns a map of vehicles by color and brand
	FindByColorAndYear(color string, year int) (result map[int]Vehicle, err error)
	// FindByDimensions is a method that returns a map of vehicles filtered by dimensions
	FindByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (result map[int]Vehicle, err error)
	// FindByWeight is a method that returns a map of vehicles filtered by weight
	FindByWeight(minWeight, maxWeight float64) (result map[int]Vehicle, err error)
	// FindByBrandAndYears is a method that returns a map of vehicles filtered by brand and years
	FindByBrandAndYears(brand string, minYear, maxYear int) (result map[int]Vehicle, err error)
	// FindByBrand is a method that returns a map of vehicles filtered by brand
	FindByBrand(brand string) (result map[int]Vehicle, err error)
	// Create is a method that creates a new vehicle
	Create(v *Vehicle) (err error)
	// CreateInBatch is a method that creates new vehicles in batch
	CreateInBatch(v *[]Vehicle) (err error)
	// Update is a method that updates a vehicle
	Update(v *Vehicle) (err error)
	// Delete is a method that deletes a vehicle
	Delete(id int) (err error)
}
