package problem58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := map[string]int{
		"Hello World":      5,
		"Hello":            5,
		"this is a string": 6,
		"":                 0,
		" ":                0,
		"a ":               1,
		"b  ":              1,
	}

	for in, want := range tests {
		got := lengthOfLastWord(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
