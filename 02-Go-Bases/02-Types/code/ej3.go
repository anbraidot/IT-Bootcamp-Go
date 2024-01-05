package main

import "fmt"

func main() {
	var month int

	fmt.Println("Ingrese el número del mes: ")
	fmt.Scanln(&month)

	switch month {
	case 1:
		fmt.Println("01, Enero")
	case 2:
		fmt.Println("02, Febrero")
	case 3:
		fmt.Println("03, Marzo")
	case 4:
		fmt.Println("04, Abril")
	case 5:
		fmt.Println("05, Mayo")
	case 6:
		fmt.Println("06, Junio")
	case 7:
		fmt.Println("07, Julio")
	case 8:
		fmt.Println("08, Agosto")
	case 9:
		fmt.Println("09, Septiembre")
	case 10:
		fmt.Println("10, Octubre")
	case 11:
		fmt.Println("11, Noviembre")
	case 12:
		fmt.Println("12, Diciembre")
	default:
		fmt.Println("No existe un mes con número ",month)
	}
}
