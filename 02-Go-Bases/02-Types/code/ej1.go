package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Ingrese su palabra: ")

	var word string
	fmt.Scanln(&word)

	fmt.Println("Su palabra es: ", word)
	fmt.Println("Cant. de letras: ", len(word))
	fmt.Println("Deletreado: ", strings.Split(word, ""))
}