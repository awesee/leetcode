package problem832

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	want [][]int
}

func TestFlipAndInvertImage(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			in: [][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			want: [][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
				{1, 0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		got := flipAndInvertImage(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
