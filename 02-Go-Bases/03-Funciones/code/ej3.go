package main

import "fmt"

func main() {
	fmt.Println("Ingrese la cantidad de minutos trabajados: ")
	var minutesWorked int
	fmt.Scanln(&minutesWorked)
	fmt.Println("Ingrese la categoria: ")
	var category string
	fmt.Scanln(&category)
	fmt.Println("El salario es: ", salaryCalc(minutesWorked, category))
}

func salaryCalc(minutesWorked int, category string) (result float32){
	switch category {
	case "A":
		result = float32(minutesWorked/60) * 1000.00
	case "B":
		result = float32(minutesWorked/60) * 1500.00
		result *= 1.2
	case "C":
		result = float32(minutesWorked/60) * 3000.00
		result *= 1.5
	}
	return
}