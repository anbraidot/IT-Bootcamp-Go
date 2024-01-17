package handler

import (
	"04-Desafio/internal/service"
	"04-Desafio/platform/web/response"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewTickerDefault(sv service.ServiceTicketDefault) *HandlerTicketDefault {
	return &HandlerTicketDefault{
		sv: sv,
	}
}

// HandlerTicketDefault represents the default handler of the tickets
type HandlerTicketDefault struct {
	// sv represents the service of the tickets
	sv service.ServiceTicketDefault
}

// GetTotalTicketsByDestinationCountry returns the total number of tickets by destination country
func (h *HandlerTicketDefault) GetTotalTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		//- get destination country from url
		country := chi.URLParam(r, "dest")
		if country == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		//process
		tickets, err := h.sv.GetTicketByDestinationCountry(context.TODO(), country)
		if err != nil {
			switch {
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		//response
		response.JSON(w, http.StatusOK, tickets)
	}
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets by destination country
func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		country := chi.URLParam(r, "dest")
		if country == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		//process
		percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(context.TODO(), country)
		if err != nil {
			switch {
			default:
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		//response
		response.JSON(w, http.StatusOK, percentage)
	}
}
