package collatzConjecture

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollatz(t *testing.T) {
	testCases := []struct {
		Input    int
		Expected int
	}{
		{
			Input:    1,
			Expected: 1,
		},
		{
			Input:    12,
			Expected: 9,
		},
		{
			Input:    19,
			Expected: 20,
		},
		{
			Input:    27,
			Expected: 111,
		},
	}

	t.Run("should successfully obtain result", func(t *testing.T) {
		for _, tc := range testCases {
			result, err := Collatz(tc.Input)

			assert.NoError(t, err)
			assert.Equal(t, result, tc.Expected)
		}
	})

	t.Run("should fail, when value is 0", func(t *testing.T) {
		_, err := Collatz(0)
		assert.Error(t, err, errors.New("invalid value"))
	})

	t.Run("should fail, when value is negative", func(t *testing.T) {
		_, err := Collatz(-2)
		assert.Error(t, err, errors.New("invalid value"))
	})
}
