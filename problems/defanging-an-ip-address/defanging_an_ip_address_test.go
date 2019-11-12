package problem1108

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestDefangIPaddr(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "1.1.1.1",
			expected: "1[.]1[.]1[.]1",
		},
		{
			input:    "255.100.50.0",
			expected: "255[.]100[.]50[.]0",
		},
	}
	for _, tc := range tests {
		output := defangIPaddr(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
