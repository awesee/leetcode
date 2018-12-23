package search_insert_position

import "testing"

type caseType struct {
	nums     []int
	target   int
	expected int
}

func TestSearchInsert(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{2, 7, 11, 15},
			target:   5,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
	}

	for _, tc := range tests {
		output := searchInsert(tc.nums, tc.target)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.nums, output, tc.expected)
		}
	}
}
