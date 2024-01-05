package main

import (
	"errors"
	"fmt"
)

var (
	ErrorMinimumHours = errors.New("Error: the worker cannot have worked less than 80 hours per month")
)

func main() {
	var hours int
	var rate float32

	fmt.Print("Enter the hours worked: ")
	fmt.Scan(&hours)
	fmt.Print("Enter the rate: ")
	fmt.Scan(&rate)

	salary, err := calculateSalary(hours, rate)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The salary is: $", salary)
	}
}

func calculateSalary(hours int, rate float32) (salary float32, err error){
	if(hours < 80){
		return 0, ErrorMinimumHours
	} else {
		salary = float32(hours) * rate
		if salary > 150000 {
			salary = salary * 0.9
		}
		return salary, nil
	}
}