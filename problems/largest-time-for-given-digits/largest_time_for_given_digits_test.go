package largest_time_for_given_digits

import "testing"

type caseType struct {
	input    []int
	expected string
}

func TestLargestTimeFromDigits(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4},
			expected: "23:41",
		},
		{
			input:    []int{5, 5, 5, 5},
			expected: "",
		},
		{
			input:    []int{2, 0, 6, 6},
			expected: "06:26",
		},
	}
	for _, tc := range tests {
		output := largestTimeFromDigits(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
