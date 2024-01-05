package main

import "fmt"

type Student struct {
	Name string
	Surname string
	DNI int
	DateOfBirth string
}

func (s *Student) Details() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Surname: ", s.Surname)
	fmt.Println("DNI: ", s.DNI)
	fmt.Println("DateOfBirth: ", s.DateOfBirth)
}

func main() {
	s := Student{
		Name: "Bill",
		Surname: "Gates",
		DNI: 12345678,
	DateOfBirth: "28/10/1955",
	}
	s.Details()
}
