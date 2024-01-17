package service_test

import (
	"04-Desafio/internal"
	"04-Desafio/internal/repository"
	"04-Desafio/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - create db
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
		}
		// - repository
		rp := repository.NewRepositoryTicketMap(db,len(db))
		// - service
		sv := service.NewServiceTicketDefault(*rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
