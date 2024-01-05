package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("customers.txt")

	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

}