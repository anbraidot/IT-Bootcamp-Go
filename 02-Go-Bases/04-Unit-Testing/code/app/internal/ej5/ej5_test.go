package ej5_test

import (
	"app/internal/ej5"
	"testing"
)

func TestAnimalDog(t *testing.T) {
	t.Run("success - case 01: quantity is greater than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 10
		result := ej5.AnimalDog(quantity)

		//assert
		var expectedResult float64 = 100.00
		if result != expectedResult {
			t.Errorf("AnimalDog(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 02: quantity is 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 0
		result := ej5.AnimalDog(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalDog(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 03: quantity is less than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := -10
		result := ej5.AnimalDog(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalDog(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})
}

func TestAnimalCat(t *testing.T) {
	t.Run("success - case 01: quantity is greater than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 10
		result := ej5.AnimalCat(quantity)

		//assert
		var expectedResult float64 = 50.00
		if result != expectedResult {
			t.Errorf("AnimalCat(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 02: quantity is 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 0
		result := ej5.AnimalCat(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalCat(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 03: quantity is less than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := -10
		result := ej5.AnimalCat(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalCat(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})
}

func TestAnimalTarantula(t *testing.T) {
	t.Run("success - case 01: quantity is greater than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 10
		result := ej5.AnimalTarantula(quantity)

		//assert
		var expectedResult float64 = 1.50
		if result != expectedResult {
			t.Errorf("AnimalTarantula(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 02: quantity is 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 0
		result := ej5.AnimalTarantula(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalTarantula(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 03: quantity is less than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := -10
		result := ej5.AnimalTarantula(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalTarantula(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})
}

func TestAnimalHamster(t *testing.T) {
	t.Run("success - case 01: quantity is greater than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 10
		result := ej5.AnimalHamster(quantity)

		//assert
		var expectedResult float64 = 2.50
		if result != expectedResult {
			t.Errorf("AnimalHamster(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 02: quantity is 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := 0
		result := ej5.AnimalHamster(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalHamster(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})

	t.Run("success - case 03: quantity is less than 0", func(t *testing.T) {
		//arrange
		// ...

		//act
		quantity := -10
		result := ej5.AnimalHamster(quantity)

		//assert
		var expectedResult float64 = 0.00
		if result != expectedResult {
			t.Errorf("AnimalHamster(%d) failed, expected %f, got %f", quantity, expectedResult, result)
		}
	})
}