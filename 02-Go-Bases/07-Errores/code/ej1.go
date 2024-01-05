package main

import "fmt"

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

	if salary < 150000 {
		err := ErrorSalary{
			Msg: "The salary entered does not reach the taxable minimum",
		}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}
}