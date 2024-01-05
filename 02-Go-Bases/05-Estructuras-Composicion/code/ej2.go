package main

import "fmt"

// Person is a struct that represents a person
type Person struct {
	// ID is the ID of the person
	ID   int
	// Name is the name of the person
	Name string
	// DateOfBirth is the date of birth of the person
	DateOfBirth string
}

// Employee is a struct that represents an employee
type Employee struct {
	// ID is the ID of the employee
	ID  int
	// Position is the position of the employee
	Position string
	// Person is the person of the employee (embedded struct)
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Println("ID: ", e.ID)
	fmt.Println("Position: ", e.Position)
	fmt.Println("Name: ", e.Name)
	fmt.Println("DateOfBirth: ", e.DateOfBirth)
}

func main() {
	p1 := Person{	
		ID: 1,
		Name: "Person 1",
		DateOfBirth: "01/01/2000",
	}

	e1 := Employee{
		ID: 1,
		Position: "Position 1",
		Person: p1,
	}

	e1.PrintEmployee()
}
