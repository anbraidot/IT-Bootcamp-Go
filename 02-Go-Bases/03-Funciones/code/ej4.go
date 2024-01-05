package main

import "fmt"

func main() {

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	function, err := operation(minimum)
	if err == "" {
		minValue := function(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("El mínimo es: ",minValue)
	}

	function, err = operation(average)
	if err == "" {
		averageValue := function(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("El promedio es: ",averageValue)
	}

	function, err = operation(maximum)
	if err == "" {
		maxValue := function(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("El máximo es: ",maxValue)
	}
}

func operation(operation string) (func(...int) float32, string) {
	switch operation {
	case "minimum":
		return minFunc, ""
	case "average":
		return averageFunc, ""
	case "maximum":
		return maxFunc, ""
	default:
		return nil, "Invalid operation"
	}
}

func minFunc(qualifications ...int) float32 {
	result := float32(qualifications[0])
	for _, value := range qualifications {
		if value < int(result) {
			result = float32(value)
		}
	}
	return result
}

func averageFunc(qualifications ...int) (result float32) {
	for _, value := range qualifications {
		result += float32(value)
	}
	result /= float32(len(qualifications))
	return
}

func maxFunc(qualifications ...int) float32 {
	result := float32(qualifications[0])
	for _, value := range qualifications {
		if value > int(result) {
			result = float32(value)
		}
	}
	return result
}