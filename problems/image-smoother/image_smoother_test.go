package problem661

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	want [][]int
}

func TestImageSmoother(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{3, 3, 4},
				{4, 5, 5},
				{6, 6, 7},
			},
		},
	}
	for _, tt := range tests {
		got := imageSmoother(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
