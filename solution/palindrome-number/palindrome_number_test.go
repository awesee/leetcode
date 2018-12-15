package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := map[int]bool{
		121:  true,
		-121: false,
		10:   false,
	}

	for input, expected := range tests {
		output := isPalindrome(input)
		if output != expected {
			t.Fatalf("Input: %v, Output: %v, Expected: %v", input, output, expected)
		}
	}
}
