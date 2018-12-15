package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for input, expected := range tests {
		output := romanToInt(input)
		if output != expected {
			t.Fatalf("Input: %v, Output: %v, Expected: %v", input, output, expected)
		}
	}
}
