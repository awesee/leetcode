package problem_415

import "testing"

type caseType struct {
	num1     string
	num2     string
	expected string
}

func TestAddStrings(t *testing.T) {
	tests := [...]caseType{
		{
			num1:     "0",
			num2:     "0",
			expected: "0",
		},
		{
			num1:     "1",
			num2:     "2",
			expected: "3",
		},
		{
			num1:     "9",
			num2:     "9",
			expected: "18",
		},
		{
			num1:     "100",
			num2:     "999",
			expected: "1099",
		},
	}
	for _, tc := range tests {
		output := addStrings(tc.num1, tc.num2)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc, output, tc.expected)
		}
	}
}
