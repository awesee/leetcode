package problem8

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

	for in, want := range tests {
		got := myAtoi(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
