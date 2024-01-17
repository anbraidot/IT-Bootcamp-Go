package application

import (
	"04-Desafio/internal/handler"
	"04-Desafio/internal/loader"
	"04-Desafio/internal/repository"
	"04-Desafio/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	ld := loader.NewLoaderTicketCSV(a.dbFile)
	db, err := ld.Load()
	if err != nil {
		return fmt.Errorf("error loading database: %v", err)
	}
	rp := repository.NewRepositoryTicketMap(db, len(db))
	sv := service.NewServiceTicketDefault(*rp)
	hd := handler.NewTickerDefault(*sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	(*a).rt.Get("/ticket/getByCountry/{dest}", hd.GetTotalTicketsByDestinationCountry())
	(*a).rt.Get("/ticket/getAverage/{dest}", hd.GetPercentageTicketsByDestinationCountry())

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
