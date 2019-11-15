package problem88

import (
	"reflect"
	"testing"
)

type testType struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	want  []int
}

func TestMerge(t *testing.T) {
	tests := [...]testType{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{-9, -6, -3, 0, 0, 0},
			m:     3,
			nums2: []int{-7, -5, -3},
			n:     3,
			want:  []int{-9, -7, -6, -5, -3, -3},
		},
		{
			nums1: []int{1, 2, 4, 5, 6, 0},
			m:     5,
			nums2: []int{3},
			n:     1,
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		if !reflect.DeepEqual(tt.nums1, tt.want) {
			t.Fatalf("in: %v %v %v %v, got: %v, want: %v", tt.nums1, tt.m, tt.nums2, tt.n, tt.nums1, tt.want)
		}
	}
}
