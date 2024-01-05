package ej3_test

import(
	"testing"
	"app/internal/ej3"
)

func TestSalaryCalc(t *testing.T) {

	t.Run("success - case 01: category A", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		minutesWorked := 60
		category := "A"
		result := ej3.SalaryCalc(minutesWorked, category)
		
		//assert
		var expectedResult float64 = 1000.00
		if result != expectedResult {
			t.Errorf("SalaryCalc(%d, %s) failed, expected %f, got %f", minutesWorked, category, expectedResult, result)
		}
	})

	t.Run("success - case 02: category B", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		minutesWorked := 60
		category := "B"
		result := ej3.SalaryCalc(minutesWorked, category)
		
		//assert
		var expectedResult float64 = 1800.00
		if result != expectedResult {
			t.Errorf("SalaryCalc(%d, %s) failed, expected %f, got %f", minutesWorked, category, expectedResult, result)
		}
	})

	t.Run("success - case 03: category C", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		minutesWorked := 60
		category := "C"
		result := ej3.SalaryCalc(minutesWorked, category)
		
		//assert
		var expectedResult float64 = 4500.00
		if result != expectedResult {
			t.Errorf("SalaryCalc(%d, %s) failed, expected %f, got %f", minutesWorked, category, expectedResult, result)
		}
	})

	t.Run("success - case 04: invalid category", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		minutesWorked := 60
		category := "D"
		result := ej3.SalaryCalc(minutesWorked, category)
		
		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("SalaryCalc(%d, %s) failed, expected %f, got %f", minutesWorked, category, expectedResult, result)
		}
	})
}