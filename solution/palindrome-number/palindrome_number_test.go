package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := map[int]bool{
		0:      true,
		121:    true,
		12321:  true,
		123321: true,
		-121:   false,
		10:     false,
		100:    false,
		12345:  false,
	}

	for input, expected := range tests {
		output := isPalindrome(input)
		if output != expected {
			t.Fatalf("Input: %v, Output: %v, Expected: %v", input, output, expected)
		}
	}
}
