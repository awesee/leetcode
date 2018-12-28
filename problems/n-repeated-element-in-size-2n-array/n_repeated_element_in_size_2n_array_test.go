package n_repeated_element_in_size_2n_array

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestRepeatedNTimes(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 3},
			expected: 3,
		},
		{
			input:    []int{2, 1, 2, 5, 3, 2},
			expected: 2,
		},
		{
			input:    []int{5, 1, 5, 2, 5, 3, 5, 4},
			expected: 5,
		},
		{
			input:    []int{1, 2, 3},
			expected: 0,
		},
	}

	for _, tc := range tests {
		output := repeatedNTimes(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
