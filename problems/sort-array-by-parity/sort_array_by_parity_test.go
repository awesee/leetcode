package problem905

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestSortArrayByParity(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
	}
	for _, tt := range tests {
		got := sortArrayByParity(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
