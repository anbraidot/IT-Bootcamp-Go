package positioner_test

import (
	"06-Testing/02-Double-Test/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPositionerDefault_GetLinearDistance tests the GetLinearDistance method
func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	t.Run("case 1: the coordinates are negative", func(t *testing.T) {
		// arrange
		// - create a position
		position1 := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}
		position2 := &positioner.Position{
			X: -2,
			Y: -2,
			Z: -2,
		}
		// - create a positioner
		pr := positioner.NewPositionerDefault()

		// act
		// - call the GetLinearDistance method
		result := pr.GetLinearDistance(position1, position2)

		// assert
		expectedResult := 1.7320508075688772
		// - check the result
		require.Equal(t, expectedResult, result)
	})

t.Run("case 2: the coordinates are positive", func(t *testing.T) {
		// arrange
		// - create a position
		position1 := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		position2 := &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}
		// - create a positioner
		pr := positioner.NewPositionerDefault()

		// act
		// - call the GetLinearDistance method
		result := pr.GetLinearDistance(position1, position2)

		// assert
		expectedResult := 1.7320508075688772
		// - check the result
		require.Equal(t, expectedResult, result)
	})

	t.Run("case 3: the coordinates return a linear distance without decimals", func(t *testing.T) {
		// arrange
		// - create a position
		position1 := &positioner.Position{
			X: 0,
			Y: 0,
			Z: 0,
		}
		position2 := &positioner.Position{
			X: 1,
			Y: 0,
			Z: 0,
		}
		// - create a positioner
		pr := positioner.NewPositionerDefault()

		// act
		// - call the GetLinearDistance method
		result := pr.GetLinearDistance(position1, position2)

		// assert
		expectedResult := 1.0
		// - check the result
		require.Equal(t, expectedResult, result)
	})
}
