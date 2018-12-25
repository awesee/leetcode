package length_of_last_word

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

	for input, expected := range tests {
		output := lengthOfLastWord(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
