package problem912

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestSortArray(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{5, 2, 3, 1},
			want: []int{1, 2, 3, 5},
		},
		{
			in:   []int{5, 1, 1, 2, 0, 0},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		got := sortArray(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
