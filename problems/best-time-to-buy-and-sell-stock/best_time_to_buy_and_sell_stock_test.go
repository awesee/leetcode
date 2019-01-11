package best_time_to_buy_and_sell_stock

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			input:    []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			input:    []int{},
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := maxProfit(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
