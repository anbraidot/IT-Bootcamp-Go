package ej2_test

import (
	"testing"
	"app/internal/ej2"
)

func TestQualificationAvg(t *testing.T) {
	t.Run("success - case 01: qualifications is not empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{10, 10, 10, 10, 10}
		result := ej2.QualificationAvg(qualifications)
		
		//assert
		var expectedResult float32 = 10.00
		if result != expectedResult {
			t.Errorf("QualificationAvg(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 02: qualifications is empty", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{}
		result := ej2.QualificationAvg(qualifications)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("QualificationAvg(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})

	t.Run("success - case 03: qualifications contains negative values", func(t *testing.T) {
		//arrange
		// ...
		
		//act
		qualifications := []int{10, 10, 10, 10, 10, -10}
		result := ej2.QualificationAvg(qualifications)
		
		//assert
		var expectedResult float32 = 0.00
		if result != expectedResult {
			t.Errorf("QualificationAvg(%v) failed, expected %f, got %f", qualifications, expectedResult, result)
		}
	})
}