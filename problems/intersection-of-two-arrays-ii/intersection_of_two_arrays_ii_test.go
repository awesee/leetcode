package problem350

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	nums1 []int
	nums2 []int
	want  []int
}

func TestIntersect(t *testing.T) {
	tests := [...]testType{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
		{
			nums1: []int{1, 3, 3, 5, 5, 5, 7, 7, 7, 9},
			nums2: []int{2, 3, 3, 5, 5, 5, 7, 8},
			want:  []int{3, 3, 5, 5, 5, 7},
		},
	}
	for _, tt := range tests {
		got := intersect(tt.nums1, tt.nums2)
		if !kit.IsEqualSliceInt(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
