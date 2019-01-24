package min_cost_climbing_stairs

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{10, 15, 20},
			expected: 15,
		},
		{
			input:    []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}
	for _, tc := range tests {
		output := minCostClimbingStairs(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
