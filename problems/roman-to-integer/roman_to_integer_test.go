package problem13

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := map[string]int{
		"III":     3,
		"IV":      4,
		"IX":      9,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}

	for in, want := range tests {
		got := romanToInt(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
