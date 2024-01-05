package tickets

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// errors
var (
	ErrInvalidLine = fmt.Errorf("the line is not valid")
	ErrInvalidTime = fmt.Errorf("the time is not valid")
)

// Ticket represents a ticket
type Ticket struct {
	Id                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         Time
	Price              int
}

// Time represents a time
type Time struct {
	Hour   int
	Minute int
}

// REQ 1
// GetTotalTicketsForDestination returns the total of tickets for a destination country
func GetTotalTicketsForDestination(destination string, tickets []Ticket) (result int) {

	// calculate the total of tickets for a destination country
	for _, ticket := range tickets {
		if ticket.DestinationCountry == destination {
			result++
		}
	}
	// return the total of tickets
	return result
}

// ParseTicket parses a line of the file and returns a Ticket
func ParseTicket(line []string) (Ticket, error) {
	var ticket Ticket

	// validate the line
	if len(line) != 6 {
		return ticket, ErrInvalidLine
	}

	// parse the line
	id, err := strconv.Atoi(line[0])
	if err != nil {
		return ticket, err
	}
	ticket.Id = id
	ticket.Name = line[1]
	ticket.Email = line[2]
	ticket.DestinationCountry = line[3]
	flightTime, err := StringToTime(line[4])
	if err != nil {
		return ticket, err
	}
	ticket.FlightTime = flightTime
	price, err := strconv.Atoi(line[5])
	if err != nil {
		return ticket, err
	}
	ticket.Price = price

	// return the ticket
	return ticket, nil
}

// REQ 2
// GetCountByPeriod returns the total of tickets for a period of the day
func GetCountByPeriod(time string, tickets []Ticket) int {
	switch time {
	case "early-morning":
		return GetEarlyMornings(tickets)
	case "morning":
		return GetMornings(tickets)
	case "afternoon":
		return GetAfternoons(tickets)
	case "night":
		return GetNights(tickets)
	default:
		return 0
	}
}

// GetEarlyMornings returns the total of tickets with early morning flight
func GetEarlyMornings(tickets []Ticket) (result int) {

	// calculate the total of tickets with early morning flight
	for _, ticket := range tickets {
		if ticket.FlightTime.Hour >= 0 && ticket.FlightTime.Hour < 7 {
			result++
		}
	}
	// return the total of tickets with early morning flight
	return result
}

// GetMornings returns the total of tickets with morning flight
func GetMornings(tickets []Ticket) (result int) {
	// calculate the total of tickets with morning flight
	for _, ticket := range tickets {
		if ticket.FlightTime.Hour >= 7 && ticket.FlightTime.Hour < 13 {
			result++
		}
	}
	// return the total of tickets with morning flight
	return result
}

// GetAfternoons returns the total of tickets with afternoon flight
func GetAfternoons(tickets []Ticket) (result int) {
	// calculate the total of tickets with afternoon flight
	for _, ticket := range tickets {
		if ticket.FlightTime.Hour >= 13 && ticket.FlightTime.Hour < 20 {
			result++
		}
	}
	// return the total of tickets with afternoon flight
	return result
}

// GetNights returns the total of tickets with night flight
func GetNights(tickets []Ticket) (result int) {
	// calculate the total of tickets with night flight
	for _, ticket := range tickets {
		if ticket.FlightTime.Hour >= 20 && ticket.FlightTime.Hour < 24 {
			result++
		}
	}
	// return the total of tickets with night flight
	return result
}

// StringToTime parses a string and returns a Time
func StringToTime(s string) (time Time, err error) {
	// split the string by :
	splitString := strings.Split(s, ":")

	// validate the string
	if len(splitString) != 2 {
		return time, ErrInvalidTime
	}

	// parse the string
	hour, err := strconv.Atoi(splitString[0])
	if err != nil {
		return time, ErrInvalidTime
	}
	if hour < 0 || hour > 23 {
		return time, ErrInvalidTime
	}
	time.Hour = hour

	minute, err := strconv.Atoi(splitString[1])
	if err != nil {
		return time, ErrInvalidTime
	}
	if minute < 0 || minute > 59 {
		return time, ErrInvalidTime
	}
	time.Minute = minute

	// return the time
	return time, nil
}

// REQ 3
// AverageDestination returns the average of tickets that fly to the destination country
func AverageDestination(destination string, tickets []Ticket) (result float64) {

	// validate that tickets is not empty
	if len(tickets) == 0 {
		return result
	}

	// calculate the average of tickets that fly to the destination country
	count := GetTotalTicketsForDestination(destination, tickets)

	total := len(tickets)

	// the result is the average of tickets between count and total of lines that contain in the file
	result = float64(count) / float64(total)

	// return the average of tickets that fly to the destination country
	return result
}

func LoadTickets(url string) (result []Ticket, err error) {
	// open the file
	file, err := os.Open(url)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a csv reader
	rd := csv.NewReader(file)

	// calculate the total of tickets
	for {
		line, err := rd.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		ticket, err := ParseTicket(line)
		if err != nil {
			fmt.Printf("error parsing ticket: %v\n", err)
		} else {
			result = append(result, ticket)
		}
	}

	// return the total of tickets
	return result, nil
}
