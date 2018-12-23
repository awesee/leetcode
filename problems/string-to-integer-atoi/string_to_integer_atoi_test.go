package string_to_integer_atoi

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := map[string]int{
		" ":               0,
		"42":              42,
		"+123":            123,
		"+123 +":          123,
		"++123":           0,
		" -42 ":           -42,
		" -123-":          -123,
		"--321":           0,
		"4193 with words": 4193,
		"words and 987":   0,
		"-91283472332":    math.MinInt32,
		"9876543210":      math.MaxInt32,
	}

	for input, expected := range tests {
		output := myAtoi(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
