package problem1042

import (
	"reflect"
	"testing"
)

type testType struct {
	n     int
	paths [][]int
	want  []int
}

func TestGardenNoAdj(t *testing.T) {
	tests := [...]testType{
		{
			n: 3,
			paths: [][]int{
				{1, 2},
				{2, 3},
				{3, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			n: 4,
			paths: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 1, 2},
		},
		{
			n: 4,
			paths: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 1},
				{1, 3},
				{2, 4},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		got := gardenNoAdj(tt.n, tt.paths)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.n, got, tt.want)
		}
	}
}
