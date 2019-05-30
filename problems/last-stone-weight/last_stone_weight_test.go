package last_stone_weight

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestLastStoneWeight(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			input:    []int{2, 7, 4, 1, 8, 1, 5},
			expected: 0,
		},
		{
			input:    []int{316, 157, 73, 106, 771, 828},
			expected: 37,
		},
	}
	for _, tc := range tests {
		output := lastStoneWeight(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
