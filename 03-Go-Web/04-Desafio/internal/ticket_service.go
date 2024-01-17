package internal

import (
	"context"
	"errors"
)

var (
	// ErrGettingTickets represents an error getting the tickets
	ErrGettingTickets = errors.New("error getting the tickets")
)

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets() (total int, err error)

	// GetTicketByDestinationCountry returns the tickets filtered by destination country
	GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	// ...

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (percentage float64, err error)
}