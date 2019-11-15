package problem18

import (
	"reflect"
	"testing"
)

type testType struct {
	in     []int
	target int
	want   [][]int
}

func TestFourSum(t *testing.T) {
	tests := [...]testType{
		{
			in:     []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			in:     []int{-1, 0, 1, 2, -1, -4},
			target: 0,
			want: [][]int{
				{-1, -1, 0, 2},
			},
		},
		{
			in:     []int{0, 0, 0, 0},
			target: 0,
			want: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			in:     []int{-2, 0, 0, 2, 2, 2},
			target: 0,
			want: [][]int{
				{-2, 0, 0, 2},
			},
		},
		{
			in:     []int{-2, 0, 0, 2, 2, 2, 2},
			target: 0,
			want: [][]int{
				{-2, 0, 0, 2},
			},
		},
		{
			in:     []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target: 0,
			want: [][]int{
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-3, 0, 0, 3},
				{-3, 0, 1, 2},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			in:     []int{1, -2, -5, -4, -3, 3, 3, 5},
			target: -11,
			want: [][]int{
				{-5, -4, -3, 1},
			},
		},
	}
	for _, tt := range tests {
		got := fourSum(tt.in, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
