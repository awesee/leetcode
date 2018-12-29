package unique_morse_code_words

import "testing"

type caseType struct {
	input    []string
	expected int
}

func TestUniqueMorseRepresentations(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"gin", "zen", "gig", "msg"},
			expected: 2,
		},
		{
			input:    []string{"hello", "word"},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := uniqueMorseRepresentations(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
