package ej1_test

import (
	"app/internal/ej1"
	"testing"
)

func TestTax(t *testing.T) {

	t.Run("success - case 01: salary is less than 50000.00", func(t *testing.T) {
		//arrange
		// ...

		//act
		salary := 40000.00
		result := ej1.Tax(salary) // Call the Tax function using the package name

		//assert
		expectedResult := 0.00
		if result != expectedResult {
			t.Errorf("Tax(%f) failed, expected %f, got %f", salary, expectedResult, result)
		}
	})

	t.Run("success - case 02: salary is greater than 50000.00 and less than 150000.00", func(t *testing.T) {
		//arrange
		// ...

		//act
		salary := 100000.00
		result := ej1.Tax(salary)

		//assert
		expectedResult := 17000.00
		if result != expectedResult {
			t.Errorf("Tax(%f) failed, expected %f, got %f", salary, expectedResult, result)
		}
	})

	t.Run("success - case 03: salary is greater than 150000.00", func(t *testing.T) {
		//arrange
		// ...

		//act
		salary := 200000.00
		result := ej1.Tax(salary)

		//assert
		expectedResult := 54000.00
		if result != expectedResult {
			t.Errorf("Tax(%f) failed, expected %f, got %f", salary, expectedResult, result)
		}
	})
}