package ej3

func SalaryCalc(minutesWorked int, category string) (result float64){
	switch category {
	case "A":
		result = float64(minutesWorked/60) * 1000.00
	case "B":
		result = float64(minutesWorked/60) * 1500.00
		result *= 1.20
	case "C":
		result = float64(minutesWorked/60) * 3000.00
		result *= 1.50
	default:
		result = 0.00
	}
	return
}