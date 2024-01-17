package application

import (
	"03-Code-Review-1/internal/handler"
	"03-Code-Review-1/internal/loader"
	"03-Code-Review-1/internal/repository"
	"03-Code-Review-1/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/", hd.GetAll())
		// - GET /vehicles/{id}
		rt.Get("/{id}", hd.GetById())
		// - GET /vehicles/dimensions?height={min_height}-{max_height}&width={min_width}-{max_width}
		rt.Get("/dimensions", hd.GetByDimensions())
		// - GET /vehicles/weight?min={min_weight}&max={max_weight}
		rt.Get("/weight", hd.GetByWeight())
		// - GET /vehicles/color/{color}/year/{year}
		rt.Get("/color/{color}/year/{year}", hd.GetByColorAndYear())
		// - GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetByBrandAndYears())
		// - GET /vehicles/average_speed/brand/{brand}
		rt.Get("/average_speed/brand/{brand}", hd.GetAverageSpeedByBrand())
		// - GET /vehicles/average_capacity/brand/{brand}
		rt.Get("/average_capacity/brand/{brand}", hd.GetAverageCapacityByBrand())
		// - GET /vehicles/fuel_type/{type}
		rt.Get("/fuel_type/{type}", hd.GetByFuelType())
		// - GET /vehicles/transmission/{type}
		rt.Get("/transmission/{type}", hd.GetByTransmission())
		// - POST /vehicles
		rt.Post("/", hd.Create())
		// - POST /vehicles/batch
		rt.Post("/batch", hd.CreateInBatch())
		// - PUT /vehicles/{id}/update_speed
		rt.Put("/{id}/update_speed", hd.UpdateSpeed())
		// - PUT /vehicles/{id}/update_fuel
		rt.Put("/{id}/update_fuel", hd.UpdateFuelType())
		// - DELETE /vehicles/{id}
		rt.Delete("/{id}", hd.Delete())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
