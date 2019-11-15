package problem922

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestSortArrayByParityII(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 2, 5, 7},
			want: []int{4, 5, 2, 7},
		},
		{
			in:   []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3},
			want: []int{2, 3, 0, 1, 4, 1, 0, 3, 4, 3},
		},
	}
	for _, tt := range tests {
		got := sortArrayByParityII(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
