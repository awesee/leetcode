package string_to_integer_atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := map[string]int{
		"42":              42,
		"+42":             42,
		"   -42":          -42,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    -2147483648,
	}

	for input, expected := range tests {
		output := myAtoi(input)
		if output != expected {
			t.Fatalf(
				"Input: %v, Output: %v, Expected: %v",
				input, output, expected,
			)
		}
	}
}
