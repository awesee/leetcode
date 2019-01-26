package contains_duplicate_ii

import "testing"

type caseType struct {
	nums     []int
	k        int
	expected bool
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := containsNearbyDuplicate(tc.nums, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.nums, tc.k, output, tc.expected)
		}
	}
}
