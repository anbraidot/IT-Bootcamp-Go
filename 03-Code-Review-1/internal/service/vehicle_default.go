package service

import (
	"03-Code-Review-1/internal"
	"errors"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a new vehicle
func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdIsAllreadyInUse):
			return internal.ErrVehicleAlreadyExists
		}
	}
	return
}

// CreateInBatch is a method that creates new vehicles in batch
func (s *VehicleDefault) CreateInBatch(v *[]internal.Vehicle) (err error) {
	err = s.rp.CreateInBatch(v)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdIsAllreadyInUse):
			return internal.ErrVehicleAlreadyExists
		}
	}
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByColorAndYear(color, year)
	return
}

// Update is a method that updates a vehicle
func (s *VehicleDefault) Update(v *internal.Vehicle) (err error) {
	// find vehicle by id
	_, err = s.rp.FindById(v.Id)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdNotFound):
			return internal.ErrVehicleNotFound
		}
	}
	// update vehicle
	err = s.rp.Update(v)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdNotFound):
			return internal.ErrVehicleNotFound
		}
	}
	return
}

// FindById is a method that returns a vehicle by id
func (s *VehicleDefault) FindById(id int) (v internal.Vehicle, err error) {
	v, err = s.rp.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdNotFound):
			return v, internal.ErrVehicleNotFound
		}
	}
	return
}

// FindByFuelType is a method that returns a map of vehicles filtered by fuel type
func (s *VehicleDefault) FindByFuelType(fuelType string) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByFuelType(fuelType)
	return
}

// FindByTransmission is a method that returns a map of vehicles filtered by transmission
func (s *VehicleDefault) FindByTransmission(transmission string) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByTransmission(transmission)
	return
}

// FindByBrand is a method that returns a map of vehicles filtered by brand
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (avg float64, err error) {
	// get vehicles by brand
	result, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}
	// calculate average speed
	if len(result) == 0 {
		return
	}
	for _, value := range result {
		avg += value.MaxSpeed
	}
	avg /= float64(len(result))

	return
}

// GetAverageCapacityByBrand is a method that returns the average capacity of vehicles by brand
func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (avg float64, err error) {
	// get vehicles by brand
	result, err := s.rp.FindByBrand(brand)
	if err != nil {
		return
	}

	// calculate average capacity
	if len(result) == 0 {
		return
	}
	for _, value := range result {
		avg += float64(value.Capacity)
	}
	avg /= float64(len(result))

	return
}

// FindByDimensions is a method that returns a map of vehicles filtered by dimensions
func (s *VehicleDefault) FindByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
	return
}

// FindByWeight is a method that returns a map of vehicles filtered by weight
func (s *VehicleDefault) FindByWeight(minWeight, maxWeight float64) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByWeight(minWeight, maxWeight)
	return
}

// FindByBrandAndYears is a method that returns a map of vehicles filtered by brand and years
func (s *VehicleDefault) FindByBrandAndYears(brand string, minYear, maxYear int) (result map[int]internal.Vehicle, err error) {
	result, err = s.rp.FindByBrandAndYears(brand, minYear, maxYear)
	return
}

// Delete is a method that deletes a vehicle
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrVehicleIdNotFound):
			return internal.ErrVehicleNotFound
		}
	}
	return
}
