package main

import (
	"errors"
	"fmt"
)

type ErrorSalary struct {
	Msg string
}

var (
	ErrLowSalary = errors.New("Error: salary is less than 10000"))

func (e ErrorSalary) Error() string {
	return e.Msg
}

func main() {
	var salary int

	fmt.Print("Enter your salary: ")
	fmt.Scan(&salary)

	err := validateSalary(salary)
	if err!= nil {
		if errors.Is(err, ErrLowSalary) {
			fmt.Println(err.Error())
		}
	}
}

func validateSalary(salary int) (err error) {
	if salary < 10000 {
		return ErrLowSalary
	}
	return nil
}