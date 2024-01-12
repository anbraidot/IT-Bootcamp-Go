package storage

import (
	"03-POST/internal"
	"fmt"
	"io"
	"os"
)

// NewProductJSON creates a new product json
func NewProductJSON(data []byte, path string) *ProductJSON {
	// default config
	if data == nil {
		data = make([]byte, 0)
	}
	if path == "" {
		path = "03-Go-Web/03-POST/internal/storage/products.json"
	}

	// return the product json
	return &ProductJSON{
		data: data,
		path: path,
	}
}

// ProductJSON is the struct for the product json
type ProductJSON struct {
	// data is the data of the json
	data []byte
	// path is the path of the json
	path string
}

// Read reads the json file and converts it to a map[string]any
func (p *ProductJSON) Read() (data []byte, err error) {
	// open the file
	file, err := os.OpenFile(p.path, os.O_RDONLY, 0644)
	if err != nil {
		err = fmt.Errorf("error opening the file: %w", internal.ErrOpenFile)
		return
	}
	// close the file
	defer file.Close()

	// read the file
	data, err = io.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("error reading the file: %w", internal.ErrReadFile)
		return
	}
	return
}

// Write writes the map to a json file
func (p *ProductJSON) Write(data []byte) (err error) {
	// open the file
	file, err := os.OpenFile(p.path, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening the file: %w", internal.ErrOpenFile)
	}
	// close the file
	defer file.Close()

	// write the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing the file: %w", internal.ErrWriteFile)
	}
	return
}

func (p *ProductJSON) CreateFile() (err error) {
	// check if the file exists
	file, err := os.OpenFile(p.path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		err = fmt.Errorf("internal server error: %w", internal.ErrCreateFile)
			return
	}
	// close the file
	defer file.Close()

	return
}
