package rotate_string

import "testing"

type caseType struct {
	A        string
	B        string
	expected bool
}

func TestRotateString(t *testing.T) {
	tests := [...]caseType{
		{
			A:        "abcde",
			B:        "cdeab",
			expected: true,
		},
		{
			A:        "abcde",
			B:        "abced",
			expected: false,
		},
		{
			A:        "abcde",
			B:        "abcdef",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := rotateString(tc.A, tc.B)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.A, tc.B, output, tc.expected)
		}
	}
}
