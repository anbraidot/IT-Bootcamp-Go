package main

import (
	"fmt"
)

type ErrorSalary struct {
	Msg string
}

func (e ErrorSalary) Error() string {
	return e.Msg
}

func main() {
	var salary int

	fmt.Print("Enter your salary: ")
	fmt.Scan(&salary)

	err := validateSalary(salary)
	if err!= nil {
		fmt.Println(err.Error())
	}
}

func validateSalary(salary int) (err error) {
	if salary < 10000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary is: %d", salary)
	}
	return nil
}