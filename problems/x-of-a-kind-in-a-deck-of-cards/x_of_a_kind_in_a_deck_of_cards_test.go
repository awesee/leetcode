package x_of_a_kind_in_a_deck_of_cards

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestHasGroupsSizeX(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4, 4, 3, 2, 1},
			expected: true,
		},
		{
			input:    []int{1, 1, 1, 2, 2, 2, 3, 3},
			expected: false,
		},
		{
			input:    []int{1},
			expected: false,
		},
		{
			input:    []int{1, 1},
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 2, 2, 2},
			expected: true,
		},
	}
	for _, tc := range tests {
		output := hasGroupsSizeX(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
