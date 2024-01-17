package service

import (
	"04-Desafio/internal"
	"04-Desafio/internal/repository"
	"context"
	"fmt"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp repository.RepositoryTicketMap
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp repository.RepositoryTicketMap) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
// gets total tickets
	t, err := s.rp.Get(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("error getting the total tickets: %w", internal.ErrGettingTickets)
	}
	total = len(t)
	return
}

// GetTicketByDestinationCountry returns the tickets filtered by destination country
func (s *ServiceTicketDefault) GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	// gets tickets by destination country
	t, err = s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return nil, fmt.Errorf("error getting the tickets by destination country: %w", internal.ErrGettingTickets)
	}
	return
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (percentage float64, err error) {
	// gets tickets by destination country
	t, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return 0, fmt.Errorf("error getting the tickets by destination country: %w", internal.ErrGettingTickets)
	}

	// gets total tickets
	total, err := s.rp.Get(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("error getting the total tickets: %w", internal.ErrGettingTickets)
	}

	// calculates the percentage of tickets by destination country
	percentage = float64(len(t)) / float64(len(total)) * 100
	
	return
}

func (s *ServiceTicketDefault) GetTotalAmountTickets() (total float64, err error) {
	// gets total tickets
	t, err := s.rp.Get(context.TODO())
	if err != nil {
		return 0, fmt.Errorf("error getting the total tickets: %w", internal.ErrGettingTickets)
	}

	// calculates the total amount of tickets
	for _, v := range t {
		total += v.Price
	}
	return
}