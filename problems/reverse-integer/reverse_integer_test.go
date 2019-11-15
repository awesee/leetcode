package problem7

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := map[int]int{
		123:           321,
		-123:          -321,
		120:           21,
		math.MaxInt32: 0,
		math.MinInt32: 0,
	}

	for in, want := range tests {
		got := reverse(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
