package merge_sorted_array

import (
	"reflect"
	"testing"
)

type caseType struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

func TestMerge(t *testing.T) {
	tests := [...]caseType{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{-9, -6, -3, 0, 0, 0},
			m:        3,
			nums2:    []int{-7, -5, -3},
			n:        3,
			expected: []int{-9, -7, -6, -5, -3, -3},
		},
		{
			nums1:    []int{1, 2, 4, 5, 6, 0},
			m:        5,
			nums2:    []int{3},
			n:        1,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range tests {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		if !reflect.DeepEqual(tc.nums1, tc.expected) {
			t.Fatalf("input: %v %v %v %v, output: %v, expected: %v", tc.nums1, tc.m, tc.nums2, tc.n, tc.nums1, tc.expected)
		}
	}
}
