package loader

import (
	"04-Desafio/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (lt *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	// open the file
	file, err := os.Open(lt.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer file.Close()

	// read the file
	r := csv.NewReader(file)

	// read the records
	t = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			err = fmt.Errorf("error reading record: %v", err)
			return nil, err
		}

		// serialize the record
		id, err := strconv.Atoi(record[0])
		if err != nil {
			err = fmt.Errorf("error converting id to int: %v", err)
			return nil, err
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			err = fmt.Errorf("error converting price to float: %v", err)
			return nil, err
		}

		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price, 
		}

		// add the ticket to the map
		t[id] = ticket
	}

	return
}
