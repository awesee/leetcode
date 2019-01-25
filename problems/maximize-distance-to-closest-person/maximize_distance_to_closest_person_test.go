package maximize_distance_to_closest_person

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMaxDistToClosest(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 0, 0, 0, 1, 0, 1},
			expected: 2,
		},
		{
			input:    []int{1, 0, 0, 0},
			expected: 3,
		},
		{
			input:    []int{0, 0, 0, 1},
			expected: 3,
		},
		{
			input:    []int{0, 0, 0, 0, 0, 1, 0},
			expected: 5,
		},
		{
			input:    []int{0, 1, 0, 0, 0, 1, 0},
			expected: 2,
		},
		{
			input:    []int{0, 1, 0, 0, 0, 0, 0},
			expected: 5,
		},
		{
			input:    []int{1, 0, 1},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := maxDistToClosest(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
