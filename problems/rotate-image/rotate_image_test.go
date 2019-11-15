package problem48

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	want [][]int
}

func TestRotate(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			in: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tt := range tests {
		l := len(tt.in)
		got := make([][]int, l)
		for i := 0; i < l; i++ {
			got[i] = make([]int, l)
			copy(got[i], tt.in[i])
		}
		rotate(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
