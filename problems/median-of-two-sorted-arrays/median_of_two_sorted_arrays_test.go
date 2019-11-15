package problem4

import "testing"

type testType struct {
	num1 []int
	num2 []int
	want float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := [...]testType{
		{
			num1: []int{1, 3},
			num2: []int{2},
			want: 2.0,
		},
		{
			num1: []int{1, 2},
			num2: []int{3, 4},
			want: 2.5,
		},
		{
			num1: []int{1, 3, 5, 7},
			num2: []int{2, 4},
			want: 3.5,
		},
		{
			num1: []int{1, 3, 5},
			num2: []int{},
			want: 3,
		},
		{
			num1: []int{1, 2},
			num2: []int{3},
			want: 2,
		},
		{
			num1: []int{1, 1},
			num2: []int{1, 1},
			want: 1,
		},
		{
			num1: []int{1, 1, 1},
			num2: []int{1, 1},
			want: 1,
		},
		{
			num1: []int{1, 1},
			num2: []int{1, 1, 1},
			want: 1,
		},
		{
			num1: []int{1},
			num2: []int{2, 3, 4},
			want: 2.5,
		},
		{
			num1: []int{2, 3, 4},
			num2: []int{1},
			want: 2.5,
		},
		{
			num1: []int{},
			num2: []int{1},
			want: 1,
		},
		{
			num1: []int{},
			num2: []int{1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		got := findMedianSortedArrays(tt.num1, tt.num2)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.num1, tt.num2, got, tt.want)
		}
	}
}
