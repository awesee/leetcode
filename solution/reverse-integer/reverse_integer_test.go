package reverse_integer

import (
	"testing"
	"math"
)

func TestReverse(t *testing.T) {
	tests := map[int]int{
		123:           321,
		-123:          -321,
		120:           21,
		math.MaxInt32: 0,
		math.MinInt32: 0,
	}

	for input, expected := range tests {
		output := reverse(input)
		if output != expected {
			t.Fatalf(
				"Input: %v, Output: %v, Expected: %v",
				input, output, expected,
			)
		}
	}
}
