package ej1

// Tax calculates the tax to pay based on the salary
func Tax(salary float64) (result float64) {
	switch{
		case salary > 150000.00:
			result += salary * 0.1
			fallthrough
		case salary > 50000:
			result += salary * 0.17
		}
		return
	}