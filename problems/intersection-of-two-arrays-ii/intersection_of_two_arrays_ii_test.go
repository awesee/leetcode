package intersection_of_two_arrays_ii

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	nums1    []int
	nums2    []int
	expected []int
}

func TestIntersect(t *testing.T) {
	tests := [...]caseType{
		{
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2, 2},
		},
		{
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
		{
			nums1:    []int{1, 3, 3, 5, 5, 5, 7, 7, 7, 9},
			nums2:    []int{2, 3, 3, 5, 5, 5, 7, 8},
			expected: []int{3, 3, 5, 5, 5, 7},
		},
	}
	for _, tc := range tests {
		output := intersect(tc.nums1, tc.nums2)
		if !IsEqualSliceInt(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.nums1, tc.nums2, output, tc.expected)
		}
	}
}
