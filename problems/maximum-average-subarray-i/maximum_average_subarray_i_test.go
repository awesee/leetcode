package maximum_average_subarray_i

import "testing"

type caseType struct {
	nums     []int
	k        int
	expected float64
}

func TestFindMaxAverage(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
	}
	for _, tc := range tests {
		output := findMaxAverage(tc.nums, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.nums, tc.k, output, tc.expected)
		}
	}
}
