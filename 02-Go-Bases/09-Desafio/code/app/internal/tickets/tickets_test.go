package tickets_test

import (
	"app/internal/tickets"
	"testing"

	"github.com/stretchr/testify/require"
)

// REQ 4
func TestGetTotalTicketsForDestination(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}

	t.Run(" success 1: destination is empty string", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetTotalTicketsForDestination("", ticketList)

		//assert
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("GetTotalTicketsForDestination() failed, expected %d, got %d", expectedResult, result)
		}
	})

	t.Run(" success 2: destination is Argentina", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetTotalTicketsForDestination("Argentina", ticketList)
		if err != nil {
			t.Errorf("GetTotalTicketsForDestination() failed, expected nil, got %v", err)
		}

		//assert
		expectedResult := 15
		if result != expectedResult {
			t.Errorf("GetTotalTicketsForDestination() failed, expected %d, got %d", expectedResult, result)
		}
	})
}

func TestParseTicket(t *testing.T) {
	t.Run(" success 1: line is valid", func(t *testing.T) {
		//arrange
		line := []string{"1", "Tait Mc Caughan", "tmc0@scribd.com", "Finland", "17:11", "785"}

		//act
		result, _ := tickets.ParseTicket(line)

		//assert
		expectedResult := tickets.Ticket{1, "Tait Mc Caughan", "tmc0@scribd.com", "Finland", tickets.Time{17, 11}, 785}
		if result != expectedResult {
			t.Errorf("ParseTicket() failed, expected %v, got %v", expectedResult, result)
		}
	})

	t.Run(" error 1: line is not valid", func(t *testing.T) {
		//arrange
		line := []string{"1", "Tait Mc Caughan", "tmc0@scribd.com", "Finland", "17:11", "785", "otherColum1", "otherColum2"}

		//act
		_, err := tickets.ParseTicket(line)

		//assert
		require.Equal(t, err, tickets.ErrInvalidLine)
	})
}

func TestGetEarlyMornings(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	ticketListEmpty, err := tickets.LoadTickets("../../docs/tickets_empty.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}

	t.Run(" success 1: there are tickets for early morning", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetEarlyMornings(ticketList)

		//assert
		expectedResult := 304
		if result != expectedResult {
			t.Errorf("GetEarlyMornings() failed, expected %d, got %d", expectedResult, result)
		}
	})

	t.Run(" success 2: there are no tickets for early morning", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetEarlyMornings(ticketListEmpty)

		//assert
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("GetEarlyMornings() failed, expected %d, got %d", expectedResult, result)
		}
	})
}

func TestoGetMornings(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	ticketListEmpty, err := tickets.LoadTickets("../../docs/tickets_empty.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}

	t.Run(" success 1: there are tickets for morning", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetMornings(ticketList)

		//assert
		expectedResult := 304
		if result != expectedResult {
			t.Errorf("GetMornings() failed, expected %d, got %d", expectedResult, result)
		}
	})

	t.Run(" success 2: there are no tickets for morning", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetMornings(ticketListEmpty)

		//assert
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("GetMornings() failed, expected %d, got %d", expectedResult, result)
		}
	})
}

func TestGetAfternoons(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	ticketListEmpty, err := tickets.LoadTickets("../../docs/tickets_empty.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}

	t.Run(" success 1: there are tickets for afternoon", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetAfternoons(ticketList)

		//assert
		expectedResult := 289
		if result != expectedResult {
			t.Errorf("GetAfternoons() failed, expected %d, got %d", expectedResult, result)
		}
	})

	t.Run(" success 2: there are no tickets for afternoon", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetAfternoons(ticketListEmpty)

		//assert
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("GetAfternoons() failed, expected %d, got %d", expectedResult, result)
		}
	})
}

func TestGetNights(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	ticketListEmpty, err := tickets.LoadTickets("../../docs/tickets_empty.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	t.Run(" success 1: there are tickets for night", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetNights(ticketList)

		//assert
		expectedResult := 151
		if result != expectedResult {
			t.Errorf("GetNights() failed, expected %d, got %d", expectedResult, result)
		}
	})

	t.Run(" success 2: there are no tickets for night", func(t *testing.T) {
		//arrange

		//act
		result := tickets.GetNights(ticketListEmpty)

		//assert
		expectedResult := 0
		if result != expectedResult {
			t.Errorf("GetNights() failed, expected %d, got %d", expectedResult, result)
		}
	})
}

func TestStringToTime(t *testing.T) {
	t.Run(" success 1: time is valid", func(t *testing.T) {
		//arrange
		time := "17:11"

		//act
		result, err := tickets.StringToTime(time)
		if err != nil {
			t.Errorf("StringToTime() failed, expected nil, got %v", err)
		}

		//assert
		expectedResult := tickets.Time{17, 11}
		if result != expectedResult {
			t.Errorf("StringToTime() failed, expected %v, got %v", expectedResult, result)
		}
	})

	t.Run(" success 2: time is not valid", func(t *testing.T) {
		//arrange
		time := "17:11:00"

		//act
		result, err := tickets.StringToTime(time)
		if err != tickets.ErrInvalidTime {
			t.Errorf("StringToTime() failed, expected %v, got %v", tickets.ErrInvalidTime, err)
		}

		//assert
		expectedResult := tickets.Time{}
		if result != expectedResult {
			t.Errorf("StringToTime() failed, expected %v, got %v", expectedResult, result)
		}
	})

	t.Run(" error 1: time contain letters", func(t *testing.T) {
		//arrange
		time := "17:11a"

		//act
		_, err := tickets.StringToTime(time)

		//assert
		require.Equal(t, err, tickets.ErrInvalidTime)
	})

	t.Run(" error 3: time is empty string", func(t *testing.T) {
		//arrange
		time := ""

		//act
		result, err := tickets.StringToTime(time)
		if err != tickets.ErrInvalidTime {
			t.Errorf("StringToTime() failed, expected %v, got %v", tickets.ErrInvalidTime, err)
		}

		//assert
		expectedResult := tickets.Time{}
		if result != expectedResult {
			t.Errorf("StringToTime() failed, expected %v, got %v", expectedResult, result)
		}
	})
}

func TestAverageDestination(t *testing.T) {

	// load the tickets
	ticketList, err := tickets.LoadTickets("../../docs/tickets.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}
	ticketListEmpty, err := tickets.LoadTickets("../../docs/tickets_empty.csv")
	if err != nil {
		t.Errorf("LoadTickets() failed, expected nil, got %v", err)
	}

	t.Run(" success 1: there are tickets for destination", func(t *testing.T) {
		//arrange

		//act
		result := tickets.AverageDestination("Argentina", ticketList)

		//assert
		expectedResult := 0.015
		if result != expectedResult {
			t.Errorf("AverageDestination() failed, expected %f, got %f", expectedResult, result)
		}
	})

	t.Run(" success 2: there are no tickets for destination", func(t *testing.T) {
		//arrange

		//act
		result := tickets.AverageDestination("Argentina", ticketListEmpty)

		//assert
		expectedResult := 0.0
		if result != expectedResult {
			t.Errorf("AverageDestination() failed, expected %f, got %f", expectedResult, result)
		}
	})
}
