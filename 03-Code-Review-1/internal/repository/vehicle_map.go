package repository

import "03-Code-Review-1/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a new vehicle
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {
	// validate that the id is not in use
	if _, ok := r.db[v.Id]; ok {
		err = internal.ErrVehicleIdIsAllreadyInUse
		return
	}

	// add vehicle to db
	r.db[v.Id] = *v
	return
}

// CreateInBatch is a method that creates new vehicles in batch
func (r *VehicleMap) CreateInBatch(v *[]internal.Vehicle) (err error) {
	// add vehicles to db
	for _, value := range *v {
		err = r.Create(&value)
		if err != nil {
			return
		}
	}
	return
}

// FindByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)

	// search vehicles by color and year
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			result[key] = value
		}
	}
	return
}

// Update is a method that updates a vehicle
func (r *VehicleMap) Update(v *internal.Vehicle) (err error) {
	// validate that the id exists
	if _, ok := r.db[v.Id]; !ok {
		err = internal.ErrVehicleIdNotFound
		return
	}

	// update vehicle in db
	r.db[v.Id] = *v
	return
}

// FindById is a method that returns a vehicle by id
func (r *VehicleMap) FindById(id int) (v internal.Vehicle, err error) {
	_, ok := r.db[id]
	if !ok {
		err = internal.ErrVehicleIdNotFound
		return
	}
	v = r.db[id]
	return
}

// FindByBrand is a method that returns a map of vehicles filtered by brand
func (r *VehicleMap) FindByBrand(brand string) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if value.Brand == brand {
			result[key] = value
		}
	}
	return
}

// FindByFuelType is a method that returns a map of vehicles filtered by fuel type
func (r *VehicleMap) FindByFuelType(fuelType string) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if value.FuelType == fuelType {
			result[key] = value
		}
	}
	return
}

// FindByTransmission is a method that returns a map of vehicles filtered by transmission
func (r *VehicleMap) FindByTransmission(transmission string) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if value.Transmission == transmission {
			result[key] = value
		}
	}
	return
}

// FindByDimensions is a method that returns a map of vehicles filtered by dimensions
func (r *VehicleMap) FindByDimensions(minHeight, maxHeight, minWidth, maxWidth float64) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)

	// filter vehicles by dimensions
	for key, value := range r.db {
		if value.Height >= minHeight && value.Height <= maxHeight && value.Width >= minWidth && value.Width <= maxWidth {
			result[key] = value
		}
	}
	return
}

// FindByWeight is a method that returns a map of vehicles filtered by weight
func (r *VehicleMap) FindByWeight(minWeight, maxWeight float64) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)

	// filter vehicles by weight
	for key, value := range r.db {
		if value.Weight >= minWeight && value.Weight <= maxWeight {
			result[key] = value
		}
	}
	return
}

// FindByBrandAndYears is a method that returns a map of vehicles filtered by brand and years
func (r *VehicleMap) FindByBrandAndYears(brand string, minYear, maxYear int) (result map[int]internal.Vehicle, err error) {
	result = make(map[int]internal.Vehicle)
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= minYear && value.FabricationYear <= maxYear {
			result[key] = value
		}
	}
	return
}

// Delete is a method that deletes a vehicle by id
func (r *VehicleMap) Delete(id int) (err error) {
	// validate that the id exists
	if _, ok := r.db[id]; !ok {
		err = internal.ErrVehicleIdNotFound
		return
	}

	// delete vehicle from db
	delete(r.db, id)
	return
}
