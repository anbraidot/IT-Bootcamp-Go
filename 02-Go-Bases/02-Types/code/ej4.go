package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	//Inciso 1
	fmt.Println("La edad de Benjamin es: ",employees["Benjamin"])

	//Inciso 2
	var employeesOver21 int

	for employe,_ := range employees {
		if employees[employe] > 21 {
			employeesOver21++
		}
	}

	println("Cant. empleados mayores de 21 años: ",employeesOver21)

	//Inciso 3
	employees["Federico"] = 25
	fmt.Println(employees)

	//Inciso 4
	delete(employees, "Pedro")
	fmt.Println(employees)

}