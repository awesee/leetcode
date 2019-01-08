package number_of_segments_in_a_string

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestCountSegments(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "Hello, my name is John",
			expected: 5,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    " ",
			expected: 0,
		},
		{
			input:    " abc ",
			expected: 1,
		},
		{
			input:    " abc  def  ",
			expected: 2,
		},
		{
			input:    "love live! mu'sic forever",
			expected: 4,
		},
	}
	for _, tc := range tests {
		output := countSegments(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
