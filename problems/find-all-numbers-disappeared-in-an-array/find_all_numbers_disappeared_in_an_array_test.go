package problem448

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestFindDisappearedNumbers(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			in:   []int{5, 4, 2, 3, 1},
			want: []int{},
		},
		{
			in:   []int{1, 1},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		got := findDisappearedNumbers(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
