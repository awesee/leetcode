package k_diff_pairs_in_an_array

import "testing"

type caseType struct {
	nums     []int
	k        int
	expected int
}

func TestFindPairs(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{1, 3, 1, 5, 4},
			k:        0,
			expected: 1,
		},
		{
			nums:     []int{3, 1, 4, 1, 5},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 4,
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        -1,
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := findPairs(tc.nums, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.nums, tc.k, output, tc.expected)
		}
	}
}
