package problem1389

import (
	"reflect"
	"testing"
)

type testType struct {
	in    []int
	index []int
	want  []int
}

func TestCreateTargetArray(t *testing.T) {
	tests := [...]testType{
		{
			in:    []int{0, 1, 2, 3, 4},
			index: []int{0, 1, 2, 2, 1},
			want:  []int{0, 4, 1, 3, 2},
		},
		{
			in:    []int{1, 2, 3, 4, 0},
			index: []int{0, 1, 2, 3, 0},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			in:    []int{1},
			index: []int{0},
			want:  []int{1},
		},
	}
	for _, tt := range tests {
		got := createTargetArray(tt.in, tt.index)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
