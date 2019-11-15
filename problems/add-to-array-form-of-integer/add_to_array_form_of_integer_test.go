package problem989

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	k    int
	want []int
}

func TestAddToArrayForm(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 0, 0},
			k:    34,
			want: []int{1, 2, 3, 4},
		},
		{
			in:   []int{2, 1, 5},
			k:    806,
			want: []int{1, 0, 2, 1},
		},
		{
			in:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			k:    1,
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		got := addToArrayForm(tt.in, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
