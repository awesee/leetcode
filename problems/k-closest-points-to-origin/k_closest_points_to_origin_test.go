package problem973

import (
	"reflect"
	"testing"
)

type testType struct {
	in   [][]int
	k    int
	want [][]int
}

func TestKClosest(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 3},
				{-2, 2},
			},
			k: 1,
			want: [][]int{
				{-2, 2},
			},
		},
		{
			in: [][]int{
				{3, 3},
				{5, -1},
				{-2, 4},
			},
			k: 2,
			want: [][]int{
				{3, 3},
				{-2, 4},
			},
		},
	}
	for _, tt := range tests {
		got := kClosest(tt.in, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
