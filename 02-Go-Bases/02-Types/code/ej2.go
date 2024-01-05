package main

import "fmt"

func main(){
	var age int
	var employed bool
	var antiquity int
	var salary float32

	for i:=0; i<3; i++ {
		fmt.Println("Edad del empleado: ")
		fmt.Scanln(&age)

		fmt.Println("Se encuentra empleado: ")
		fmt.Scanln(&employed)
		
		fmt.Println("Salario: ")
		fmt.Scanln(&salary)

		fmt.Println("Antigüedad: ")
		fmt.Scanln(&antiquity)

		if age>22 && employed==true && antiquity>=1 {
			if salary >= 100000.00 {
				fmt.Println("El empleado puede recibir préstamos sin interes.")
			} else {
				fmt.Println("El empleado puede recibir préstamos.")
			}
		} else {
			fmt.Println("El empleado no puede recibir préstamos.")
		}
	}
}