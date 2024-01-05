package ej4_test

import (
	"testing"
	"app/internal/ej4"
)

func TestMinFunction(t *testing.T) {

	t.Run("success - case 01: qualifications is not empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, 3, 4, 5}
		result := ej4.MinFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 1.00
		if result != expectedResult {
			t.Errorf("MinFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 02: qualifications is empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{}
		result := ej4.MinFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("MinFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 03: qualifications contains negative numbers", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, -3, 4, 5}
		result := ej4.MinFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("MinFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})
}

func TestAverageFunction(t *testing.T) {

	t.Run("success - case 01: qualifications is not empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, 3, 4, 5}
		result := ej4.AverageFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 3.00
		if result != expectedResult {
			t.Errorf("AverageFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 02: qualifications is empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{}
		result := ej4.AverageFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("AverageFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 03: qualifications contains negative numbers", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, -3, 4, 5}
		result := ej4.AverageFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("AverageFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})
}

func TestMaxFunction(t *testing.T) {

	t.Run("success - case 01: qualifications is not empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, 3, 4, 5}
		result := ej4.MaxFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 5.00
		if result != expectedResult {
			t.Errorf("MaxFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 02: qualifications is empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{}
		result := ej4.MaxFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("MaxFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 03: qualifications contains negative numbers", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{1, 2, -3, 4, 5}
		result := ej4.MaxFunc(qualifications...)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("MaxFunc(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})
}