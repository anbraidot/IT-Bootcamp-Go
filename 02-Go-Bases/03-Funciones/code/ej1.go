package main

import "fmt"

func main() {
	var salary float32
	fmt.Println("Ingrese su salario: ")
	fmt.Scanln(&salary)
	fmt.Println("El impuesto a pagar es: ", tax(salary))
}

func tax(salary float32) (result float32) {
	switch{
		case salary > 150000.00:
			result += salary * 0.1
			fallthrough
		case salary > 50000:
			result += salary * 0.17
		}
		return
	}