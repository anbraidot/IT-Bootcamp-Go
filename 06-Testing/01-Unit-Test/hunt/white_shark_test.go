package hunt_test

import (
	"06-Testing/01-Unit-Test/hunt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrange
		// - create a white shark
		w := hunt.NewWhiteShark(true, false, 10)
		// - create a tuna
		tuna := hunt.NewTuna("tuna 1", 5)

		// act
		err := w.Hunt(tuna)

		// assert
		if err != nil {
			t.Errorf("expected error nil, got %v", err)
		}
		if w.Hungry {
			t.Errorf("expected hungry false, got true")
		}
		if !w.Tired {
			t.Errorf("expected tired true, got false")
		}
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrange
		// - create a white shark
		w := hunt.NewWhiteShark(false, false, 10)
		// - create a tuna
		tuna := hunt.NewTuna("tuna 1", 5)

		// act
		err := w.Hunt(tuna)

		// assert
		expectError := hunt.ErrSharkIsNotHungry
		require.Error(t, expectError, err)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrange
		// - create a white shark
		w := hunt.NewWhiteShark(true, true, 10)
		// - create a tuna
		tuna := hunt.NewTuna("tuna 1", 5)

		// act
		err := w.Hunt(tuna)

		// assert
		expectError := hunt.ErrSharkIsTired
		require.Error(t, expectError, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrange
		// - create a white shark
		w := hunt.NewWhiteShark(true, false, 5)
		// - create a tuna
		tuna := hunt.NewTuna("tuna 1", 10)

		// act
		err := w.Hunt(tuna)

		// assert
		expectError := hunt.ErrSharkIsSlower
		require.Error(t, expectError, err)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrange
		// - create a white shark
		w := hunt.NewWhiteShark(true, false, 10)

		// act
		err := w.Hunt(nil)

		// assert
		expectError := hunt.ErrTunaIsNil
		require.Error(t, expectError, err)
	})
}
