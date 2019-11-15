package problem867

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	want [][]int
}

func TestTranspose(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tt := range tests {
		got := transpose(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
