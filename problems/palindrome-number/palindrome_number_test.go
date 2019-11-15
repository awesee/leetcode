package problem9

import (
	"math"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := map[int]bool{
		0:             true,
		121:           true,
		12321:         true,
		123321:        true,
		-121:          false,
		10:            false,
		100:           false,
		12345:         false,
		math.MaxInt32: false,
	}

	for in, want := range tests {
		got := isPalindrome(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
