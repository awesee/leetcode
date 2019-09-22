package problem_65

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestIsNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "0",
			expected: true,
		},
		{
			input:    " 0.1",
			expected: true,
		},
		{
			input:    "abc",
			expected: false,
		},
		{
			input:    "1 a",
			expected: false,
		},
		{
			input:    "2e10",
			expected: true,
		},
		{
			input:    " -90e3   ",
			expected: true,
		},
		{
			input:    " 1e",
			expected: false,
		},
		{
			input:    "e3",
			expected: false,
		},
		{
			input:    " 6e-1",
			expected: true,
		},
		{
			input:    " 99e2.5 ",
			expected: false,
		},
		{
			input:    "53.5e93",
			expected: true,
		},
		{
			input:    " --6 ",
			expected: false,
		},
		{
			input:    "-+3",
			expected: false,
		},
		{
			input:    "95a54e53",
			expected: false,
		},
		{
			input:    "078332e437",
			expected: true,
		},
	}
	for _, tc := range tests {
		output := isNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
