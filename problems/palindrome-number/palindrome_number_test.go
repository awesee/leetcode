package palindrome_number

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

	for input, expected := range tests {
		output := isPalindrome(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
