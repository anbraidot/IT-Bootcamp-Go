package main

import (
	"app/internal/tickets"
	"fmt"
)

func main() {
	ticketsList, err := tickets.LoadTickets("./docs/tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	total := tickets.GetTotalTicketsForDestination("Argentina", ticketsList)
	fmt.Println(" the total of tickets for Argentina is:", total)
}
