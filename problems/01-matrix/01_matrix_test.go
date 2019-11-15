package problem542

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	want [][]int
}

func TestUpdateMatrix(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			in: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		got := updateMatrix(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
