package problem_1154

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestDayOfYear(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "2019-01-09",
			expected: 9,
		},
		{
			input:    "2019-02-10",
			expected: 41,
		},
		{
			input:    "2003-03-01",
			expected: 60,
		},
		{
			input:    "2004-03-01",
			expected: 61,
		},
		{
			input:    "2000-08-01",
			expected: 214,
		},
		{
			input:    "1993-12-11",
			expected: 345,
		},
	}
	for _, tc := range tests {
		output := dayOfYear(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
