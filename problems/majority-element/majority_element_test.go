package majority_element

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMajorityElement(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 2, 3},
			expected: 3,
		},
		{
			input:    []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			input:    []int{2, 2, 2, 2, 1, 1, 1},
			expected: 2,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := majorityElement(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
